package entity

type CheckResult struct {
	Developer string
	Placement Placement
	Base      Placement
	Url       string
	Message   string
	Status    StatusCheck
}
