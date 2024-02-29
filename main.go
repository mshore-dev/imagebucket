package main

import (
	"flag"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/handlebars/v2"

	"github.com/mshore-dev/imagebucket/config"
	"github.com/mshore-dev/imagebucket/database"
	"github.com/mshore-dev/imagebucket/middleware"
	"github.com/mshore-dev/imagebucket/routes"
	"github.com/mshore-dev/imagebucket/utils"
)

func main() {

	configFile := flag.String("config", "config.toml", "Path the config file")
	flag.Parse()

	log.Println("imagebucket testing")

	config.LoadConfig(*configFile)

	database.OpenDB()

	utils.Setup()

	app := fiber.New(fiber.Config{
		Views: handlebars.New("./assets/templates", ".hbs"),
	})

	app.Static("/static", "./assets/static")

	// GetSession middleware adds authentication information to the request
	// context, so that we can easily pull it later, without having to add
	// extra complexity with sharing the session store.
	app.Use(middleware.GetSession)

	routes.RegisterRoutes(app)

	log.Fatal(app.Listen(":8080"))

}
