package app

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func GetCatalogConfigurationEndpoint(w http.ResponseWriter, req *http.Request) {
	country, _ := mux.Vars(req)["country"]
	city, _ := mux.Vars(req)["city"]

	fmt.Println(country)
	fmt.Println(city)
	var config CatalogConfigurationResponse
	config = CatalogConfigurationResponse{
		Amenities:      Amenities,
		PropertyTypes:  PropertyTypes,
		Neighbourhoods: []string{""},
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(config)
}

func ListCitiesEndpoint(w http.ResponseWriter, req *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	w.WriteHeader(http.StatusOK)
}
