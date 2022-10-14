package controllers

import (
	"Biodata/models"
	"fmt"
)

func GetPersonById(id string, People models.Person) string {
	var conditional bool = false
	var peopleIndex = 0

	for index, value := range models.People {
		if value.Id == id {
			conditional = true
			peopleIndex = index
			break
		}
	}

	if conditional {
		return fmt.Sprint("Data person ditemukan:", "\n", models.People[peopleIndex])
	} else {
		return "Data tidak ditemukan !"
	}
}
