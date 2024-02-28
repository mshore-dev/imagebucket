package main

import (
	"flag"
	"log"

	"github.com/gofiber/fiber/v2"

	"github.com/mshore-dev/imagebucket/config"
	"github.com/mshore-dev/imagebucket/routes"
)

func main() {

	configFile := flag.String("config", "config.toml", "Path the config file")
	flag.Parse()

	log.Println("imagebucket testing")

	config.LoadConfig(*configFile)

	app := fiber.New()

	app.Static("/static", "./assets/static")

	routes.RegisterRoutes(app)

}
