package models

type Dotes struct{
	Title string
	Description string
}

type Address struct {
	City string
	District string
	Street string
	HouseHumber uint16
}

type Hotel struct {
	PriceInDay uint64
	Description string
	Address 	Address
}

type Excurs struct{
	Dotes Dotes
	Description string
	Title string
}

type Travel struct {
	Excurs []Excurs
	Description string
}