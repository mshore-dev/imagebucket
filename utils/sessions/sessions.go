package sessions

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
)

var store *session.Store

func init() {

	// TODO: allow options to be configured
	store = session.New()

}

func Get(c *fiber.Ctx) (*session.Session, error) {
	return store.Get(c)
}
