package main

import (
	mgo "gopkg.in/mgo.v2"
)

//GetMongoSession return a session of MongoDB.
func GetMongoSession() (session *mgo.Session) {
	session, err := mgo.Dial("localhost")
	if err != nil {
		panic(err)
	}
	return
}
