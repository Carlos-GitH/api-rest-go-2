package main

import (
	"github.com/Carlos-GitH/api-rest-gin-2/database"
	"github.com/Carlos-GitH/api-rest-gin-2/routes"
)

func main() {
	database.ConectaComBancoDeDados()
	routes.HandleRequest()
}
