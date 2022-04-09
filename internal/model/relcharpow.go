package model

type RelCharPow struct {
	CharactersID int
	PowersID     int
}

func (r RelCharPow) TableName() string {
	return "realcharpower"
}
