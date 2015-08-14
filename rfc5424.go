package rfc5424

import "time"

const severityMask = 0x07
const facilityMask = 0xf8

// Message represents a log message as defined by RFC-5424
// (https://tools.ietf.org/html/rfc5424)
type Message struct {
	Severity  int       `json:"severity,omitempty"`
	Facility  int       `json:"facility,omitempty"`
	Timestamp time.Time `json:"timestamp,omitempty"`
	Hostname  string    `json:"hostname,omitempty"`
	AppName   string    `json:"appname,omitempty"`
	ProcessID string    `json:"process_id,omitempty"`
	MessageID string    `json:"message_id,omitempty"`

	StructuredData []StructuredData `json:"structured_data,omitempty"`
	Message        []byte           `json:"message,omitempty"`
}

// SDParam represents parameters for structured data
type SDParam struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

// StructuredData represents structured data within a log message
type StructuredData struct {
	ID         string    `json:"id"`
	Parameters []SDParam `json:"parameters"`
}

// AddParam adds the given name and value to this structured data's parameters
func (sd *StructuredData) AddParam(name, value string) {
	sd.Parameters = append(sd.Parameters, SDParam{Name: name, Value: value})
}

// AddDatum adds structured data to a log message
func (m *Message) AddDatum(ID string, Name string, Value string) {
	if m.StructuredData == nil {
		m.StructuredData = []StructuredData{}
	}
	for i, sd := range m.StructuredData {
		if sd.ID == ID {
			sd.Parameters = append(sd.Parameters, SDParam{Name: Name, Value: Value})
			m.StructuredData[i] = sd
			return
		}
	}

	m.StructuredData = append(m.StructuredData, StructuredData{
		ID: ID,
		Parameters: []SDParam{
			SDParam{
				Name:  Name,
				Value: Value,
			},
		},
	})
}
