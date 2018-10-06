package TwitterIOT

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"net/http"

	"google.golang.org/appengine"
	"google.golang.org/appengine/log"
)

// ComputeHmac256 returns the CRC check token string
func ComputeHmac256(message string, secret string) string {
	key := []byte(secret)
	h := hmac.New(sha256.New, key)
	h.Write([]byte(message))
	return base64.StdEncoding.EncodeToString(h.Sum(nil))
}

// TwitterResponse represents the json response required for twitter CRC
type TwitterResponse struct {
	ResponseToken string `json:"response_token"`
}

// GetCRCResponse returns the Json object required for Twitter CRC check
// Parameters
// 		message (string) - message from twitter
//		secret (string) - Twitter consumer secret
//
// Example JSON response (from twitter docs):
//		{
//  		"response_token": "sha256=x0mYd8hz2goCTfcNAaMqENy2BFgJJfJOb4PdvTffpwg="
//		}
//
func GetCRCResponse(message string, secret string, w http.ResponseWriter) error {
	s := ComputeHmac256(message, secret)
	tr := &TwitterResponse{
		ResponseToken: "sha256=" + s,
	}

	// encode the response as json and send to web output stream, in this case: ResponseWriter
	encoder := json.NewEncoder(w)
	if err := encoder.Encode(tr); err != nil {
		return err
	}

	return nil
}

// testCRCResponse is a quick test function for the CRC check required for twitter.
//
//		URL: /app/test/crt?message=a9sd87f98s6a7f
func testCRCResponse(w http.ResponseWriter, r *http.Request) {

	ctx := appengine.NewContext(r)

	if r.Method == http.MethodGet {
		// this should be the consumer secret that is provided to the twitter account
		secret := "89s76d9fasdasdf"

		// messsage is pulled from the URL query but twitter will provide it as a JSON object that
		// needs to be decoded
		message := r.URL.Query().Get("message")

		// if there is an error decoding we should log it
		err := GetCRCResponse(message, secret, w)
		if err != nil {
			log.Errorf(ctx, "there was an error encoding the response %s", err)
		}
	} else {
		// Twitter will make a call to the webhook endpoint in the form of a POST for actual data
		// but will call the same endpoint using GET to do the CRC check
		log.Debugf(ctx, "Method is not GET it is %s", r.Method)
	}
}
