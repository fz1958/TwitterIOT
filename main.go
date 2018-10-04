package TwitterIOT

import (
	"fmt"
	"math/rand"
	"net/http"
	"time"

	"google.golang.org/appengine"
	"google.golang.org/appengine/datastore"
)

func init() {
	http.HandleFunc("/app/twitter/register", registerTwitterUser)
	http.HandleFunc("/app/twitter/code", storeCode)
	http.HandleFunc("/app/verify", verifyCode)
	http.HandleFunc("/app/testGen", TestGen)
	rand.Seed(time.Now().UnixNano())
}

//This struct will store
type twitter struct {
	UserName string
	Code     string
	Claimed  bool
}

const kind string = "twitter"

//number made for rune
var numberRune = []rune("0123456789")

//CreateCode -- Makes a 7 digit code and stops it from
func CreateCode() string {
	//Creates a random 7 digit code
	genCode := make([]rune, 7)
	for i := range genCode {
		genCode[i] = numberRune[rand.Intn(len(numberRune))]
	}
	return string(genCode)
}

//storeCode -- Used for testing storing codes
//			URL /app/twitter/code
func storeCode(w http.ResponseWriter, r *http.Request) {
	//store the users code with the persons ID
	//Check if its been used

	ctx := appengine.NewContext(r)
	code := r.URL.Query().Get("code")
	k := datastore.NewKey(ctx, kind, code, 0, nil)
	e := &twitter{
		UserName: "Austin",
	}

	if _, err := datastore.Put(ctx, k, e); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

}

func verifyCode(w http.ResponseWriter, r *http.Request) {
	//Query in the new code
	//Don't actually store the code just grab it
	//Cycle throught the data to see if new code matches

	w.Header().Set("Content-Type", "text/plain; charset=utf-8")

	ctx := appengine.NewContext(r)
	newCode := r.URL.Query().Get("code")

	k := datastore.NewKey(ctx, kind, newCode, 0, nil)
	e := new(twitter)
	if err := datastore.Get(ctx, k, e); err != nil {
		http.Error(w, err.Error(), 500)
		fmt.Fprint(w, "Bad Code")
		return
	}

	fmt.Fprintf(w, "good Code", e)
}

//datastore.put
//First you put the variable you are putting in
//next is the key that leads you to that

func registerTwitterUser(w http.ResponseWriter, r *http.Request) {
	//First we need to read the twitter username/ID
	//send the username to the database
	//Send the code to claim prize

	ctx := appengine.NewContext(r)
	twitterID := r.URL.Query().Get("ID")
	generateCode := CreateCode()
	k := datastore.NewKey(ctx, kind, generateCode, 0, nil)
	e := &twitter{
		Code:     generateCode,
		UserName: twitterID,
	}

	if _, err := datastore.Put(ctx, k, e); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

}

func TestGen(w http.ResponseWriter, r *http.Request) {
	i := 0
	duplicate := false
	var m map[string]string
	m = make(map[string]string)

	for !duplicate {

		CodeTest := CreateCode()

		if len(m[CodeTest]) > 1 {
			fmt.Fprintf(w, "DUPLICATE")
			duplicate = true
		}

		m[CodeTest] = "whatever"

		i++
	}
	fmt.Fprintf(w, "Total runs %v", i)

}

//Current Goal: Is to querey the code from the url when someone puts in the code.
//After getting the code someone puts in I will check it against every code in the database.
//Once the code is checked and finds a match I will set the variable boolean equal to true.
