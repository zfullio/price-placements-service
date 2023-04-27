package entity

type AreaType string

const (
	LIVING   AreaType = "Жилая недвижимость"
	COMMERCE AreaType = "Коммерческая недвижимость"
)

var AreaTypesAll = []AreaType{LIVING, COMMERCE}
