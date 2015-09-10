package rfc5424

import (
	"io"
	"reflect"
	"time"
)

var TimeNow = time.Now

func Encode(ob interface{}) *Message {
	mt := reflect.TypeOf(ob)
	mv := reflect.ValueOf(ob)

	reflection := Reflect(mt)

	m := Message{}

	severity := reflection.SeverityDefault
	if reflection.SeverityFieldIndex >= 0 {
		severity = mv.Field(reflection.SeverityFieldIndex).Interface().(Severity)
	}

	facility := reflection.FacilityDefault
	if reflection.FacilityFieldIndex >= 0 {
		facility = mv.Field(reflection.FacilityFieldIndex).Interface().(Facility)
	}
	m.Priority = int(severity-Emergency) | (int(facility-Kernel) << 3)

	if reflection.TimestampFieldIndex >= 0 {
		m.Timestamp = mv.Field(reflection.TimestampFieldIndex).Interface().(time.Time)
	} else {
		m.Timestamp = TimeNow().UTC()
	}

	if reflection.HostnameFieldIndex >= 0 {
		m.Hostname = mv.Field(reflection.HostnameFieldIndex).Interface().(string)
	} else {
		m.Hostname = defaultHostname
	}

	if reflection.AppNameFieldIndex >= 0 {
		m.AppName = mv.Field(reflection.AppNameFieldIndex).Interface().(string)
	} else {
		m.AppName = reflection.AppNameDefault
	}

	if reflection.ProcessIDFieldIndex >= 0 {
		m.ProcessID = mv.Field(reflection.ProcessIDFieldIndex).String()
	} else {
		m.ProcessID = defaultProcessID
	}

	if reflection.MessageIDFieldIndex >= 0 {
		m.MessageID = mv.Field(reflection.MessageIDFieldIndex).String()
	} else {
		m.MessageID = reflection.MessageIDDefault
	}

	for _, fieldReflection := range reflection.StructuredDataFieldReflections {
		v := mv.Field(fieldReflection.FieldIndex)
		if fieldReflection.OmitEmpty {
			zeroValue := reflect.Zero(mt.Field(fieldReflection.FieldIndex).Type)
			if zeroValue.Interface() == v.Interface() {
				continue
			}
		}
		m.AddDatum(fieldReflection.SdID, fieldReflection.FieldName, v.String())
	}

	if reflection.MessageFieldIndex >= 0 {
		m.Message = mv.Field(reflection.MessageFieldIndex).Interface().([]byte)
	}
	return &m
}

type Encoder struct {
	Writer io.Writer
}

func NewEncoder(w io.Writer) *Encoder {
	return &Encoder{Writer: w}
}

func (e Encoder) Encode(ob interface{}) error {
	m := Encode(ob)
	_, err := m.WriteTo(e.Writer)
	return err
}
