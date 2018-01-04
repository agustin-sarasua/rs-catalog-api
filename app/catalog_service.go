package app

import (
	"errors"

	c "github.com/agustin-sarasua/rs-common"
	m "github.com/agustin-sarasua/rs-model"
)

func GetCatalogConfiguration(countryCode string, citiyCode string) (*CatalogConfigurationResponse, error) {

	countryCitites, exists := m.CountryCitites[countryCode]
	if !exists {
		return nil, errors.New("Country not valid")
	}
	city, exists := countryCitites.Cities[citiyCode]
	if !exists {
		return nil, errors.New("City not valid")
	}
	var config CatalogConfigurationResponse

	config = CatalogConfigurationResponse{
		Amenities:      c.MapKeys(m.Amenities),
		PropertyTypes:  c.MapKeys(m.PropertyTypes),
		Neighbourhoods: city.Neigbourhoods,
		Guarantees:     c.MapKeys(m.Guarantee)}
	return &config, nil
}

func GetCatalogCountryConfiguration(countryCode string) (*m.CountryConfig, error) {
	countryCitites, exists := m.CountryCitites[countryCode]
	if !exists {
		return nil, errors.New("Country not valid")
	}
	return &countryCitites, nil
}
