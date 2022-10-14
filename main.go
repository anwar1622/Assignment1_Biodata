package main

import (
	"Biodata/controllers"
	"Biodata/models"
	"fmt"
	"os"
)

func main() {
	var args = os.Args

	fmt.Println(controllers.GetPersonById(args[1], models.Person{}))
}
