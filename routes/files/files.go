package files

import (
	"log"
	"path"

	"github.com/gofiber/fiber/v2"
	"github.com/mshore-dev/imagebucket/config"
	"github.com/mshore-dev/imagebucket/database"
	"github.com/mshore-dev/imagebucket/middleware"
	"github.com/mshore-dev/imagebucket/utils"
)

func RegisterRoutes(app *fiber.App) {
	app.Post("/files/upload", middleware.RequireAuthentication, routePostFilesUpload)
}

func routePostFilesUpload(c *fiber.Ctx) error {

	file, err := c.FormFile("file")
	if err != nil {
		log.Printf("failed to get formfile from request: %v\n", err)

		return c.Render("errors/500", fiber.Map{
			"title":   "Error",
			"message": "Could not find file in POST request.",
		})
	}

	localFilename := utils.GenerateID() + path.Ext(file.Filename)

	log.Printf("got file to /files/upload! name: %s | size: %d | db filename: %s\n", file.Filename, file.Size, localFilename)

	err = c.SaveFile(file, path.Join(config.Config.Uploads.Folder, localFilename))
	if err != nil {
		log.Printf("failed to save uploaded file: %v\n", err)
	}

	err = database.CreateFile(localFilename, file.Filename, "", c.Context().UserValue("userid").(int))
	if err != nil {
		log.Printf("failed to add file to db: %v\n", err)
	}

	return c.Redirect("/")
}
