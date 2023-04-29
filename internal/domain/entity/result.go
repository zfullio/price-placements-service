package entity

type CheckResult struct {
	Developer string
	Placement Placement
	Base      Placement
	URL       string
	Message   string
	Status    StatusCheck
}
