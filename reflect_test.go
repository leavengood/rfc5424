package rfc5424

import (
	"reflect"
	"time"

	. "gopkg.in/check.v1"
)

var _ = Suite(&ReflectTest{})

type ReflectTest struct {
}

type struct1 struct {
	Severity                Severity `log:"error"`
	Facility                Facility `log:"local2"`
	Timestamp               time.Time
	Hostname                string
	AppName                 string `log:"myAppName"`
	ProcessID               int
	MessageID               string
	Message                 []byte
	MyCustomInt             int
	MyCustomString          string
	MyCustomBool            bool
	myUnexportedValue       string
	myUnexportedTaggedValue string `log:"myUnexportedTaggedValue"`
}

var expectedReflection1 = reflection{
	Type:                reflect.TypeOf(struct1{}),
	SeverityFieldIndex:  0,
	SeverityDefault:     Error,
	FacilityFieldIndex:  1,
	FacilityDefault:     Local2,
	TimestampFieldIndex: 2,
	HostnameFieldIndex:  3,
	AppNameFieldIndex:   4,
	AppNameDefault:      "myAppName",
	ProcessIDFieldIndex: 5,
	MessageIDFieldIndex: 6,
	MessageIDDefault:    "struct1",
	MessageFieldIndex:   7,
	StructuredDataFieldReflections: []structuredDataFieldReflection{
		structuredDataFieldReflection{
			FieldIndex: 8,
			FieldName:  "MyCustomInt",
			SdID:       "0@local",
		},
		structuredDataFieldReflection{
			FieldIndex: 9,
			FieldName:  "MyCustomString",
			SdID:       "0@local",
		},
		structuredDataFieldReflection{
			FieldIndex: 10,
			FieldName:  "MyCustomBool",
			SdID:       "0@local",
		},
		structuredDataFieldReflection{
			FieldIndex: 12,
			FieldName:  "myUnexportedTaggedValue",
			SdID:       "0@local",
		},
	},
}

type struct2 struct {
	MyCustomInt             int
	MyCustomString          string
	MyCustomBool            bool
	myUnexportedValue       string
	myUnexportedTaggedValue string `log:"5516@sbc myUnexportedTaggedValue"`
}

var expectedReflection2 = reflection{
	Type:                reflect.TypeOf(struct2{}),
	SeverityFieldIndex:  -1,
	SeverityDefault:     Info,
	FacilityFieldIndex:  -1,
	FacilityDefault:     Local0,
	TimestampFieldIndex: -1,
	HostnameFieldIndex:  -1,
	AppNameFieldIndex:   -1,
	AppNameDefault:      "rfc5424.test",
	ProcessIDFieldIndex: -1,
	MessageIDFieldIndex: -1,
	MessageIDDefault:    "struct2",
	MessageFieldIndex:   -1,
	StructuredDataFieldReflections: []structuredDataFieldReflection{
		structuredDataFieldReflection{
			FieldIndex: 0,
			FieldName:  "MyCustomInt",
			SdID:       "0@local",
		},
		structuredDataFieldReflection{
			FieldIndex: 1,
			FieldName:  "MyCustomString",
			SdID:       "0@local",
		},
		structuredDataFieldReflection{
			FieldIndex: 2,
			FieldName:  "MyCustomBool",
			SdID:       "0@local",
		},
		structuredDataFieldReflection{
			FieldIndex: 4,
			FieldName:  "myUnexportedTaggedValue",
			SdID:       "5516@sbc",
		},
	},
}

type struct3 struct {
	Severity       Severity
	Facility       Facility
	Timestamp      time.Time
	Hostname       string
	AppName        string
	ProcessID      int
	MessageID      string
	Message        []byte // To test the tag below
	RealMessage    []byte `log:",message"`
	MyCustomString string
}

var expectedReflection3 = reflection{
	Type:                reflect.TypeOf(struct3{}),
	SeverityFieldIndex:  0,
	SeverityDefault:     Info,
	FacilityFieldIndex:  1,
	FacilityDefault:     Local0,
	TimestampFieldIndex: 2,
	HostnameFieldIndex:  3,
	AppNameFieldIndex:   4,
	AppNameDefault:      "rfc5424.test",
	ProcessIDFieldIndex: 5,
	MessageIDFieldIndex: 6,
	MessageIDDefault:    "struct3",
	MessageFieldIndex:   8,
	StructuredDataFieldReflections: []structuredDataFieldReflection{
		structuredDataFieldReflection{
			FieldIndex: 9,
			FieldName:  "MyCustomString",
			SdID:       "0@local",
		},
	},
}

func (s *ReflectTest) TestCanReflect(c *C) {
	{
		v := struct1{}
		r := Reflect(reflect.TypeOf(v))
		c.Assert(r, DeepEquals, &expectedReflection1)
	}
	{
		v := struct2{}
		r := Reflect(reflect.TypeOf(v))
		c.Assert(r, DeepEquals, &expectedReflection2)
	}

	{
		v := struct3{}
		r := Reflect(reflect.TypeOf(v))
		c.Assert(r, DeepEquals, &expectedReflection3)
	}

}
