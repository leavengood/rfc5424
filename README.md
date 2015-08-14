
[![Build Status](https://travis-ci.org/secureworks/rfc5424.png)](https://travis-ci.org/secureworks/rfc5424)

[![](https://godoc.org/github.com/secureworks/rfc5424?status.png)](http://godoc.org/github.com/secureworks/rfc5424)

This is a Go library for encoding and decoding RFC-5424 Syslog messages in the style of `encoding/json` or `encoding/xml` from the standard library.

[It is very much a work in progress. Please don't use it at all yet!]

Encoding concept:

```
type MyError struct {
    Severity rfc5424.Severity   `log:"error"`
    Facility int                `log:"kern"`
    Timestamp time.Time
    SessionID string            `log:"5516@xyz sessionID"`
    HumanReadableMessage string `log:",message"`
}

myError := MyError{SessionID: "1234", HumanReadableMessage: "Invalid Frob"}
logSocket, err := net.Dial("logserver:514")
err := rfc5424.Encoder(logSocket).Encode(&myError)
```

Decoding concept:

```
d := rfc5424.Decoder(r)
for {
    myError := MyError{}
    m, err := d.Decode(&myError)
    if err != nil {
        break
    }
}
```

TODO: 
 - check types in Reflect
 - require annotations for all special fields. don't use magic names
 - basic types in SD
 - complex types in SD
 - be clear about RFC violating long names
 - be clear about stream mode (with a length) vs. non-stream mode
