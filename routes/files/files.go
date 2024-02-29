package files

import (
	"log"
	"path"

	"github.com/gofiber/fiber/v2"
	"github.com/mshore-dev/imagebucket/config"
	"github.com/mshore-dev/imagebucket/database"
	"github.com/mshore-dev/imagebucket/utils"
)

func RegisterRoutes(app *fiber.App) {
	app.Post("/files/upload", routeFilesUpload)
}

func routeFilesUpload(c *fiber.Ctx) error {

	file, err := c.FormFile("file")
	if err != nil {
		c.Render("errors/500", fiber.Map{
			"message": err.Error(),
		})
		panic(err)
	}

	localFilename := utils.GenerateID() + path.Ext(file.Filename)

	log.Printf("got file to /files/upload! name: %s | size: %d | db filename: %s\n", file.Filename, file.Size, localFilename)

	err = c.SaveFile(file, path.Join(config.Config.Uploads.Folder, localFilename))
	if err != nil {
		log.Printf("failed to save uploaded file: %v\n", err)
	}

	err = database.CreateFile(localFilename, file.Filename, "", 1)
	if err != nil {
		log.Printf("failed to add file to db: %v\n", err)
	}

	return nil
}
