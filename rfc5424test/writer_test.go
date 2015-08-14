package rfc5424test

import (
	"errors"

	. "gopkg.in/check.v1"

	"github.com/secureworks/rfc5424"
)

var _ = Suite(&MultiMessageWriterTest{})

type MultiMessageWriterTest struct {
}

func (testSuite *MultiMessageWriterTest) TestStuff(c *C) {
	mmw := rfc5424.MultiMessageWriter{}

	// works with no writers
	err := mmw.WriteMessage(rfc5424.Message{MessageID: "one"})
	c.Assert(err, IsNil)

	// Writes to two writers when they both work
	fw1 := NewFakeWriter()
	fw2 := NewFakeWriter()
	mmw.Writers = append(mmw.Writers, fw1)
	mmw.Writers = append(mmw.Writers, fw2)
	err = mmw.WriteMessage(rfc5424.Message{MessageID: "one"})
	c.Assert(err, IsNil)
	m := <-fw1.Messages
	c.Assert(m, Equals, "<0>1 0001-01-01T00:00:00Z - - - one -")
	m = <-fw2.Messages
	c.Assert(m, Equals, "<0>1 0001-01-01T00:00:00Z - - - one -")

	// Writes to one and the other fails
	fw1.Error = errors.New("Couldn't frob the grob")
	err = mmw.WriteMessage(rfc5424.Message{MessageID: "one"})
	c.Assert(err, ErrorMatches, "Couldn't frob the grob")
	m = <-fw2.Messages
	c.Assert(m, Equals, "<0>1 0001-01-01T00:00:00Z - - - one -")

	// Closes upstream on close
	mmw.Close()
	c.Assert(fw1.Messages, IsNil)
	c.Assert(fw2.Messages, IsNil)
}
