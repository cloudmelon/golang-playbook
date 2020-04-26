package main

import (
	"cloudmelon/webservice/models"
	"fmt"
)

func main() {
	u := models.User{
		ID:        2,
		FirstName: "Lucie",
		LastName:  "Cat",
	}
	fmt.Println(u)
}
