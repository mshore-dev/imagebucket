package auth

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"

	"github.com/mshore-dev/imagebucket/database"
	"github.com/mshore-dev/imagebucket/middleware"
	"github.com/mshore-dev/imagebucket/utils/sessions"
)

func RegisterRoutes(app *fiber.App) {
	app.Get("/auth/login", routeGetLogin)
	app.Post("/auth/login", routePostLogin)
	app.Post("/auth/logout", middleware.RequireAuthentication, routePostLogout)

	app.Get("/auth/register", middleware.RestrictPrivateMode, routeGetRegister)
	app.Post("/auth/register", middleware.RestrictPrivateMode, routePostRegister)

	// special :D
	app.Get("/auth/test", middleware.RequireAuthentication, routeGetTest)
}

func routeGetLogin(c *fiber.Ctx) error {

	ctx := c.Context()

	if ctx.UserValue("authenticated") != nil {
		// user is already signed in
		c.Redirect("/")
	}

	c.Render("login", fiber.Map{
		"title": "Login",
	})

	return nil
}

func routePostLogin(c *fiber.Ctx) error {

	username := c.FormValue("username", "")
	password := c.FormValue("password", "")

	u, err := database.GetUserByUsername(username)
	if err != nil {
		// TODO: properly handle the "no rows" error for user's that don't exist
		log.Printf("failed to look up user \"%s\": %v\n", username, err)

		return c.Render("errors/500", fiber.Map{
			"title":   "Error",
			"message": "Failed to look up user in database.",
		})
	}

	ok, err := database.CheckUserLogin(username, password)
	if err != nil {
		// TODO: properly handle the "no rows" error for user's that don't exist
		log.Printf("failed to check password for \"%s\": %v\n", username, err)

		return c.Render("errors/500", fiber.Map{
			"title":   "Error",
			"message": "Failed to validate password.",
		})
	}

	if !ok {
		// incorrect credentials
		// TODO: properly show user an error on the login page, don't reuse 500
		return c.Render("errors/500", fiber.Map{
			"title":   "Error",
			"message": "Incorrect username or password.",
		})
	}

	// all OK?
	sess, err := sessions.Get(c)
	if err != nil {
		log.Printf("failed to get session after register: %v\n", err)

		return c.Render("errors/500", fiber.Map{
			"title":   "Error",
			"message": "Failed to get session after register.",
		})
	}

	sess.Set("authenticated", true)
	sess.Set("userid", u.ID)
	sess.Save()

	c.Redirect("/")

	return nil
}

func routePostLogout(c *fiber.Ctx) error {

	return nil
}

func routeGetRegister(c *fiber.Ctx) error {

	c.Render("register", fiber.Map{
		"title": "Register",
	})

	return nil
}

func routePostRegister(c *fiber.Ctx) error {

	// TODO: input validation

	username := c.FormValue("username", "")
	password := c.FormValue("password", "")
	password2 := c.FormValue("password2", "")

	if username == "" || password == "" || password2 == "" {
		// TODO: better error page
		return c.Render("errors/500", fiber.Map{
			"title":   "Error",
			"message": "You must fill out all fields.",
		})
	}

	if password != password2 {
		// TODO: better error page
		return c.Render("errors/500", fiber.Map{
			"title":   "Error",
			"message": "Passwords must match.",
		})
	}

	err := database.CreateUser(username, password)
	if err != nil {
		log.Printf("failed to CreateUser(\"%s\": %v\n)", username, err)

		// TODO: better error page
		return c.Render("errors/500", fiber.Map{
			"title":   "Error",
			"message": "Failed to create user.",
		})
	}

	return nil
}

func routeGetTest(c *fiber.Ctx) error {

	ctx := c.Context()

	c.SendString(fmt.Sprintf("authenticated: %b\nuserid: %d\nusername: %s\n", ctx.UserValue("authenticated"), ctx.UserValue("userid"), ctx.UserValue("username")))

	return nil
}
