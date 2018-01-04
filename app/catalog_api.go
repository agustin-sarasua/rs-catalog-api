package app

import (
	m "github.com/agustin-sarasua/rs-model"
)

type CatalogConfigurationResponse struct {
	Neighbourhoods []m.NameCode
	PropertyTypes  []string
	Amenities      []string
	Guarantees     []string
}
