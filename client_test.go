package ngerest

import (
	"bytes"
	"net/url"
	"testing"
)

func TestGenerateSignature(t *testing.T) {
	var (
		apiKey     = "L68h3Fn3QjRaKqhaU82Y"
		apiSecret  = "2AWH0b47nVxWxRZ14q7x8KH4pAgRp20n96oLHC6PQlI6WT4oUfJFZiQh0429t7p7I633j3vv55Q9DiuNESdOvnmKux6n01Ogj3y"
		bodyString = "{\"symbol\": \"XBTUSD\", \"orderQty\": 1, \"price\": 3536.7, \"side\": \"Buy\"}"
		signature  = "c4e9ea35fe2ae4ca8f08406b957fae906f5a54f88149aea689451538db78ef16"
		expires    = 1551691465
	)
	client := NewAPIClient(NewConfiguration())

	url, err := url.Parse("/api/v1/order")

	if err != nil {
		t.Error(err)
		return
	}

	body := &bytes.Buffer{}

	body.WriteString(bodyString)

	sig := client.generateSignature(apiSecret, "post", url, expires, body)

	if sig != signature {
		t.Error("singnature checking failed.")
	} else {
		t.Logf("%s api key checking success.", apiKey)
	}
}
