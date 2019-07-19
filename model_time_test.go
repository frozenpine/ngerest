package ngerest

import (
	// "encoding/json"
	"os"
	"testing"
	"time"
)

func TestTimeMarshalJSON(t *testing.T) {
	tm := NGETime{Time: time.Now()}

	loc, _ := time.LoadLocation("UTC")
	tm.In(loc)

	data, err := tm.MarshalJSON()
	if err != nil {
		t.Error(err)
	} else {
		t.Log(string(data))
		t.Log(tm.String())
	}

	os.Setenv("USE_TIMESTAMP", "TRUE")

	data, err = tm.MarshalJSON()
	if err != nil {
		t.Error(err)
	} else {
		t.Log(string(data))
	}
}
