package ngerest

import (
	// "os"
	"strconv"
	"strings"
	"time"
)

// NGETime NGE timestamp competibal with UTC time string & timestamp
type NGETime struct {
	time.Time
}

// func (t *NGETime) String() string {
// 	return t.Format("2006-01-02 15:04:05.000Z")
// }

// UnmarshalJSON convert time string or timestamp(ms)
func (t *NGETime) UnmarshalJSON(data []byte) error {
	if data == nil {
		t.Time = time.Unix(0, 0)
		return nil
	}

	dataStr := string(data)
	var err error

	if strings.HasPrefix(dataStr, "\"") && strings.HasSuffix(dataStr, "\"") {
		dataStr = strings.Trim(dataStr, "\" ")
		
		t.Time, err = time.ParseInLocation(
			"2006-01-02 15:04:05.000Z", dataStr, time.UTC)

		return err
	}

	timestamp, err := strconv.Atoi(dataStr)
	sec := int64(timestamp / 1000)
	nsec := (int64(timestamp) - sec*1000) * 1000
	t.Time = time.Unix(sec, nsec)

	return nil
}

// MarshalJSON convert Time structure in json
// func (t *NGETime) MarshalJSON() ([]byte, error) {
// 	useTimestamp := os.Getenv("USE_TIMESTAMP")

// 	useTimestamp = strings.Trim(useTimestamp, " \"")

// 	switch useTimestamp {
// 	case "1", "y", "Y", "t", "T", "yes", "Yes", "YES", "true", "True", "TRUE":
// 		ms := int(t.UnixNano() / 1000)
// 		data := []byte(strconv.Itoa(ms))
// 		return data, nil
// 	default:
// 		return []byte("\"" + t.String() + "\""), nil
// 	}
// }
