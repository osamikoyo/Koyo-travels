package models

type Sort struct {
	Rait bool
	Alf  bool
}

type HotelFilter struct {
	MaxPrice uint64
	City     string
	District string
}

type Filter struct {
	Country        string
	MoreThanItRait float32
	HotelFilter HotelFilter
}
