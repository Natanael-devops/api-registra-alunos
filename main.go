package main

import (
	"github.com/Natanael-devops/api-go-gin/database"
	"github.com/Natanael-devops/api-go-gin/routes"
)

func main() {
	database.ConectaComBancoDeDados()
	routes.HandleRequests()
}
