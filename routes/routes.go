package routes

import (
	"github.com/gofiber/fiber/v2"

	"github.com/mshore-dev/imagebucket/routes/admin"
	"github.com/mshore-dev/imagebucket/routes/api"
	"github.com/mshore-dev/imagebucket/routes/auth"
	"github.com/mshore-dev/imagebucket/routes/files"
	"github.com/mshore-dev/imagebucket/routes/gallery"
	"github.com/mshore-dev/imagebucket/routes/home"
)

// this top-level function will bring in routes from all of
// the packages below it.
func RegisterRoutes(app *fiber.App) {
	files.RegisterRoutes(app)
	gallery.RegisterRoutes(app)
	api.RegisterRoutes(app)
	admin.RegisterRoutes(app)
	auth.RegisterRoutes(app)

	// I think this needs to be last, since we (may) have to register
	// a route for /filename.ext, which would probably take priority
	// over the rest of the above routes.
	home.RegisterRoutes(app)
}
