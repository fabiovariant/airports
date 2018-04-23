package main

import (
	"gopkg.in/mgo.v2/bson"
)

//Airport struct to JSON responses
type Airport struct {
	ID         bson.ObjectId `json:"id" bson:"_id,omitempty"`
	Name       string        `json:"name" bson:"name"`
	SgAirport  string        `json:"sg_airport" bson:"sg_airport"`
	Country    string        `json:"country" bson:"country"`
	State      string        `json:"state" bson:"state"`
	City       string        `json:"city" bson:"city"`
	TxBoarding float64       `json:"tx_boarding" bson:"tx_boarding"`
	Cad        string        `json:"cad" bson:"cad"`
	IsBlocked  bool          `json:"is_blocked" bson:"is_blocked"`
}

// Country is a struct type of countries collection
type Country struct {
	ID        bson.ObjectId `json:"id,omitempty" bson:"_id,omitempty"`
	CdCountry string        `json:"value,omitempty" bson:"cd_country"`
	NmCountry string        `json:"text,omitempty" bson:"nm_country"`
	States    []State       `json:"states,omitempty" bson:"states"`
}

// State is the struct of a State collection
type State struct {
	ID      bson.ObjectId `json:"id" bson:"_id,omitempty"`
	CdState string        `json:"value" bson:"cd_state"`
	NmState string        `json:"text" bson:"nm_state"`
}

// City is the strct of a City Collection.
type City struct {
	ID     bson.ObjectId `json:"id"        bson:"_id,omitempty"`
	CdCity string        `json:"value" bson:"cd_city"`
	NmCity string        `json:"text" bson:"nm_city"`
}
