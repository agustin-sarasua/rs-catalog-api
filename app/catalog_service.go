package app

import (
	"errors"

	m "github.com/agustin-sarasua/rs-model"
)

func GetCatalogConfiguration(countryCode string, citiyCode string) (*CatalogConfigurationResponse, error) {

	countryCitites, exists := m.CountryCitites[countryCode]
	if !exists {
		return nil, errors.New("Country not valid")
	}
	city, exists := countryCitites.Cities[citiyCode]
	if !exists {
		return nil, errors.New("Country not valid")
	}
	var config CatalogConfigurationResponse

	config = CatalogConfigurationResponse{
		Amenities:      mapKeys(m.Amenities),
		PropertyTypes:  mapKeys(m.PropertyTypes),
		Neighbourhoods: city.Neigbourhoods,
	}
	return &config, nil
}
