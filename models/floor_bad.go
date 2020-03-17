package models

type FloorBad struct {
	Id int
	HighCi string
}

func (m *FloorBad) TableName() string {
	return TableName("floor_bed")
}