package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/agustin-sarasua/rs-catalog-api/app"
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/catalog/{country:[UY|AR]+}/{city:.+}", app.GetCatalogConfigurationEndpoint).Methods("GET")
	router.HandleFunc("/catalog/country/{country:[UY|AR]+}", app.GetCatalogCountryConfigurationEndpoint).Methods("GET")
	router.HandleFunc("/catalog/country/{country:[UY|AR]+}/property-types", app.GetCatalogCountryPropertyTypesEndpoint).Methods("GET")

	fmt.Println("Hello there")
	log.Fatal(http.ListenAndServe(":8080", router))
}

func use(h http.HandlerFunc, middleware ...func(http.HandlerFunc) http.HandlerFunc) http.HandlerFunc {
	for _, m := range middleware {
		h = m(h)
	}
	return h
}
