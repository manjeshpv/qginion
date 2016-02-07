package main

import (
//	"net/http"
//	"encoding/json"
	"github.com/manjeshpv/qginion/config"
	"fmt"
	"log"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
//	"errors"
)
type Person struct {
	Name string
	Phone string
}



func main() {

	db := dbconfig.DB{}

	session, err := db.DoDial()

	if err != nil {
		panic(err)
	}

	defer session.Close()

	// Optional. Switch the session to a monotonic behavior.
	session.SetMode(mgo.Monotonic, true)

	c := session.DB("tsc").C("people")

	err = c.Insert(&Person{"Ale", "+55 53 8116 9639"},
		&Person{"Cla", "+55 53 8402 8510"})
	if err != nil {
		log.Fatal(err)
	}

	result := Person{}
	err = c.Find(bson.M{"name": "Ale"}).One(&result)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Phone:", result.Phone)
}
//type Profile struct {
//	Name    string
//	Hobbies []string
//}
//
//func main() {
//	http.HandleFunc("/", foo)
//	http.ListenAndServe(":3000", nil)
//}
//
//func foo(w http.ResponseWriter, r *http.Request) {
////	profile := Profile{"Alex", []string{"snowboarding", "programming"}}
////
////	js, err := json.Marshal(profile)
////	if err != nil {
////		http.Error(w, err.Error(), http.StatusInternalServerError)
////		return
////	}
////
////	w.Header().Set("Content-Type", "application/json")
////	w.Write(js)
//
//
//}
