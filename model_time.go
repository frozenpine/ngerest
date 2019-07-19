package ngerest

import (
	"strconv"
	"strings"
	"time"
)

// Time NGE timestamp competibal with UTC time string & timestamp
type Time struct {
	time.Time
}

func (t *Time) String() string {
	return t.Format("2006-01-02 15:04:05.000Z")
}

// UnmarshalJSON convert time string or timestamp(ms)
func (t *Time) UnmarshalJSON(data []byte) error {
	if data == nil {
		t.Time = time.Unix(0, 0)
		return nil
	}

	dataStr := string(data)
	var err error

	if strings.HasPrefix(dataStr, "\"") && strings.HasSuffix(dataStr, "\"") {
		dataStr = strings.TrimSuffix(
			strings.TrimPrefix(dataStr, "\""), "\"")
		t.Time, err = time.Parse("2006-01-02 15:04:05.000Z", dataStr)

		return err
	}

	timestamp, err := strconv.Atoi(dataStr)
	sec := int64(timestamp / 1000)
	nsec := (int64(timestamp) - sec*1000) * 1000
	t.Time = time.Unix(sec, nsec)

	return nil
}

// MarshalJSON convert Time structure in json
func (t *Time) MarshalJSON() ([]byte, error) {
	ms := int(t.UnixNano() / 1000)

	data := []byte(strconv.Itoa(ms))

	return data, nil
}
