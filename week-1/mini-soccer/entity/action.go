package entity

import (
	"fmt"
	"strconv"
	"unicode"
)

func IsString(input string) bool {
	for _, v := range input {
		if !unicode.IsLetter(v) {
			return false
		}
	}

	return true
}

func (l *Lapangan) ListLapangan(input []Lapangan) {
	fmt.Println("Daftar Lapangan:")
	for _, v := range input {
		fmt.Printf("Lapangan: %s\nJenis Rumput: %s\nHarga: Rp.%d / jam\n", v.Nama, v.Rumput, v.Perjam)
	}
}

func (b *Booking) GetBooking(input []Lapangan) string {
	dataBooking := []Booking{}

	var inputCustomer string
	fmt.Println("Masukkan nama:")
	fmt.Scanln(&inputCustomer)

	var timeBook string
	fmt.Println("Booking berapa jam:")
	fmt.Scanln(&timeBook)

	if !IsString(inputCustomer) {
		return "Invalid input: Masukkan nama dengan huruf"
	}

	jamBooking, err := strconv.Atoi(timeBook)
	if err != nil {
		return "Invalid input: Masukkan waktu booking dengan angka"
	}

	dataBooking = append(dataBooking, dataBooking...)

	return fmt.Sprintf("Booking berhasil\nNama: %s\nLapangan: %s\nJenis Rumput: %s\nHarga: Rp.%d / jam\nBooking: %d jam\nTotal: Rp.%d", b.Customer, b.Lapangan.Nama, b.Lapangan.Rumput, b.Lapangan.Perjam, jamBooking, b.Lapangan.Perjam*jamBooking)
}
