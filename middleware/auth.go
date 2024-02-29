package middleware

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/mshore-dev/imagebucket/database"
	"github.com/mshore-dev/imagebucket/utils/sessions"
)

func GetSession(c *fiber.Ctx) error {

	sess, err := sessions.Get(c)
	if err != nil {
		log.Printf("failed to get session: %v\n", err)

		return c.Render("errors/500", fiber.Map{
			"title":   "Error",
			"message": "Could not look up your session.",
		})
	}

	requestContext := c.Context()

	userid := sess.Get("userid")

	requestContext.SetUserValue("authenticated", sess.Get("authenticated"))
	requestContext.SetUserValue("userid", userid)

	if userid != nil {
		// userid is set, so the user *should* be logged in
		// TODO: double-check there's no weird edge cases here.

		u, err := database.GetUserByID(userid.(int))
		if err != nil {
			log.Printf("failed to get user for context: %v\n", err)

			return c.Render("errors/500", fiber.Map{
				"title":   "Error",
				"message": "Could not look up your account in db.",
			})
		}

		requestContext.SetUserValue("username", u.Username)
	}

	return c.Next()
}

func RequireAuthentication(c *fiber.Ctx) error {

	requestContext := c.Context()

	if requestContext.UserValue("authenticated") == nil {
		// user is not authenticated
		return c.Redirect("/auth/login")
	}

	// fall through
	return c.Next()
}
