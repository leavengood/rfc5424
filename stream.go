package rfc5424

import (
	"fmt"
	"io"
	"io/ioutil"
)

// WriteTo writes the message to a stream of messages in the style defined
// by RFC-5425. (It does not implement the TLS stuff described in the RFC, just
// the length delimiting.
func (m Message) WriteTo(w io.Writer) (int64, error) {
	b, err := m.MarshalBinary()
	if err != nil {
		return 0, err
	}
	n, err := fmt.Fprintf(w, "%d %s", len(b), b)
	return int64(n), err
}

// ReadFrom reads a single record from an RFC-5425 style stream of messages
func (m *Message) ReadFrom(r io.Reader) (int64, error) {
	var length int64
	n1, err := fmt.Fscanf(r, "%d ", &length)
	if err != nil {
		return 0, err
	}
	r2 := io.LimitReader(r, length)
	buf, err := ioutil.ReadAll(r2)
	if err != nil {
		return int64(n1 + len(buf)), err
	}
	if len(buf) != int(length) {
		return int64(n1 + len(buf)), fmt.Errorf("Expected to read %d bytes, got %d", length, len(buf))
	}
	err = m.UnmarshalBinary(buf)
	if err != nil {
		return 0, err
	}
	return int64(n1 + len(buf)), err
}
