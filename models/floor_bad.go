package models

type FloorBad struct {
	Id     int
	HighCi string
	HighIn string
	HighCm string
	HighMm string
	Price  string
}

func (m *FloorBad) TableName() string {
	return TableName("floor_bed")
}