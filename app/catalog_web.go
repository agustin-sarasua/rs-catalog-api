package app

import (
	"encoding/json"
	"log"
	"net/http"

	m "github.com/agustin-sarasua/rs-model"
	"github.com/gorilla/mux"
)

func GetCatalogConfigurationEndpoint(w http.ResponseWriter, req *http.Request) {
	country, _ := mux.Vars(req)["country"]
	city, _ := mux.Vars(req)["city"]

	if v, err := GetCatalogConfiguration(country, city); err != nil {
		log.Printf("Error getting catalog")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(m.BuildErrorResponse([]error{err}))
	} else {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(v)
	}
}

func GetCatalogCountryConfigurationEndpoint(w http.ResponseWriter, req *http.Request) {
	country, _ := mux.Vars(req)["country"]
	if v, err := GetCatalogCountryConfiguration(country); err != nil {
		log.Printf("Error getting catalog")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(m.BuildErrorResponse([]error{err}))
	} else {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(v)
	}
}

func GetCatalogCountryPropertyTypesEndpoint(w http.ResponseWriter, req *http.Request) {
	country, _ := mux.Vars(req)["country"]
	if v, err := GetCatalogCountryPropertyTypes(country); err != nil {
		log.Printf("Error getting propertyTypes")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(m.BuildErrorResponse([]error{err}))
	} else {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(v)
	}
}

func GetCatalogCountryAmenitiesEndpoint(w http.ResponseWriter, req *http.Request) {
	country, _ := mux.Vars(req)["country"]
	if v, err := GetCatalogCountryAmenities(country); err != nil {
		log.Printf("Error getting amenities")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(m.BuildErrorResponse([]error{err}))
	} else {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(v)
	}
}