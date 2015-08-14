package rfc5424

import (
	"io"
	"reflect"
	"time"
)

var TimeNow = time.Now

type Encoder struct {
	Writer io.Writer
}

func NewEncoder(w io.Writer) *Encoder {
	return &Encoder{Writer: w}
}

type Severity int

const (
	DefaultSeverity = iota
	Emergency       = iota
	Alert           = iota
	Critical        = iota
	Error           = iota
	Warning         = iota
	Notice          = iota
	Info            = iota
	Debug           = iota
)

var severityNames = map[string]Severity{
	"emergency":     Emergency,
	"emerg":         Emergency,
	"alert":         Alert,
	"critical":      Critical,
	"crit":          Critical,
	"error":         Error,
	"warning":       Warning,
	"warn":          Warning,
	"notice":        Notice,
	"informational": Info,
	"info":          Info,
	"debug":         Debug,
}

type Facility int

const (
	DefaultFacility = iota
	Kernel          = iota
	User            = iota
	Mail            = iota
	Daemon          = iota
	Auth            = iota
	Syslog          = iota
	LPR             = iota
	News            = iota
	UUCP            = iota
	Clock           = iota
	AuthPriv        = iota
	FTP             = iota
	NTP             = iota
	Audit           = iota
	LogAlert        = iota
	Cron            = iota
	Local0          = iota
	Local1          = iota
	Local2          = iota
	Local3          = iota
	Local4          = iota
	Local5          = iota
	Local6          = iota
	Local7          = iota
)

var facilityNames = map[string]Facility{
	"kernel":   Kernel,
	"user":     User,
	"mail":     Mail,
	"daemon":   Daemon,
	"auth":     Auth,
	"syslog":   Syslog,
	"lpr":      LPR,
	"news":     News,
	"uucp":     UUCP,
	"clock":    Clock,
	"authpriv": AuthPriv,
	"ftp":      FTP,
	"ntp":      NTP,
	"audit":    Audit,
	"logalert": LogAlert,
	"cron":     Cron,
	"local0":   Local0,
	"local1":   Local1,
	"local2":   Local2,
	"local3":   Local3,
	"local4":   Local4,
	"local5":   Local5,
	"local6":   Local6,
	"local7":   Local7,
}

func (e Encoder) encode(ob interface{}) *Message {
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

func (e Encoder) Encode(ob interface{}) error {
	m := e.encode(ob)
	_, err := m.WriteTo(e.Writer)
	return err
}
