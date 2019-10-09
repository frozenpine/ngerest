package ngerest

import (
	"encoding/json"
	"testing"
	"time"
)

func TestNGETime(t *testing.T) {
	tm := NGETime(time.Now().UTC())
	result, _ := json.Marshal(tm)
	t.Log(string(result))

	tm.FromTimestamp(1570639648629)
	result, _ = json.Marshal(tm)

	t.Log(string(result))
}
