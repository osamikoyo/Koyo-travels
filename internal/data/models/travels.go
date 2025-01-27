package models

import "time"

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

type Review struct {
	Username string
	Content string
	CreatedAt time.Time
}

type Travel struct {
	Title string
	CreatedAd time.Time
	Hotel Hotel
	Excurs []Excurs
	Description string
}
