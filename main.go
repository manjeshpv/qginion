package main

import (
//	"net/http"
//	"encoding/json"
	"time"
	"fmt"
	"log"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)
type Person struct {
	Name string
	Phone string
}
//const MongoDb details
const (
	hosts      = "ds063439.mongolab.com:63439"
	database   = "tsc"
	username   = "qnotify"
	password   = "quezx123"
	collection = "messages"
)


func main() {
//	ds063439.mongolab.com:63439
	info := &mgo.DialInfo{
		Addrs:    []string{hosts},
		Timeout:  60 * time.Second,
		Database: database,
		Username: username,
		Password: password,
	}

	session, err := mgo.DialWithInfo(info)
//	session, err := mgo.Dial("qnotify:quezx123@ds063439.mongolab.com:63439")
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
