package ngerest

import (
	"encoding/json"
	"testing"
	"time"
)

func TestNGETime(t *testing.T) {
	tm := NGETime(time.Now().UTC())
	// t.Log(tm.String())

	result, _ := json.Marshal(tm)

	t.Log(string(result))
}
