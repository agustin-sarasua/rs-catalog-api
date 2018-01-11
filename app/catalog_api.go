package app

import (
	m "github.com/agustin-sarasua/rs-model"
)

type CatalogConfigurationResponse struct {
	Neighbourhoods []m.NameCode
	Guarantees     []string
}

type ConfigurationResponse struct{
	Items []m.NameCode
}