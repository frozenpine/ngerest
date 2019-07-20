package ngerest

import (
	"regexp"
	"strconv"
	"strings"
	"time"
)

// StringFloat json string format float64
type StringFloat float64

// UnmarshalJSON unmarshal float from json
func (f *StringFloat) UnmarshalJSON(data []byte) error {
	dataStr := string(data)
	dataStr = strings.Trim(dataStr, "\" ")

	if dataStr == "" {
		*f = 0
		return nil
	}

	value, err := strconv.ParseFloat(dataStr, 64)
	if err != nil {
		return err
	}

	*f = StringFloat(value)
	return nil
}

var timestampPattern = regexp.MustCompile("^[0-9]+$")
var timeStringPattern = regexp.MustCompile("([0-9]{2}:){3}[0-9]{3}Z$")

// NGETime NGE timestamp competibal with UTC time string & timestamp
type NGETime struct {
	time.Time
}

// UnmarshalJSON convert time string or timestamp(ms)
func (t *NGETime) UnmarshalJSON(data []byte) error {
	if data == nil {
		t.Time = time.Unix(0, 0)
		return nil
	}

	dataStr := string(data)
	dataStr = strings.Trim(dataStr, "\" ")

	var err error

	if timestampPattern.MatchString(dataStr) {
		timestamp, err := strconv.Atoi(dataStr)

		if err != nil {
			return err
		}

		sec := int64(timestamp / 1000)
		nsec := (int64(timestamp) - sec*1000) * 1000
		t.Time = time.Unix(sec, nsec)

		return nil
	}

	if dataStr == "" {
		t.Time = time.Unix(0, 0)
		return nil
	}

	if timeStringPattern.MatchString(dataStr) {
		seps := strings.Split(dataStr, ":")

		head := strings.Join(seps[:3], ":")

		dataStr = head + "." + seps[3]
	}

	t.Time, err = time.ParseInLocation(
		"2006-01-02 15:04:05.000Z", dataStr, time.UTC)

	t.Time = t.In(time.Local)

	return err
}
