package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"gopkg.in/mgo.v2"

	"github.com/gorilla/mux"
	"gopkg.in/mgo.v2/bson"
)

// GetAirportByInitials returns a JSON of airports by initials
func GetAirportByInitials(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	sgAirport := params["sg_airport"]
	s := GetMongoSession()
	defer s.Close()

	var airport Airport
	c := s.DB("airports").C("airposts_list")
	err := c.Find(bson.M{"sg_airport": sgAirport}).One(&airport)
	if err != nil {
		w.WriteHeader(http.StatusNoContent)
		return
	}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err = json.NewEncoder(w).Encode(airport); err != nil {
		panic(err)
	}
}

// GetCountries returnus all registred Countries on Dabase
func GetCountries(w http.ResponseWriter, r *http.Request) {
	s := GetMongoSession()
	defer s.Close()
	var v []Country

	c := s.DB("airports").C("countries")
	err := c.Find(bson.M{}).All(&v)
	if err != nil {
		panic(err)
	}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err = json.NewEncoder(w).Encode(v); err != nil {
		panic(err)
	}
}

// GetStatesByCountry returns all states of a Country
func GetStatesByCountry(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	cdCountry := params["cd_country"]
	var country Country

	s := GetMongoSession()
	defer s.Close()

	c := s.DB("airports").C("countries")
	err := c.Find(bson.M{"cd_country": cdCountry}).Select(bson.M{"states": 1}).One(&country)
	if err != nil {
		panic(err)
	}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err = json.NewEncoder(w).Encode(country); err != nil {
		panic(err)
	}
}

// GetCitiesByState returns a array o cities of a state
func GetCitiesByState(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	cdCountry := params["cd_country"]
	cdState := params["cd_state"]
	var state State

	s := GetMongoSession()
	defer s.Close()

	c := s.DB("airports").C("cities")
	query := bson.M{"$and": []bson.M{
		bson.M{"cd_country": cdCountry},
		bson.M{"cd_state": cdState},
	}}
	err := c.Find(query).One(&state)
	if err != nil {
		if err != mgo.ErrNotFound {
			panic(err)
		}
		fmt.Println(err)
	}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err = json.NewEncoder(w).Encode(state); err != nil {
		panic(err)
	}
}

// AddAirport insert a airport on the mongo database
func AddAirport(w http.ResponseWriter, r *http.Request) {
	var air Airport
	if err := json.NewDecoder(r.Body).Decode(air); err != nil {
		panic(err)
	}
	i := bson.NewObjectId()
	air.ID = i

	s := GetMongoSession()
	defer s.Close()

	c := s.DB("airports").C("airposts_list")
	err := c.Insert(air)
	if err != nil {
		panic(err)
	}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err = json.NewEncoder(w).Encode(air); err != nil {
		panic(err)
	}
}
