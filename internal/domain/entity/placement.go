package entity

type Placement string

const (
	UnknownPlacement Placement = "unknown"
	Cian             Placement = "циан"
	Realty           Placement = "я.недвижимость"
	Avito            Placement = "авито"
	M2               Placement = "м2"
	DomClick         Placement = "домклик"
	JCat             Placement = "jcat"
	CalugaHouse      Placement = "калугахаус"
	FlatOutlet       Placement = "flatoutlet"
	GdeTotDom        Placement = "gdeetotdom"
	Hybrid           Placement = "hybrid"
	Avaho            Placement = "avaho"
	NovostroyM       Placement = "новострой-м"
	ClientFeed       Placement = "клиентский фид"
)

var PlacementsAll = []Placement{
	Cian,
	Realty,
	Avito,
	M2,
	DomClick,
	JCat,
	CalugaHouse,
	FlatOutlet,
	GdeTotDom,
	Hybrid,
	Avaho,
	NovostroyM,
	ClientFeed,
}
