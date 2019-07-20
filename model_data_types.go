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

	if strings.Contains(dataStr, "T") {
		t.Time, err = time.ParseInLocation(
			"2006-01-02T15:04:05:000Z", dataStr, time.UTC)
	} else {
		t.Time, err = time.ParseInLocation(
			"2006-01-02 15:04:05.000Z", dataStr, time.UTC)
	}

	return err
}
