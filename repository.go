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

//Return a copy of the mongodb session and a collection conn.
func getSessionCopyAndCollectionConn(s *mgo.Session) (session *mgo.Session, c *mgo.Collection) {
	session = s.Copy()
	c = session.DB("airports").C("airposts_list")
	return session, c
}
