package app

import (
	"encoding/json"
	"log"
	"net/http"

	m "github.com/agustin-sarasua/rs-model"
	"github.com/gorilla/mux"
)

func mapKeys(m map[string]struct{}) []string {
	keys := make([]string, len(m))
	i := 0
	for k := range m {
		keys[i] = k
		i++
	}
	return keys
}

func GetCatalogConfigurationEndpoint(w http.ResponseWriter, req *http.Request) {
	country, _ := mux.Vars(req)["country"]
	city, _ := mux.Vars(req)["city"]

	if v, err := GetCatalogConfiguration(country, city); err != nil {
		log.Printf("Error getting catalog")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(m.BuildErrorResponse(err))
	} else {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(v)
	}
}

func ListCitiesEndpoint(w http.ResponseWriter, req *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	w.WriteHeader(http.StatusOK)
}
