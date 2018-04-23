package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func main() {

	router := mux.NewRouter()
	router.HandleFunc("/register/countries", GetCountries).Methods("GET")
	router.HandleFunc("/register/countries/states/{cd_country}", GetStatesByCountry).Methods("GET")
	router.HandleFunc("/airport/name/{sg_airport}", GetAirportByInitials).Methods("GET")

	c := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedHeaders: []string{"Access-Control-Allow-Origin"},

		// Enable Debugging for testing, consider disabling in production
		Debug: true,
	})
	handler := c.Handler(router)

	log.Fatal(http.ListenAndServe(":8076", handler))
}
