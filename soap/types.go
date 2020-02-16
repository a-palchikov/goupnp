package soap

import (
	"bytes"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"net/url"
	"regexp"
	"strconv"
	"time"
	"unicode/utf8"

	"github.com/pkg/errors"
)

var (
	// localLoc acts like time.Local for this package, but is faked out by the
	// unit tests to ensure that things stay constant (especially when running
	// this test in a place where local time is UTC which might mask bugs).
	localLoc = time.Local
)

// UPnP SOAP data types
// Reference: http://upnp.org/specs/arch/UPnP-arch-DeviceArchitecture-v2.0.pdf

// Ui1 represents a SOAP "ui1" type
type Ui1 uint8

func (r Ui1) MarshalText() ([]byte, error) {
	return []byte(strconv.FormatUint(uint64(r), 10)), nil
}

func (r *Ui1) UnmarshalText(text []byte) error {
	v, err := strconv.ParseUint(string(text), 10, 8)
	if err != nil {
		return err
	}
	*r = Ui1(v)
	return nil
}

// Ui2 represents the SOAP "ui2" type
type Ui2 uint16

func (r Ui2) MarshalText() ([]byte, error) {
	return []byte(strconv.FormatUint(uint64(r), 10)), nil
}

func (r *Ui2) UnmarshalText(text []byte) error {
	v, err := strconv.ParseUint(string(text), 10, 16)
	if err != nil {
		return err
	}
	*r = Ui2(v)
	return nil
}

// Ui4 represents the SOAP "ui4" type
type Ui4 uint32

func (r Ui4) MarshalText() ([]byte, error) {
	return []byte(strconv.FormatUint(uint64(r), 10)), nil
}

func (r *Ui4) UnmarshalText(text []byte) error {
	v, err := strconv.ParseUint(string(text), 10, 32)
	if err != nil {
		return err
	}
	*r = Ui4(v)
	return nil
}

// Ui8 represents the SOAP "ui8" type
type Ui8 uint64

func (r Ui8) MarshalText() ([]byte, error) {
	return []byte(strconv.FormatUint(uint64(r), 10)), nil
}

func (r *Ui8) UnmarshalText(text []byte) error {
	v, err := strconv.ParseUint(string(text), 10, 64)
	if err != nil {
		return err
	}
	*r = Ui8(v)
	return nil
}

// I1 represents the SOAP "i1" type
type I1 int8

func (r I1) MarshalText() ([]byte, error) {
	return []byte(strconv.FormatInt(int64(r), 10)), nil
}

func (r *I1) UnmarshalText(text []byte) error {
	v, err := strconv.ParseInt(string(text), 10, 8)
	if err != nil {
		return err
	}
	*r = I1(v)
	return nil
}

// I2 represents the SOAP "i2" type
type I2 int16

func (r I2) MarshalText() ([]byte, error) {
	return []byte(strconv.FormatInt(int64(r), 10)), nil
}

func (r *I2) UnmarshalText(text []byte) error {
	v, err := strconv.ParseInt(string(text), 10, 16)
	if err != nil {
		return err
	}
	*r = I2(v)
	return nil
}

// I4 represents the SOAP "i4" type
type I4 int32

func (r I4) MarshalText() ([]byte, error) {
	return []byte(strconv.FormatInt(int64(r), 10)), nil
}

func (r *I4) UnmarshalText(text []byte) error {
	v, err := strconv.ParseInt(string(text), 10, 32)
	if err != nil {
		return err
	}
	*r = I4(v)
	return nil
}

// Int represents the SOAP "int" type
type Int int64

func (r Int) MarshalText() ([]byte, error) {
	return []byte(strconv.FormatInt(int64(r), 10)), nil
}

func (r *Int) UnmarshalText(text []byte) error {
	v, err := strconv.ParseInt(string(text), 10, 64)
	if err != nil {
		return err
	}
	*r = Int(v)
	return nil
}

// R4 represents the SOAP "r4" type
type R4 float32

func (r R4) MarshalText() ([]byte, error) {
	return []byte(strconv.FormatFloat(float64(r), 'G', -1, 32)), nil
}

func (r *R4) UnmarshalText(text []byte) error {
	v, err := strconv.ParseFloat(string(text), 32)
	if err != nil {
		return err
	}
	*r = R4(v)
	return nil
}

// R8 represents the SOAP "r8" float
type R8 float64

func (r R8) MarshalText() ([]byte, error) {
	return []byte(strconv.FormatFloat(float64(r), 'G', -1, 64)), nil
}

func (r *R8) UnmarshalText(text []byte) error {
	v, err := strconv.ParseFloat(string(text), 64)
	if err != nil {
		return err
	}
	*r = R8(v)
	return nil
}

// Fixed14_4 represents the SOAP "fixed.14.4" type
type Fixed14_4 float64

// MarshalText marshals this to SOAP "fixed.14.4" type.
func (r Fixed14_4) MarshalText() ([]byte, error) {
	if r >= 1e14 || r <= -1e14 {
		return nil, fmt.Errorf("soap fixed14_4: value %v out of bounds", r)
	}
	return []byte(strconv.FormatFloat(float64(r), 'f', 4, 64)), nil
}

// UnmarshalFixed14_4 unmarshals float64 from SOAP "fixed.14.4" type.
func (r *Fixed14_4) UnmarshalText(text []byte) error {
	v, err := strconv.ParseFloat(string(text), 64)
	if err != nil {
		return err
	}
	if v >= 1e14 || v <= -1e14 {
		return errors.Errorf("soap fixed14_4: value %q out of bounds", text)
	}
	*r = Fixed14_4(v)
	return nil
}

// Char represents the SOAP "char" type
type Char rune

// MarshalText marshals this to SOAP "char" type.
func (r Char) MarshalText() ([]byte, error) {
	if r == 0 {
		return nil, errors.New("soap char: rune 0 is not allowed")
	}
	return []byte(string(r)), nil
}

// UnmarshalChar unmarshals rune from SOAP "char" type.
func (r *Char) UnmarshalText(text []byte) error {
	if len(text) == 0 {
		return errors.New("soap char: got empty string")
	}
	rn, n := utf8.DecodeRune(text)
	if n != len(text) {
		return errors.Errorf("soap char: value %q is not a single rune", text)
	}
	*r = Char(rn)
	return nil
}

// String represents the SOAP "string" type
type String string

// MarshalText returns this string as a byte slice
func (r String) MarshalText() ([]byte, error) {
	return []byte(r), nil
}

// UnmarshalText interprets text as string
func (r *String) UnmarshalText(text []byte) error {
	*r = String(text)
	return nil
}

// Date represents a SOAP "date" type
type Date time.Time

// MarshalText marshals time.Time to SOAP "date" type. Note that this converts
// to local time, and discards the time-of-day components.
func (r Date) MarshalText() ([]byte, error) {
	return []byte((time.Time)(r).In(localLoc).Format("2006-01-02")), nil
}

// UnmarshalText unmarshals time.Time from SOAP "date" type. This outputs the
// date as midnight in the local time zone.
func (r *Date) UnmarshalText(text []byte) error {
	year, month, day, err := parseDateParts(string(text))
	if err != nil {
		return err
	}
	*r = Date(time.Date(year, time.Month(month), day, 0, 0, 0, 0, localLoc))
	return nil
}

// TimeOfDay implements the SOAP "time" or "time.tz" types
type TimeOfDay struct {
	// Duration of time since midnight.
	FromMidnight time.Duration

	// Set to true if Offset is specified. If false, then the timezone is
	// unspecified (and by ISO8601 - implies some "local" time).
	HasOffset bool

	// Offset is non-zero only if time.tz is used. It is otherwise ignored. If
	// non-zero, then it is regarded as a UTC offset in seconds. Note that the
	// sub-minutes is ignored by the marshal function.
	Offset int
}

// MarshalText marshals TimeOfDay to the "time" type.
func (r TimeOfDay) MarshalText() ([]byte, error) {
	d := int64(r.FromMidnight / time.Second)
	hour := d / 3600
	d = d % 3600
	minute := d / 60
	second := d % 60

	return []byte(fmt.Sprintf("%02d:%02d:%02d", hour, minute, second)), nil
}

// UnmarshalText unmarshals TimeOfDay from the "time" type.
func (r *TimeOfDay) UnmarshalText(text []byte) error {
	var t TimeOfDayTz
	if err := t.UnmarshalText(text); err != nil {
		return err
	} else if t.HasOffset {
		return errors.Errorf("soap time: value %q contains unexpected timezone", text)
	}
	*r = TimeOfDay(t)
	return nil
}

// TimeOfDayTz represents a SOAP "time.tz" type
type TimeOfDayTz TimeOfDay

// MarshalTimeOfDayTz marshals TimeOfDay to the "time.tz" type.
func (r TimeOfDayTz) MarshalText() ([]byte, error) {
	d := int64(r.FromMidnight / time.Second)
	hour := d / 3600
	d = d % 3600
	minute := d / 60
	second := d % 60

	tz := ""
	if r.HasOffset {
		if r.Offset == 0 {
			tz = "Z"
		} else {
			offsetMins := r.Offset / 60
			sign := '+'
			if offsetMins < 1 {
				offsetMins = -offsetMins
				sign = '-'
			}
			tz = fmt.Sprintf("%c%02d:%02d", sign, offsetMins/60, offsetMins%60)
		}
	}

	return []byte(fmt.Sprintf("%02d:%02d:%02d%s", hour, minute, second, tz)), nil
}

// UnmarshalText unmarshals TimeOfDay from the "time.tz" type.
func (r *TimeOfDayTz) UnmarshalText(text []byte) (err error) {
	zoneIndex := bytes.IndexAny(text, "Z+-")
	var timePart string
	var hasOffset bool
	var offset int
	if zoneIndex == -1 {
		hasOffset = false
		timePart = string(text)
	} else {
		hasOffset = true
		timePart = string(text[:zoneIndex])
		if offset, err = parseTimezone(string(text[zoneIndex:])); err != nil {
			return err
		}
	}

	hour, minute, second, err := parseTimeParts(timePart)
	if err != nil {
		return err
	}

	fromMidnight := time.Duration(hour*3600+minute*60+second) * time.Second

	// ISO8601 special case - values up to 24:00:00 are allowed, so using
	// strictly greater-than for the maximum value.
	if fromMidnight > 24*time.Hour || minute >= 60 || second >= 60 {
		return errors.Errorf("soap time.tz: value %q has value(s) out of range", text)
	}

	*r = TimeOfDayTz{
		FromMidnight: time.Duration(hour*3600+minute*60+second) * time.Second,
		HasOffset:    hasOffset,
		Offset:       offset,
	}
	return nil
}

// DateTime represents the SOAP "dateTime" type
type DateTime time.Time

// MarshalText marshals time.Time to SOAP "dateTime" type. Note that this
// converts to local time.
func (r DateTime) MarshalText() ([]byte, error) {
	return []byte((time.Time)(r).In(localLoc).Format("2006-01-02T15:04:05")), nil
}

// UnmarshalText unmarshals time.Time from the SOAP "dateTime" type. This
// returns a value in the local timezone.
func (r *DateTime) UnmarshalText(text []byte) (err error) {
	dateStr, timeStr, zoneStr, err := splitCompleteDateTimeZone(string(text))
	if err != nil {
		return err
	}

	if len(zoneStr) != 0 {
		return errors.Errorf("soap datetime: unexpected timezone in %q", text)
	}

	year, month, day, err := parseDateParts(dateStr)
	if err != nil {
		return err
	}

	var hour, minute, second int
	if len(timeStr) != 0 {
		hour, minute, second, err = parseTimeParts(timeStr)
		if err != nil {
			return err
		}
	}

	*r = DateTime(time.Date(year, time.Month(month), day, hour, minute, second, 0, localLoc))
	return nil
}

// DateTimeTz represents the SOAP "dateTime.tz" type
type DateTimeTz time.Time

// MarshalText marshals time.Time to SOAP "dateTime.tz" type.
func (r DateTimeTz) MarshalText() ([]byte, error) {
	return []byte((time.Time)(r).Format("2006-01-02T15:04:05-07:00")), nil
}

// UnmarshalText unmarshals time.Time from the SOAP "dateTime.tz" type.
// This returns a value in the local timezone when the timezone is unspecified.
func (r *DateTimeTz) UnmarshalText(text []byte) (err error) {
	dateStr, timeStr, zoneStr, err := splitCompleteDateTimeZone(string(text))
	if err != nil {
		return err
	}

	year, month, day, err := parseDateParts(dateStr)
	if err != nil {
		return err
	}

	var hour, minute, second int
	var location *time.Location = localLoc
	if len(timeStr) != 0 {
		hour, minute, second, err = parseTimeParts(timeStr)
		if err != nil {
			return
		}
		if len(zoneStr) != 0 {
			var offset int
			offset, err = parseTimezone(zoneStr)
			if err != nil {
				return errors.Wrap(err, "parsing timezone")
			}
			if offset == 0 {
				location = time.UTC
			} else {
				location = time.FixedZone("", offset)
			}
		}
	}

	*r = DateTimeTz(time.Date(year, time.Month(month), day, hour, minute, second, 0, location))
	return nil
}

// Bool represents the SOAP "boolean" type
type Bool bool

// MarshalText marshals bool to SOAP "boolean" type.
func (r Bool) MarshalText() ([]byte, error) {
	if r {
		return []byte("1"), nil
	}
	return []byte("0"), nil
}

// UnmarshalText unmarshals bool from the SOAP "boolean" type.
func (r *Bool) UnmarshalText(text []byte) error {
	switch string(text) {
	case "0", "false", "no":
		*r = Bool(false)
		return nil
	case "1", "true", "yes":
		*r = Bool(true)
		return nil
	}
	return errors.Errorf("soap boolean: %q is not a valid boolean value", text)
}

// BinBase64 represents the SOAP "bin.base64" type
type BinBase64 []byte

// MarshalText marshals []byte to SOAP "bin.base64" type.
func (r BinBase64) MarshalText() ([]byte, error) {
	buf := make([]byte, base64.StdEncoding.EncodedLen(len(r)))
	base64.StdEncoding.Encode(buf, r)
	return buf, nil
}

// UnmarshalText unmarshals []byte from the SOAP "bin.base64" type.
func (r *BinBase64) UnmarshalText(text []byte) error {
	buf := make([]byte, base64.StdEncoding.DecodedLen(len(text)))
	//nolint:errcheck
	base64.StdEncoding.Decode(buf, text)
	*r = BinBase64(buf)
	return nil
}

// BinHex represents the SOAP "bin.hex" type
type BinHex []byte

// MarshalBinHex marshals []byte to SOAP "bin.hex" type.
func (r BinHex) MarshalText() ([]byte, error) {
	buf := make([]byte, hex.EncodedLen(len(r)))
	hex.Encode(buf, r)
	return buf, nil
}

// UnmarshalText unmarshals []byte from the SOAP "bin.hex" type.
func (r *BinHex) UnmarshalText(text []byte) error {
	buf := make([]byte, hex.DecodedLen(len(text)))
	//nolint:errcheck
	hex.Decode(buf, text)
	*r = BinHex(buf)
	return nil
}

// URI represents the SOAP "uri" type
type URI url.URL

// MarshalText marshals *url.URL to SOAP "uri" type.
func (r URI) MarshalText() ([]byte, error) {
	return []byte((*url.URL)(&r).String()), nil
}

// UnmarshalText unmarshals *url.URL from the SOAP "uri" type.
func (r *URI) UnmarshalText(text []byte) error {
	u, err := url.Parse(string(text))
	if err != nil {
		return err
	}
	*r = URI(*u)
	return nil
}

func parseInt(s string, err *error) int {
	v, parseErr := strconv.ParseInt(s, 10, 64)
	if parseErr != nil {
		*err = parseErr
	}
	return int(v)
}

var dateRegexps = []*regexp.Regexp{
	// yyyy[-mm[-dd]]
	regexp.MustCompile(`^(\d{4})(?:-(\d{2})(?:-(\d{2}))?)?$`),
	// yyyy[mm[dd]]
	regexp.MustCompile(`^(\d{4})(?:(\d{2})(?:(\d{2}))?)?$`),
}

func parseDateParts(s string) (year, month, day int, err error) {
	var parts []string
	for _, re := range dateRegexps {
		parts = re.FindStringSubmatch(s)
		if parts != nil {
			break
		}
	}
	if parts == nil {
		err = fmt.Errorf("soap date: value %q is not in a recognized ISO8601 date format", s)
		return
	}

	year = parseInt(parts[1], &err)
	month = 1
	day = 1
	if len(parts[2]) != 0 {
		month = parseInt(parts[2], &err)
		if len(parts[3]) != 0 {
			day = parseInt(parts[3], &err)
		}
	}

	if err != nil {
		err = fmt.Errorf("soap date: %q: %v", s, err)
	}

	return
}

var timeRegexps = []*regexp.Regexp{
	// hh[:mm[:ss]]
	regexp.MustCompile(`^(\d{2})(?::(\d{2})(?::(\d{2}))?)?$`),
	// hh[mm[ss]]
	regexp.MustCompile(`^(\d{2})(?:(\d{2})(?:(\d{2}))?)?$`),
}

func parseTimeParts(s string) (hour, minute, second int, err error) {
	var parts []string
	for _, re := range timeRegexps {
		parts = re.FindStringSubmatch(s)
		if parts != nil {
			break
		}
	}
	if parts == nil {
		err = fmt.Errorf("soap time: value %q is not in ISO8601 time format", s)
		return
	}

	hour = parseInt(parts[1], &err)
	if len(parts[2]) != 0 {
		minute = parseInt(parts[2], &err)
		if len(parts[3]) != 0 {
			second = parseInt(parts[3], &err)
		}
	}

	if err != nil {
		err = fmt.Errorf("soap time: %q: %v", s, err)
	}

	return
}

// (+|-)hh[[:]mm]
var timezoneRegexp = regexp.MustCompile(`^([+-])(\d{2})(?::?(\d{2}))?$`)

func parseTimezone(s string) (offset int, err error) {
	if s == "Z" {
		return 0, nil
	}
	parts := timezoneRegexp.FindStringSubmatch(s)
	if parts == nil {
		err = fmt.Errorf("soap timezone: value %q is not in ISO8601 timezone format", s)
		return
	}

	offset = parseInt(parts[2], &err) * 3600
	if len(parts[3]) != 0 {
		offset += parseInt(parts[3], &err) * 60
	}
	if parts[1] == "-" {
		offset = -offset
	}

	if err != nil {
		err = fmt.Errorf("soap timezone: %q: %v", s, err)
	}

	return
}

var completeDateTimeZoneRegexp = regexp.MustCompile(`^([^T]+)(?:T([^-+Z]+)(.+)?)?$`)

// splitCompleteDateTimeZone splits date, time and timezone apart from an
// ISO8601 string. It does not ensure that the contents of each part are
// correct, it merely splits on certain delimiters.
// e.g "2010-09-08T12:15:10+0700" => "2010-09-08", "12:15:10", "+0700".
// Timezone can only be present if time is also present.
func splitCompleteDateTimeZone(s string) (dateStr, timeStr, zoneStr string, err error) {
	parts := completeDateTimeZoneRegexp.FindStringSubmatch(s)
	if parts == nil {
		err = fmt.Errorf("soap date/time/zone: value %q is not in ISO8601 datetime format", s)
		return
	}
	dateStr = parts[1]
	timeStr = parts[2]
	zoneStr = parts[3]
	return
}
