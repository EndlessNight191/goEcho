package main

import (
	"goEcho/configs"
	"goEcho/internal/database"
	"goEcho/internal/routes"
	"log"
)

func main() {
	if err := configs.GetConfig(); err != nil {
		log.Fatalf("error trying initialize config: %v", err)
	}

	if err := database.CreateConnection(); err != nil {
		log.Fatal(err)
	}

	routes.InitRoutes()
}
