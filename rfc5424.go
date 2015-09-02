package rfc5424

const severityMask = 0x07
const facilityMask = 0xf8

type Severity int

const (
	DefaultSeverity = iota
	Emergency
	Alert
	Critical
	Error
	Warning
	Notice
	Info
	Debug
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
	Kernel
	User
	Mail
	Daemon
	Auth
	Syslog
	LPR
	News
	UUCP
	Clock
	AuthPriv
	FTP
	NTP
	Audit
	LogAlert
	Cron
	Local0
	Local1
	Local2
	Local3
	Local4
	Local5
	Local6
	Local7
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
