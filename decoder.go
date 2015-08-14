package rfc5424

import (
	"io"
	"reflect"
)

type Decoder struct {
	Reader io.Reader
}

func NewDecoder(r io.Reader) *Decoder {
	return &Decoder{Reader: r}
}

func (d Decoder) Decode(ob interface{}) error {
	m := Message{}
	if _, err := m.ReadFrom(d.Reader); err != nil {
		return err
	}
	return d.decode(&m, ob)
}

func (d Decoder) decode(m *Message, ob interface{}) error {
	mt := reflect.TypeOf(ob)
	mv := reflect.ValueOf(ob)
	reflection := Reflect(mt)

	if reflection.SeverityFieldIndex >= 0 {
		severity := Emergency + (m.Priority & severityMask)
		mv.Field(reflection.SeverityFieldIndex).Set(reflect.ValueOf(severity))
	}
	if reflection.FacilityFieldIndex >= 0 {
		facility := Kernel + ((m.Priority & facilityMask) >> 3)
		mv.Field(reflection.FacilityFieldIndex).Set(reflect.ValueOf(facility))
	}
	if reflection.TimestampFieldIndex >= 0 {
		mv.Field(reflection.TimestampFieldIndex).Set(reflect.ValueOf(m.Timestamp))
	}
	if reflection.HostnameFieldIndex >= 0 {
		mv.Field(reflection.HostnameFieldIndex).Set(reflect.ValueOf(m.Hostname))
	}
	if reflection.AppNameFieldIndex >= 0 {
		mv.Field(reflection.AppNameFieldIndex).Set(reflect.ValueOf(m.AppName))
	}
	if reflection.ProcessIDFieldIndex >= 0 {
		mv.Field(reflection.ProcessIDFieldIndex).Set(reflect.ValueOf(m.ProcessID))
	}
	if reflection.MessageIDFieldIndex >= 0 {
		mv.Field(reflection.MessageIDFieldIndex).Set(reflect.ValueOf(m.MessageID))
	}
	if reflection.MessageFieldIndex >= 0 {
		mv.Field(reflection.MessageIDFieldIndex).Set(reflect.ValueOf(m.Message))
	}
	for _, sd := range m.StructuredData {
		for _, param := range sd.Parameters {
			fieldReflection := reflection.GetStructuredDataFieldReflection(
				sd.ID, param.Name)
			if fieldReflection != nil {
				continue
			}
			// XXX: handle other types than strings
			mv.Field(fieldReflection.FieldIndex).SetString(param.Value)
		}
	}
	return nil
}
