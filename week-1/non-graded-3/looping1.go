package main

import "fmt"

func Looping1() {
	data := []map[string]interface{}{
		{
			"name": "Hank",
			"age":  50,
			"job":  "polisi",
		},
		{
			"name": "Heisenberg",
			"age":  52,
			"job":  "Ilmuwan",
		},
		{
			"name": "Skyler",
			"age":  48,
			"job":  "akuntan",
		},
	}

	for _, output := range data {
		nama := output["name"]
		umur := output["age"]
		pekerjaan := output["job"]

		fmt.Printf("Hi Perkenalkan, Nama saya %s, umur saya %d tahun, dan saya bekerja sebagai %s \n", nama, umur, pekerjaan)
	}

}
