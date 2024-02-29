package home

import (
	"log"

	"github.com/gofiber/fiber/v2"

	"github.com/mshore-dev/imagebucket/config"
	"github.com/mshore-dev/imagebucket/database"
)

func RegisterRoutes(app *fiber.App) {
	app.Get("/", routeHome)

	// only enable is the user has requsted it
	if config.Config.ServeFiles {
		log.Println("Serving uploads directly")
		app.Static("/", config.Config.Uploads.Folder)
	}

}

func routeHome(c *fiber.Ctx) error {

	files, err := database.GetFilesByUser(1, 2, 1)
	if err != nil {
		panic(err)
	}

	c.Render("home", fiber.Map{
		"title": "Home",
		"files": files,
	})

	return nil
}
