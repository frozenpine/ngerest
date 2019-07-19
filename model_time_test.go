package ngerest

import (
	"encoding/json"
	"os"
	"testing"
	"time"
)

func TestTimeMarshalJSON(t *testing.T) {
	tm := NGETime{Time: time.Now()}

	data, err := json.Marshal(tm)
	if err != nil {
		t.Error(err)
	} else {
		t.Log(string(data))
	}

	os.Setenv("USE_TIMESTAMP", "TRUE")

	data, err = json.Marshal(tm)
}
