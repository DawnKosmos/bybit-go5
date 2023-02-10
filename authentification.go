package bybit

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"net/http"
	"strconv"
	"time"
)

var RECV_WINDOW = "5000"

/*
Parameters for Authenticated Endpoints
The following http header keys must be used for authentication:

X-BAPI-API-KEY - API key
X-BAPI-TIMESTAMP - UTC timestamp in milliseconds
X-BAPI-SIGN - a signature derived from the request's parameters
We also provide X-BAPI-RECV-WINDOW (unit in millisecond and default value is 5,000) to specify how long an HTTP request
is valid. It is also used to prevent replay attacks.

Basic steps:

1 timestamp + API key + (recv_window) + (queryString | jsonBodyString)
2 Use the HMAC_SHA256 algorithm to sign the string in step 1, and convert it to a hex string to obtain the sign parameter.
3 Append the sign parameter to request header, and send the HTTP request.
	Note: the plain text for GET and POST requests is different. Please refer to blew examples.
*/

func SignGET(a *Account, req *http.Request, params string) {
	tNow := time.Now().UnixMilli()
	msg := fmt.Sprintf("%d%s%s%s", tNow, a.PublicKey, RECV_WINDOW, params)
	hmac := hmac.New(sha256.New, []byte(a.SecretKey))
	hmac.Write([]byte(msg))
	sign := hex.EncodeToString(hmac.Sum(nil))

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-BAPI-API-KEY", a.PublicKey)
	req.Header.Set("X-BAPI-TIMESTAMP", strconv.FormatInt(tNow, 10))
	req.Header.Set("X-BAPI-SIGN", sign)
	req.Header.Set("X-BAPI-RECV-WINDOW", RECV_WINDOW)
}

func SignPOST(a *Account, req *http.Request, jsonParams []byte) {
	tNow := time.Now().UnixMilli()
	msg := fmt.Sprintf("%d%s%s%s", tNow, a.PublicKey, RECV_WINDOW, string(jsonParams))
	hmac := hmac.New(sha256.New, []byte(a.SecretKey))
	hmac.Write([]byte(msg))
	sign := hex.EncodeToString(hmac.Sum(nil))

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-BAPI-API-KEY", a.PublicKey)
	req.Header.Set("X-BAPI-TIMESTAMP", strconv.FormatInt(tNow, 10))
	req.Header.Set("X-BAPI-SIGN", sign)
	req.Header.Set("X-BAPI-RECV-WINDOW", RECV_WINDOW)
}
