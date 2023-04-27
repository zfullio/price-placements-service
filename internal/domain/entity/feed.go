package entity

type Feed struct {
	Developer string
	Placement Placement
	AreaType  AreaType
	Status    StatusFeed
	Url       string
	BaseFeed  Placement
}
