package main

import (
	"gopkg.in/mgo.v2/bson"
)

//Airport struct to JSON responses
type Airport struct {
	ID        bson.ObjectId `json:"id"        bson:"_id,omitempty"`
	Name      string        `json:"name" bson:"name"`
	SgAirport string        `json:"sg_airport" bson:"sg_airport"`
}
