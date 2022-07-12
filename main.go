package main

import (
	"alura-go-restapi/database"
	"alura-go-restapi/routes"
	"fmt"
)

func main() {
	database.ConexaoDB()

	fmt.Println("Iniciando servidor Go")
	routes.HandleRequest()
}
