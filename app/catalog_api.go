package app

import (
	m "github.com/agustin-sarasua/rs-model"
)

type CatalogConfigurationResponse struct {
	Neighbourhoods []m.Neigbourhood
	PropertyTypes  []string
	Amenities      []string
	Guarantees     []string
}

type CatalogContryResponse struct {
	Citites []string
}
