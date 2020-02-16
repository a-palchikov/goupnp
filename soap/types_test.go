package soap

import (
	"bytes"
	"encoding"
	"net/url"
	"reflect"
	"strings"
	"testing"
	"time"
)

type convTest interface {
	encoding.TextMarshaler
	encoding.TextUnmarshaler
}

type testCase struct {
	// value specifies the value for the Marshal test
	value convTest
	// str specifies the input string for Unmashal test
	str string
	// expected specifies the expected value
	expected convTest
	// err optionally specifies the expected error
	err string
}

func TestSerdeTypes(t *testing.T) {
	const time010203 time.Duration = (1*3600 + 2*60 + 3) * time.Second
	const time0102 time.Duration = (1*3600 + 2*60) * time.Second
	const time01 time.Duration = (1 * 3600) * time.Second
	const time235959 time.Duration = (23*3600 + 59*60 + 59) * time.Second

	var testCases = []testCase{
		{str: "", expected: ui1(0), err: `parsing "": invalid syntax`},       // d
		{str: " ", expected: ui1(0), err: `parsing " ": invalid syntax`},     // d
		{str: "abc", expected: ui1(0), err: `parsing "abc": invalid syntax`}, // d
		{str: "-1", expected: ui1(0), err: `parsing "-1": invalid syntax`},   // d
		{str: "0", expected: ui1(0)},
		{str: "1", expected: ui1(1)},
		{str: "255", expected: ui1(255)},
		{str: "255", value: ui1(255)},
		{str: "256", expected: ui1(0), err: `parsing "256": value out of range`},
		// ui2
		{str: "65535", expected: ui2(65535)},
		{str: "65536", expected: ui2(0), err: `parsing "65536": value out of range`},
		// ui4
		{str: "4294967295", expected: ui4(4294967295)},
		{str: "4294967296", expected: ui4(0), err: `parsing "4294967296": value out of range`},
		// ui8
		{str: "18446744073709551615", expected: ui8(18446744073709551615)},
		{str: "18446744073709551616", expected: ui8(0), err: `parsing "18446744073709551616": value out of range`},
		// i1
		{str: "", expected: i1(0), err: `parsing "": invalid syntax`},       // d
		{str: " ", expected: i1(0), err: `parsing " ": invalid syntax`},     // d
		{str: "abc", expected: i1(0), err: `parsing "abc": invalid syntax`}, // d
		{str: "0", expected: i1(0)},
		{str: "-1", expected: i1(-1)},
		{str: "127", expected: i1(127)},
		{str: "-128", expected: i1(-128)},
		{str: "128", expected: i1(0), err: `parsing "128": value out of range`},
		{str: "-129", expected: i1(0), err: `parsing "-129": value out of range`},
		// i2
		{str: "32767", expected: i2(32767)},
		{str: "-32768", expected: i2(-32768)},
		{str: "32768", expected: i2(0), err: `parsing "32768": value out of range`},
		{str: "-32769", expected: i2(0), err: `parsing "-32769": value out of range`},
		// i4
		{str: "2147483647", expected: i4(2147483647)},
		{str: "-2147483648", expected: i4(-2147483648)},
		{str: "2147483648", expected: i4(0), err: `parsing "2147483648": value out of range`},
		{str: "-2147483649", expected: i4(0), err: `parsing "-2147483649": value out of range`},
		// int
		{str: "9223372036854775807", expected: int_(9223372036854775807)},
		{str: "-9223372036854775808", expected: int_(-9223372036854775808)},
		{str: "9223372036854775808", expected: int_(0), err: `parsing "9223372036854775808": value out of range`},
		{str: "-9223372036854775809", expected: int_(0), err: `parsing "-9223372036854775809": value out of range`},
		// fixed14_4
		{str: "0.0000", expected: fixed14_4(0)},
		{str: "1.0000", expected: fixed14_4(1)},
		{str: "1.2346", expected: fixed14_4(1.23456)},
		{str: "-1.0000", expected: fixed14_4(-1)},
		{str: "-1.2346", expected: fixed14_4(-1.23456)},
		{str: "10000000000000.0000", expected: fixed14_4(1e13)},
		{str: "100000000000000.0000", expected: fixed14_4(1e14), err: `value "100000000000000.0000" out of bounds`},
		{str: "-10000000000000.0000", expected: fixed14_4(-1e13)},
		{str: "-100000000000000.0000", expected: fixed14_4(-1e14), err: `value "-100000000000000.0000" out of bounds`},
		// char
		{str: "a", expected: char('a')},
		{str: "z", expected: char('z')},
		{str: "\u1234", expected: char(0x1234)},
		{str: "aa", expected: char(0), err: `value "aa" is not a single rune`},
		{str: "", expected: char(0), err: "got empty string"},
		// date
		{str: "2013-10-08", expected: date(Date(time.Date(2013, 10, 8, 0, 0, 0, 0, localLoc)))},                                               // d
		{str: "20131008", expected: date(Date(time.Date(2013, 10, 8, 0, 0, 0, 0, localLoc)))},                                                 // d
		{str: "2013-10-08T10:30:50", expected: date(Date{}), err: `value "2013-10-08T10:30:50" is not in a recognized ISO8601 date format`},   // d
		{str: "2013-10-08T10:30:50Z", expected: date(Date{}), err: `value "2013-10-08T10:30:50Z" is not in a recognized ISO8601 date format`}, // d
		{str: "", expected: date(Date{}), err: `value "" is not in a recognized ISO8601 date format`},
		{str: "-1", expected: date(Date{}), err: `value "-1" is not in a recognized ISO8601 date format`},
		// time
		{str: "00:00:00", expected: timeOfDay(TimeOfDay{FromMidnight: 0})},
		{str: "000000", expected: timeOfDay(TimeOfDay{FromMidnight: 0})},
		{str: "24:00:00", expected: timeOfDay(TimeOfDay{FromMidnight: 24 * time.Hour})}, // ISO8601 special case
		{str: "24:01:00", expected: timeOfDay(TimeOfDay{}), err: `value "24:01:00" has value(s) out of range`},
		{str: "24:00:01", expected: timeOfDay(TimeOfDay{}), err: `value "24:00:01" has value(s) out of range`},
		{str: "25:00:00", expected: timeOfDay(TimeOfDay{}), err: `value "25:00:00" has value(s) out of range`},
		{str: "00:60:00", expected: timeOfDay(TimeOfDay{}), err: `value "00:60:00" has value(s) out of range`},
		{str: "00:00:60", expected: timeOfDay(TimeOfDay{}), err: `value "00:00:60" has value(s) out of range`},
		{str: "01:02:03", expected: timeOfDay(TimeOfDay{FromMidnight: time010203})},
		{str: "010203", expected: timeOfDay(TimeOfDay{FromMidnight: time010203})},
		{str: "23:59:59", expected: timeOfDay(TimeOfDay{FromMidnight: time235959})},
		{str: "235959", expected: timeOfDay(TimeOfDay{FromMidnight: time235959})},
		{str: "01:02", expected: timeOfDay(TimeOfDay{FromMidnight: time0102})},
		{str: "0102", expected: timeOfDay(TimeOfDay{FromMidnight: time0102})},
		{str: "01", expected: timeOfDay(TimeOfDay{FromMidnight: time01})},
		{str: "foo 01:02:03", expected: timeOfDay(TimeOfDay{}), err: `value "foo 01:02:03" is not in ISO8601 time format`},
		{str: "foo\n01:02:03", expected: timeOfDay(TimeOfDay{}), err: `value "foo\n01:02:03" is not in ISO8601 time format`},
		{str: "01:02:03 foo", expected: timeOfDay(TimeOfDay{}), err: `value "01:02:03 foo" is not in ISO8601 time format`},
		{str: "01:02:03\nfoo", expected: timeOfDay(TimeOfDay{}), err: `value "01:02:03\nfoo" is not in ISO8601 time format`},
		{str: "01:02:03Z", expected: timeOfDay(TimeOfDay{}), err: `value "01:02:03Z" contains unexpected timezone`},
		{str: "01:02:03+01", expected: timeOfDay(TimeOfDay{}), err: `value "01:02:03+01" contains unexpected timezone`},
		{str: "01:02:03+01:23", expected: timeOfDay(TimeOfDay{}), err: `value "01:02:03+01:23" contains unexpected timezone`},
		{str: "01:02:03+0123", expected: timeOfDay(TimeOfDay{}), err: `value "01:02:03+0123" contains unexpected timezone`},
		{str: "01:02:03-01", expected: timeOfDay(TimeOfDay{}), err: `value "01:02:03-01" contains unexpected timezone`},
		{str: "01:02:03-01:23", expected: timeOfDay(TimeOfDay{}), err: `value "01:02:03-01:23" contains unexpected timezone`},
		{str: "01:02:03-0123", expected: timeOfDay(TimeOfDay{}), err: `value "01:02:03-0123" contains unexpected timezone`},
		// time.tz
		{str: "24:00:01", expected: timeOfDayTz(TimeOfDayTz{}), err: `value "24:00:01" has value(s) out of range`},
		{str: "01Z", expected: timeOfDayTz(TimeOfDayTz{time01, true, 0})},
		{str: "01:02:03Z", expected: timeOfDayTz(TimeOfDayTz{time010203, true, 0})},
		{str: "01+01", expected: timeOfDayTz(TimeOfDayTz{time01, true, 3600})},
		{str: "01:02:03+01", expected: timeOfDayTz(TimeOfDayTz{time010203, true, 3600})},
		{str: "01:02:03+01:23", expected: timeOfDayTz(TimeOfDayTz{time010203, true, 3600 + 23*60})},
		{str: "01:02:03+0123", expected: timeOfDayTz(TimeOfDayTz{time010203, true, 3600 + 23*60})},
		{str: "01:02:03-01", expected: timeOfDayTz(TimeOfDayTz{time010203, true, -3600})},
		{str: "01:02:03-01:23", expected: timeOfDayTz(TimeOfDayTz{time010203, true, -(3600 + 23*60)})},
		{str: "01:02:03-0123", expected: timeOfDayTz(TimeOfDayTz{time010203, true, -(3600 + 23*60)})},
		// dateTime
		{str: "2013-10-08T00:00:00", expected: dateTime(DateTime(time.Date(2013, 10, 8, 0, 0, 0, 0, localLoc)))}, // d
		{str: "20131008", expected: dateTime(DateTime(time.Date(2013, 10, 8, 0, 0, 0, 0, localLoc)))},
		{str: "2013-10-08T10:30:50", expected: dateTime(DateTime(time.Date(2013, 10, 8, 10, 30, 50, 0, localLoc)))}, // d
		{str: "2013-10-08T10:30:50T", expected: dateTime(DateTime{}), err: `value "10:30:50T" is not in ISO8601 time format`},
		{str: "2013-10-08T10:30:50+01", expected: dateTime(DateTime{}), err: `unexpected timezone in "2013-10-08T10:30:50+01"`},       // d
		{str: "2013-10-08T10:30:50+01:23", expected: dateTime(DateTime{}), err: `unexpected timezone in "2013-10-08T10:30:50+01:23"`}, // d
		{str: "2013-10-08T10:30:50+0123", expected: dateTime(DateTime{}), err: `unexpected timezone in "2013-10-08T10:30:50+0123"`},   // d
		{str: "2013-10-08T10:30:50-01", expected: dateTime(DateTime{}), err: `unexpected timezone in "2013-10-08T10:30:50-01"`},       // d
		{str: "2013-10-08T10:30:50-01:23", expected: dateTime(DateTime{}), err: `unexpected timezone in "2013-10-08T10:30:50-01:23"`}, // d
		{str: "2013-10-08T10:30:50-0123", expected: dateTime(DateTime{}), err: `unexpected timezone in "2013-10-08T10:30:50-0123"`},   // d
		// dateTime.tz
		{str: "2013-10-08T10:30:50", expected: dateTimeTz(DateTimeTz(time.Date(2013, 10, 8, 10, 30, 50, 0, localLoc)))},
		{str: "2013-10-08T10:30:50+01", expected: dateTimeTz(DateTimeTz(time.Date(2013, 10, 8, 10, 30, 50, 0, time.FixedZone("+01:00", 3600))))},
		{str: "2013-10-08T10:30:50+01:23", expected: dateTimeTz(DateTimeTz(time.Date(2013, 10, 8, 10, 30, 50, 0, time.FixedZone("+01:23", 3600+23*60))))},
		{str: "2013-10-08T10:30:50+0123", expected: dateTimeTz(DateTimeTz(time.Date(2013, 10, 8, 10, 30, 50, 0, time.FixedZone("+01:23", 3600+23*60))))},
		{str: "2013-10-08T10:30:50-01", expected: dateTimeTz(DateTimeTz(time.Date(2013, 10, 8, 10, 30, 50, 0, time.FixedZone("-01:00", -3600))))},
		{str: "2013-10-08T10:30:50-01:23", expected: dateTimeTz(DateTimeTz(time.Date(2013, 10, 8, 10, 30, 50, 0, time.FixedZone("-01:23", -(3600+23*60)))))},
		{str: "2013-10-08T10:30:50-0123", expected: dateTimeTz(DateTimeTz(time.Date(2013, 10, 8, 10, 30, 50, 0, time.FixedZone("-01:23", -(3600+23*60)))))},
		// boolean
		{str: "0", expected: bool_(false)},
		{str: "1", expected: bool_(true)},
		{str: "false", expected: bool_(false)},
		{str: "true", expected: bool_(true)},
		{str: "no", expected: bool_(false)},
		{str: "yes", expected: bool_(true)},
		{str: "", expected: bool_(false), err: `"" is not a valid boolean value`},
		{str: "other", expected: bool_(false), err: `"other" is not a valid boolean value`},
		{str: "2", expected: bool_(false), err: `"2" is not a valid boolean value`},
		{str: "-1", expected: bool_(false), err: `"-1" is not a valid boolean value`},
		// bin.base64
		{str: "", expected: binBase64(nil)},
		{str: "YQ==", expected: binBase64(BinBase64("a"))},
		{str: "TG9uZ2VyIFN0cmluZy4=", expected: binBase64(BinBase64("Longer String."))},
		{str: "TG9uZ2VyIEFsaWduZWQu", expected: binBase64(BinBase64("Longer Aligned."))},
		// bin.hex
		{str: "", expected: binHex(nil)},
		{str: "61", expected: binHex(BinHex("a"))},
		{str: "4c6f6e67657220537472696e672e", expected: binHex(BinHex("Longer String."))},
		{str: "4C6F6E67657220537472696E672E", expected: binHex(BinHex("Longer String."))},
		// uri
		{str: "http://example.com/path", expected: uri(URI(url.URL{Scheme: "http", Host: "example.com", Path: "/path"}))},
	}

	for _, testCase := range testCases {
		var err error
		var out []byte
		switch {
		case testCase.value == nil:
			typ := reflect.TypeOf(testCase.expected).Elem()
			value := reflect.New(typ).Interface().(convTest)
			err = value.UnmarshalText([]byte(testCase.str))
		case testCase.value != nil:
			out, err = testCase.value.MarshalText()
			if !bytes.Equal(out, []byte(testCase.str)) {
				t.Errorf("expected %q but got %q", testCase.str, out)
			}
		}
		if err == nil && testCase.err != "" {
			t.Error("expected error:", testCase.err)
		}
		if err != nil {
			if testCase.err != "" && !strings.HasSuffix(err.Error(), testCase.err) {
				t.Errorf("expected error %v but got %v", err, testCase.err)
			} else if testCase.err == "" {
				t.Error("failed: ", err)
			}
		}
	}
}

func ui1(v Ui1) convTest                 { return &v }
func ui2(v Ui2) convTest                 { return &v }
func ui4(v Ui4) convTest                 { return &v }
func ui8(v Ui8) convTest                 { return &v }
func i1(v I1) convTest                   { return &v }
func i2(v I2) convTest                   { return &v }
func i4(v I4) convTest                   { return &v }
func int_(v Int) convTest                { return &v }
func fixed14_4(v Fixed14_4) convTest     { return &v }
func char(v Char) convTest               { return &v }
func bool_(v Bool) convTest              { return &v }
func date(v Date) convTest               { return &v }
func dateTime(v DateTime) convTest       { return &v }
func dateTimeTz(v DateTimeTz) convTest   { return &v }
func timeOfDay(v TimeOfDay) convTest     { return &v }
func timeOfDayTz(v TimeOfDayTz) convTest { return &v }
func binBase64(v BinBase64) convTest     { return &v }
func binHex(v BinHex) convTest           { return &v }
func uri(v URI) convTest                 { return &v }
