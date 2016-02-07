package subsriptiondao

import (
	"errors"
	subscription "github.com/manjeshpv/qginion/api/subscription/model"
	"github.com/manjeshpv/qginion/config"
	"gopkg.in/mgo.v2/bson"
	"time"
	"log"
	"fmt"
)

const col string = "subscriptions"

type Person struct {
	Name string
	Phone string
}

func All() (subscription.Subscriptions, error) {

	db := dbconfig.DB{}

	ts := subscription.Subscriptions{}

	s, err := db.DoDial()

	if err != nil {
		return ts, errors.New("There was an error trying to connect with the DB.")
	}

	defer s.Close()

	c := s.DB(db.Name()).C(col)

	err = c.Find(bson.M{}).All(&ts)

	if err != nil {
		return ts, errors.New("There was an error trying to find the todos.")
	}

	return ts, err
}

func NewSubscription(subscriptionToSave subscription.Subscription) (subscription.Subscription, error) {

	db := dbconfig.DB{}

	subscriptionToSave.Id = bson.NewObjectId()

	subscriptionToSave.CreatedAt = time.Now()

	s, err := db.DoDial()

	if err != nil {
		return subscriptionToSave, errors.New("There was an error trying to connect to the DB.")
	}

	defer s.Close()

	c := s.DB(db.Name()).C(col)


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

	err = c.Insert(subscriptionToSave)

	if err != nil {
		return subscriptionToSave, errors.New("There was an error trying to insert the doc to the DB.")
	}

	return subscriptionToSave, err
}
//
//func DeleteTodo(id string) error {
//	db := dbconfig.DB{}
//
//	s, err := db.DoDial()
//
//	if err != nil {
//		return errors.New("There was an error trying to connect with the DB.")
//	}
//
//	idoi := bson.ObjectIdHex(id)
//
//	defer s.Close()
//
//	c := s.DB(db.Name()).C(col)
//
//	err = c.RemoveId(idoi)
//
//	if err != nil {
//		return errors.New("There was an error trying to remove the todo.")
//	}
//
//	return err
//}
