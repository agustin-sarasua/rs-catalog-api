package app

type Neighbourhood struct {
	name    string
	city    string
	country string
}

var Neighbourhoods = []*Neighbourhood{
	&Neighbourhood{name: "Pocitos", city: "MVD", country: "UY"},
	&Neighbourhood{name: "Punta Carretas", city: "MVD", country: "UY"},
	&Neighbourhood{name: "Centro", city: "MVD", country: "UY"},
	&Neighbourhood{name: "Tres Cruces", city: "MVD", country: "UY"},
}

var PropertyTypes = []string{"APARTAMENTO", "CASA",
	"PENTHOUSE", "CAMPO", "LOCAL_COMERCIAL", "PROPIEDAD_HORIZONTAL", "TERRENO", "GALPON"}

var Amenities = []string{"JACUZZI", "POOL", "BBQ", "GARAGE", "PORTERIA_24", "VIGILANCIA", "GYM"}

// const Cities = []string{"MVD", "MDES", "ART",}
