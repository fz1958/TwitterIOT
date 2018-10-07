package TwitterIOT

import (
	"encoding/json"
	"net/http"

	"google.golang.org/appengine"
	"google.golang.org/appengine/log"
)

//Twitter Webhook
type testStruct struct {
	Test string
}

func twitterhook(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)
	//Log that this function is called
	log.Debugf(ctx, "Webhook success!")

	switch r.Method {
	//If requesting a "Get" run CRC
	case http.MethodGet:
		log.Debugf(ctx, "This is a Get method")
		secret := "8LRLAFoS7FlUE6q2hJKMG2kbtBXLPUTyl6btc1PABayI3416IV"
		message := r.URL.Query().Get("message")
		GetCRCResponse(secret, message, w)

		//If post parse data
	case http.MethodPost:
		log.Debugf(ctx, "This is a Post method")
		decoder := json.NewDecoder(r.Body)
		var t testStruct
		err := decoder.Decode(&t)
		if err != nil {
			panic(err)
		}
		log.Debugf(ctx, t.Test)
		//When Put is called, it is asking to renew so have CRC run
	case http.MethodPut:
		log.Debugf(ctx, "This is a Put method")
		secret := "8LRLAFoS7FlUE6q2hJKMG2kbtBXLPUTyl6btc1PABayI3416IV"
		message := r.URL.Query().Get("message")
		GetCRCResponse(secret, message, w)
	default: //Else this gets awkward
		log.Debugf(ctx, "Yikes this wasn't supposed to happen (Not a GET or POST)")
	}
}
