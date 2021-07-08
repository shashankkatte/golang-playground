package main

import (
	"fmt"

	"github.com/shashankkatte/golang-playground/05_webservice/models"
)

func main() {

	u := models.User{
		ID:        2,
		FirstName: "Albert",
		LastName:  "Einstein",
	}
	fmt.Println(u)
}
