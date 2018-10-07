package TwitterIOT

import (
	"net/http"

	"google.golang.org/appengine"
	"google.golang.org/appengine/log"
)

func twitterhook(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)
	log.Debugf(ctx, "Webhook success!")
}
