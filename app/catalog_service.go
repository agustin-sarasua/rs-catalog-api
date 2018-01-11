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

	var city *m.CityConfig
	for _, v := range countryCitites.Cities {
		if v.Code == citiyCode {
			// Found!
			city = &v
			break
		}
	}
	if city == nil {
		return nil, errors.New("City not valid")
	}

	var config CatalogConfigurationResponse

	config = CatalogConfigurationResponse{
		Neighbourhoods: city.Neighbourhoods,
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

func GetCatalogCountryPropertyTypes(countryCode string) (*PropertyTypesResponse, error){
	pt := m.CountryPropertyTypes[countryCode]
	return &PropertyTypesResponse{PropertyTypes: pt}, nil
}