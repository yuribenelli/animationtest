package model

type Powers struct {
	ID          int
	Name        string
	Description string
	Damage      int
}

func (p Powers) TableName() string {
	return "powers"
}
