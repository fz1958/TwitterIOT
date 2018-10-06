package TwitterIOT //All files should have the same Package Name [VPO]

import (
	"fmt"
	"net/http"

	"google.golang.org/appengine"
	"google.golang.org/appengine/datastore"
)

//This function set the claim_code flag to TRUE when a robot has been dispensed.
//This function receives the signal from the RaspberryPi [RPi] when the request has been executed successfully
//Main tasks to be executed in this function:
//1. Check Code: [Code] Property/Field
//2. Verify if the Robot has been taken: [Claimed] Property/Field
//3. Provide the Status- if the robot has not been claimed then this function will change the FLAG [Claimed] from False to True
func claimcode(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "[Claim code]")

	w.Header().Set("Content-Type", "text/plain; charset=utf-8")

	ctx := appengine.NewContext(r)
	newCode := r.URL.Query().Get("code") //Query in the code, don't actually store the code just grab it

	k := datastore.NewKey(ctx, kind, newCode, 0, nil)
	e := new(twitter)

	if err := datastore.Get(ctx, k, e); err != nil {
		//Cycle throught the data to see if the code matches
		//http.Error(w, err.Error(), 500) //datastore: no such entity
		fmt.Fprint(w, "The code was not found, it does not exist in the DataBase") //The code does not match
		return
	}

	if e.Claimed == true {
		fmt.Fprintf(w, "The code exists in the DataBase\n")                    //The code matches
		fmt.Fprintf(w, "You already claimed your little robot - Thank you!\n") //The user claimed the little robot successfully
	}

	if e.Claimed == false {
		//if the little robot has not been claimed then the FLAG [Claimed] will set its value to True
		fmt.Fprintf(w, "The code exists in the DataBase\n") //The code matches
		//fmt.Fprintln(w, "You have not clamed your robot!", e.Claimed) //This print statement can be removed, it is used for test purposes to see the output, or it can be keep and just remove: , e.Claimed
		e.Claimed = true
		fmt.Fprintln(w, "Thank you for claiming your little robot!", e.Claimed) //This print statement can be removed, it is used for test purposes to see the output, or it can be keep and just remove: , e.Claimed
		if _, err := datastore.Put(ctx, k, e); err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
	}

}
