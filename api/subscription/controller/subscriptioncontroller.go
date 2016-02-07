package subscriptioncontroller

import (
	"github.com/manjeshpv/qginion/api/subscription/dao"
	"encoding/json"
	"github.com/manjeshpv/qginion/api/subscription/model"
	"github.com/julienschmidt/httprouter"
	"io/ioutil"
	"net/http"
	"fmt"

)

type Person struct {
	Name string
	Phone string
}


func GetAll(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")
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
func NewSubscription(res http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	res.Header().Set("Content-Type", "application/json")

	schema := subscriptionmodel.Subscription{}

	subscriptionBody, err := ioutil.ReadAll(req.Body)



	if err != nil {
		res.WriteHeader(http.StatusBadRequest)
		return
	}

	fmt.Println("request body",string(subscriptionBody))

	fmt.Println("schema:", &schema)

	defer req.Body.Close()

	err = json.Unmarshal(subscriptionBody, &schema)
	fmt.Println("errr",err)
	if err != nil {
		fmt.Println(err)
		res.WriteHeader(http.StatusInternalServerError)
		return
	}

	nt, err := subsriptiondao.NewSubscription(schema)

	fmt.Println("err",err)

	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		return
	}

	fmt.Println("nt:", nt)

	ntm, err := json.Marshal(nt)

	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		return
	}
	fmt.Println("ntm:", nt)

	res.WriteHeader(http.StatusCreated)

	res.Write(ntm)
}
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
