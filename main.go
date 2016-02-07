package main

import (
	"fmt"
	"net/http"
	"github.com/manjeshpv/qginion/routes"
	"github.com/julienschmidt/httprouter"
)

const port string = ":3334"


func main() {

	r := httprouter.New()

	routes.Init(r)
	fmt.Printf("Running at %v\n", port)
	panic(http.ListenAndServe(port, r))



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
