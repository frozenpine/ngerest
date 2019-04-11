package ngerest_test

import (
	"encoding/json"
	"fmt"
	"testing"
	"time"

	"github.com/frozenpine/ngerest"
)

func TestJavaTime(t *testing.T) {
	timestamp := int64(time.Now().UnixNano() / 1000)
	var javaTime ngerest.JavaTime

	err := json.Unmarshal([]byte(fmt.Sprintf("%d", timestamp)), &javaTime)

	if err != nil {
		t.Fatal(err)
	}

	result, _ := json.Marshal(javaTime)
	t.Log(string(result))
}
