/*
 * BitMEX API
 *
 * ## REST API for the BitMEX Trading Platform  [View Changelog](/app/apiChangelog)    #### Getting Started  Base URI: [https://www.bitmex.com/api/v1](/api/v1)  ##### Fetching Data  All REST endpoints are documented below. You can try out any query right from this interface.  Most table queries accept `count`, `start`, and `reverse` params. Set `reverse=true` to get rows newest-first.  Additional documentation regarding filters, timestamps, and authentication is available in [the main API documentation](/app/restAPI).  *All* table data is available via the [Websocket](/app/wsAPI). We highly recommend using the socket if you want to have the quickest possible data without being subject to ratelimits.  ##### Return Types  By default, all data is returned as JSON. Send `?_format=csv` to get CSV data or `?_format=xml` to get XML data.  ##### Trade Data Queries  *This is only a small subset of what is available, to get you started.*  Fill in the parameters and click the `Try it out!` button to try any of these queries.  * [Pricing Data](#!/Quote/Quote_get)  * [Trade Data](#!/Trade/Trade_get)  * [OrderBook Data](#!/OrderBook/OrderBook_getL2)  * [Settlement Data](#!/Settlement/Settlement_get)  * [Exchange Statistics](#!/Stats/Stats_history)  Every function of the BitMEX.com platform is exposed here and documented. Many more functions are available.  ##### Swagger Specification  [⇩ Download Swagger JSON](swagger.json)    ## All API Endpoints  Click to expand a section.
 *
 * API version: 1.2.0
 * Contact: support@bitmex.com
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package ngerest

import (
	"bytes"
	"strconv"
	"time"
)

// XAny x any
type XAny struct {
}

// JavaTime Java timestamp in million sec
type JavaTime time.Time

// UnmarshalJSON unmarshal java timestamp
func (t *JavaTime) UnmarshalJSON(data []byte) error {
	millis, err := strconv.ParseInt(string(data), 10, 64)

	if err != nil {
		return err
	}

	*t = JavaTime(time.Unix(0, millis*int64(time.Millisecond)))
	return nil
}

// MarshalJSON marshal Time to java timestamp
func (t *JavaTime) MarshalJSON() (data []byte, err error) {
	var buf bytes.Buffer

	buf.WriteString(strconv.FormatInt(time.Time(*t).UnixNano()/int64(time.Millisecond), 10))

	return buf.Bytes(), nil
}
