package subscriptioncontroller

import (
//	"github.com/manjeshpv/qginion/config"
//	"fmt"
//	"log"
//	"gopkg.in/mgo.v2"
//	"gopkg.in/mgo.v2/bson"
	"github.com/manjeshpv/qginion/api/subscription/dao"
	"encoding/json"
//	"github.com/manjeshpv/qgotify/server/api/todo/dao"
//	todo "github.com/manjeshpv/qgotify/server/api/todo/model"
	"github.com/julienschmidt/httprouter"
//	"io/ioutil"
	"net/http"
)

type Person struct {
	Name string
	Phone string
}


func GetAll(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	ts, err := subsriptiondao.All()

//	db := dbconfig.DB{}
//
//	session, err := db.DoDial()
//
//	if err != nil {
//		panic(err)
//	}
//
//	defer session.Close()
//
//	// Optional. Switch the session to a monotonic behavior.
//	session.SetMode(mgo.Monotonic, true)
//
//	c := session.DB("tsc").C("people")
//
//	err = c.Insert(&Person{"Ale", "+55 53 8116 9639"},
//		&Person{"Cla", "+55 53 8402 8510"})
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	result := Person{}
//	err = c.Find(bson.M{"name": "Ale"}).One(&result)
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	fmt.Println("Phone:", result.Phone)

	tsm, err := json.Marshal(ts)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(tsm)

}
//func GetAll(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
//	ts, err := tododao.All()
//
//	w.Header().Set("Content-Type", "application/json")
//
//	if err != nil {
//		w.WriteHeader(http.StatusInternalServerError)
//		return
//	}
//
//	tsm, err := json.Marshal(ts)
//
//	if err != nil {
//		w.WriteHeader(http.StatusInternalServerError)
//		return
//	}
//
//	w.WriteHeader(http.StatusOK)
//
//	w.Write(tsm)
//}
//
//func NewTodo(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
//	w.Header().Set("Content-Type", "application/json")
//
//	t := todo.Todo{}
//
//	tf, err := ioutil.ReadAll(r.Body)
//
//	if err != nil {
//		w.WriteHeader(http.StatusBadRequest)
//		return
//	}
//
//	defer r.Body.Close()
//
//	err = json.Unmarshal(tf, &t)
//
//	if err != nil {
//		w.WriteHeader(http.StatusInternalServerError)
//		return
//	}
//
//	nt, err := tododao.NewTodo(t)
//
//	if err != nil {
//		w.WriteHeader(http.StatusInternalServerError)
//		return
//	}
//
//	ntm, err := json.Marshal(nt)
//
//	if err != nil {
//		w.WriteHeader(http.StatusInternalServerError)
//		return
//	}
//
//	w.WriteHeader(http.StatusCreated)
//
//	w.Write(ntm)
//}
//
//func RemoveTodo(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
//
//	id := ps.ByName("id")
//
//	err := tododao.DeleteTodo(id)
//
//	if err != nil {
//		w.WriteHeader(http.StatusInternalServerError)
//		return
//	}
//
//	w.WriteHeader(http.StatusOK)
//}
