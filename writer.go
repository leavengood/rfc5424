package rfc5424

import (
	"github.com/secureworks/errset"
)

// MessageWriter is an interface for structures that can write RFC-5424 messages
type MessageWriter interface {
	WriteMessage(m Message) error
	Close() error
}

// MultiMessageWriter is a MessageWriter that acts like a fan-out, writing the
// provided message to each of the supplied MessageWriters in Writers.
type MultiMessageWriter struct {
	Writers []MessageWriter
}

// WriteMessage writes the message `m` to each of the fanout MessageWriters. If
// any of the writers fail, the others are still tried. Returns a non-nil error
// if any of the writers fail.
func (mmw MultiMessageWriter) WriteMessage(m Message) error {
	errCh := make(chan error, len(mmw.Writers))

	for _, w := range mmw.Writers {
		go func(w MessageWriter) {
			errCh <- w.WriteMessage(m)
		}(w)
	}

	errs := errset.ErrSet{}
	for _ = range mmw.Writers {
		err := <-errCh
		if err != nil {
			errs = append(errs, err)
		}
	}
	return errs.ReturnValue()
}

// Close closes each of the underlying writers.
func (mmw MultiMessageWriter) Close() error {
	errs := errset.ErrSet{}
	for _, w := range mmw.Writers {
		if err := w.Close(); err != nil {
			errs = append(errs, err)
		}
	}
	return errs.ReturnValue()
}
