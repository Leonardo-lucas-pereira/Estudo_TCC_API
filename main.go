package main

import (
	"fmt"

	"github.com/Leonardo-lucas-pereira/tcc-api/database"
	"github.com/Leonardo-lucas-pereira/tcc-api/server"
)

func main() {
	fmt.Println("Iniciando servidor....")
	database.StartDB()

	server := server.NewServer()
	server.Run()
}
