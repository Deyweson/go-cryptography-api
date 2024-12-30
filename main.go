package main

import (
	"github.com/deyweson/go-cryptography-api/database"
	"github.com/deyweson/go-cryptography-api/routes"
)

func main() {

	database.DatabaseInitializer()
	routes.HandleRequests()

}
