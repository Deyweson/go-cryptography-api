package main

import (
	"fmt"

	"github.com/deyweson/go-cryptography-api/database"
	"github.com/deyweson/go-cryptography-api/routes"
)

func main() {
	fmt.Println("Hello")
	database.DatabaseInitializer()
	routes.HandleRequests()
}
