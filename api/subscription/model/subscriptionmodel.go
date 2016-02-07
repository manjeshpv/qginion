package subscriptionmodel

import (
	"gopkg.in/mgo.v2/bson"
	"time"
)

type Subscription struct {
	Id          bson.ObjectId `json:"_id,omitempty" bson:"_id,omitempty"`
	Name string        `json:"todoMessage,omitempty" bson:"Name"`
	Phone string        `json:"todoMessage,omitempty" bson:"Phone"`
	CreatedAt   time.Time     `json:"createdAt,omitempty" bson:"createdAt"`
}

func (t Subscription) IsValid() bool {

	if l := len(t.Name); l > 4 {
		return true
	}

	return false
}

type Subscriptions []Subscription
