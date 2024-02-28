package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/mshore-dev/imagebucket/routes/image"
)

func RegisterRoutes(app *fiber.App) {
	image.RegisterRoutes(app)
}
