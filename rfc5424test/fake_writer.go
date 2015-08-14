package rfc5424test

import (
	"github.com/secureworks/rfc5424"
)

// NewFakeWriter returns a fake writer.
func NewFakeWriter() *FakeWriter {
	return &FakeWriter{Messages: make(chan string, 4096)}
}

// FakeWriter implements rfc5424.MessageWriter for testing.
type FakeWriter struct {
	Messages chan string
	Error    error
}

// WriteMessage serializes and forwards the message to the Messages buffered
// channel. If Error is non-nil it returns Error rather than fake-sending the
// message.
func (fw *FakeWriter) WriteMessage(m rfc5424.Message) error {
	if fw.Error != nil {
		return fw.Error
	}

	str, err := m.MarshalBinary()
	if err != nil {
		return nil
	}
	fw.Messages <- string(str)
	return nil
}

// Close closes the Messages channel.
func (fw *FakeWriter) Close() error {
	close(fw.Messages)
	fw.Messages = nil
	return nil
}
