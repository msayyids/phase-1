package main

import (
	"mini-soccer/entity"
)

func main() {

	data := []entity.Lapangan{
		{Nama: "old trafford", Rumput: "sintetis", Perjam: 700000},
		{Nama: "old trafford", Rumput: "sintetis", Perjam: 700000},
		{Nama: "old trafford", Rumput: "sintetis", Perjam: 700000},
	}

	l := entity.Lapangan{}
	l.ListLapangan(data)
	b := entity.Booking{}
	b.GetBooking(data)

}
