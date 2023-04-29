package entity

type Phone struct {
	Number         int
	Developer      string
	Object         string
	ObjectExtended []string
	Placement      Placement
}

func (p Phone) CheckNumber(obj string, num int) bool {
	if p.Object == obj && p.Number == num {
		return true
	}

	return false
}
