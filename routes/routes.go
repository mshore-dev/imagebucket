package routes

import (
	"github.com/gofiber/fiber/v2"

	"github.com/mshore-dev/imagebucket/routes/admin"
	"github.com/mshore-dev/imagebucket/routes/api"
	"github.com/mshore-dev/imagebucket/routes/files"
	"github.com/mshore-dev/imagebucket/routes/gallery"
	"github.com/mshore-dev/imagebucket/routes/home"
)

func RegisterRoutes(app *fiber.App) {
	files.RegisterRoutes(app)
	gallery.RegisterRoutes(app)
	api.RegisterRoutes(app)
	admin.RegisterRoutes(app)

	// I think this needs to be last, since we (may) have to register
	// a route for /<fileid>.<ext>, which would probably take priority
	// over the rest of the above routes.
	home.RegisterRoutes(app)
}
