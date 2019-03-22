package ngerest

import (
	"context"
	"testing"

	"github.com/frozenpine/pkcs8"
)

func TestGetPublicKey(t *testing.T) {
	client := NewAPIClient(NewConfiguration())
	client.ChangeBasePath("http://trade/api/v1")

	ctx := context.Background()

	key, rsp, err := client.KeyExchange.GetPublicKey(ctx)
	if rsp.StatusCode > 300 || err != nil {
		t.Fatal(err.(GenericSwaggerError).Error(), string(err.(GenericSwaggerError).Body()))
	}
	t.Log(key)

	login := map[string]string{
		"email":      "yuanyang@quantdo.com.cn",
		"password":   key.Encrypt("yuanyang"),
		"type":       "account",
		"verifyCode": "",
	}
	auth, rspLogin, err := client.UserApi.UserLogin(ctx, login)
	if rspLogin.StatusCode > 300 || err != nil {
		t.Fatal(err.(GenericSwaggerError).Error(), string(err.(GenericSwaggerError).Body()))
	}
	t.Log(auth.Value(ContextQuantToken))

	privateKey := pkcs8.GeneratePriveKey(2048)

	defaultKey, rspKey, err := client.UserApi.UserGetDefaultAPIKey(auth, privateKey)
	if rspKey.StatusCode > 300 || err != nil {
		t.Fatal(err.(GenericSwaggerError).Error(), string(err.(GenericSwaggerError).Body()))
	}
	t.Log(defaultKey)
}
