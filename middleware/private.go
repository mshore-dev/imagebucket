package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/mshore-dev/imagebucket/config"
)

func RestrictPrivateMode(c *fiber.Ctx) error {
	if config.Config.Private {
		// TODO: don't use 500 here, bad :(
		return c.Render("errors/500", fiber.Map{
			"title":   "Error",
			"message": "This instance is in private mode, you cannot access this page.",
		})
	}

	return c.Next()
}
