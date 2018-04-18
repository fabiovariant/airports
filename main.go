package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"gopkg.in/mgo.v2/bson"
)

func main() {

	router := mux.NewRouter()
	router.HandleFunc("/airport/name/{sg_airport}", GetAirportByInitials).Methods("GET")
	c := cors.New(cors.Options{
		AllowedOrigins: []string{"http://*"},
		AllowedHeaders: []string{"Access-Control-Allow-Origin"},

		// Enable Debugging for testing, consider disabling in production
		Debug: true,
	})
	handler := c.Handler(router)

	log.Fatal(http.ListenAndServe(":8076", handler))
}

//GetAirportByInitials returns a JSON of airports by initials
func GetAirportByInitials(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	sgAirport := params["sg_airport"]
	s := GetMongoSession()
	defer s.Close()

	var airport Airport
	_, c := getSessionCopyAndCollectionConn(s)
	err := c.Find(bson.M{"sg_airport": sgAirport}).One(&airport)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err = json.NewEncoder(w).Encode(airport); err != nil {
		panic(err)
	}
}
