package models

type HighBad struct {
	Id     int
	HighCi string
	HighIn string
	HighCm string
	HighMm string
	Price  string
}

func (m *HighBad) TableName() string {
	return TableName("high_bed")
}