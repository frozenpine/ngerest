package ngerest

import (
	"testing"
	"time"
)

func TestNGETime(t *testing.T) {
	tm := NGETime{
		Time: time.Now().UTC(),
	}

	result, _ := tm.MarshalJSON()

	t.Log(string(result))
}
