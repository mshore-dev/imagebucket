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

	ctx := c.Context()

	if ctx.UserValue("authenticated") == true {

		files, err := database.GetFilesByUser(ctx.UserValue("userid").(int), 10, 0)
		if err != nil {
			panic(err)
		}

		return c.Render("home", fiber.Map{
			"title":         "Home",
			"authenticated": ctx.UserValue("authenticated"),
			"username":      ctx.UserValue("username"),
			"private":       config.Config.Private,
			"files":         files,
		})
	}

	return c.Render("home", fiber.Map{
		"title":   "Home",
		"private": config.Config.Private,
	})
}
