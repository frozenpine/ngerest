package ngerest

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"

	"github.com/frozenpine/pkcs8"
)

var (
	_ context.Context
)

// KeyExchangeService public key exchange service for NGE
type KeyExchangeService service

// GetPublicKey get public key from nge
func (a *KeyExchangeService) GetPublicKey(ctx context.Context) (pkcs8.PublicKey, *http.Response, error) {
	var (
		localVarHttpMethod  = strings.ToUpper("POST")
		localVarPostBody    interface{}
		localVarFileName    string
		localVarFileBytes   []byte
		localVarReturnValue pkcs8.PublicKey
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/user/getPublicKey"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{"application/json", "application/x-www-form-urlencoded"}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{"application/json", "application/xml", "text/xml", "application/javascript", "text/javascript"}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}

	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
	localVarHttpResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	newErr := GenericSwaggerError{
		body:  localVarBody,
		error: localVarHttpResponse.Status,
	}

	if localVarHttpResponse.StatusCode < 300 {
		data := make(map[string]interface{})
		err := json.Unmarshal(localVarBody, &data)

		if err != nil {
			newErr.error = err.Error()
			newErr.body = localVarBody
			return localVarReturnValue, localVarHttpResponse, newErr
		}

		pubKeyString, ok := data["result"].(string)
		if !ok {
			newErr.error = "invalid result type in response"
			newErr.body = localVarBody
			return localVarReturnValue, localVarHttpResponse, newErr
		}

		localVarReturnValue, err := pkcs8.ParseFromPublicKeyString(pubKeyString, pkcs8.PKCS8)
		if err != nil {
			newErr.error = err.Error()
			newErr.body = localVarBody
			return localVarReturnValue, localVarHttpResponse, newErr
		}

		return localVarReturnValue, localVarHttpResponse, err
	} else {
		newErr.error = "request failed"
		newErr.body = localVarBody
		return localVarReturnValue, localVarHttpResponse, newErr
	}
}
