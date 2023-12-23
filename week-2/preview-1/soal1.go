package main

import "fmt"

type Data struct {
	Nama      string
	Alamat    string
	Pekerjaan string
	Alasan    string
}

func TampilkanData(data []Data) {
	for _, v := range data {
		fmt.Printf("Nama: %s\nAlamat: %s\nPekerjaan: %s\nAlasan memilih kelas golang: %s\n",
			v.Nama, v.Alamat, v.Pekerjaan, v.Alasan)
	}
}
