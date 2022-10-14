package models

type Person struct {
	Id        string
	Nama      string
	Alamat    string
	Pekerjaan string
	Alasan    string
}

var People []Person = []Person{
	{
		Id:        "1",
		Nama:      "Muhamad Chaerul Anwar",
		Alamat:    "Jakarta",
		Pekerjaan: "Software Enginer",
		Alasan:    "Menyenangkan",
	},
	{
		Id:        "2",
		Nama:      "Gilang",
		Alamat:    "Bandung",
		Pekerjaan: "Software Enginer",
		Alasan:    "Hobby",
	},
	{
		Id:        "3",
		Nama:      "Vindy",
		Alamat:    "Jakarta",
		Pekerjaan: "Software Enginer",
		Alasan:    "Karir",
	},
}
