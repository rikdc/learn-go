package main

import (
	"fmt"

	"github.com/pluralsight/webservice/models"
)

func main() {
	u := models.User{
		ID:        2,
		FirstName: "Peter",
		LastName:  "Pan",
	}
	fmt.Println(u)

	port := 3000
	_, err := startWebServer(port)
	fmt.Println(err)
}

func startWebServer(port int) (int, error) {
	fmt.Println("Starting webserver...")
	fmt.Println("Server started on port", port)

	return port, nil
}
