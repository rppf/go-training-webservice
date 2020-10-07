package main

import (
	"fmt"

	"github.com/rppf/go-training-webservice/models"
)

func main() {
	u := models.User{
		ID:        2,
		FirstName: "Raymond",
		LastName:  "Florendo",
	}
	fmt.Println(u)
}
