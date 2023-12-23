package handler

import "github.com/gofiber/fiber/v2"

func RenderHomeHandler() fiber.Handler {
	return func(c *fiber.Ctx) error {
		return c.Render("home", fiber.Map{
			"Title": "Hello, world 2",
		}, "layouts/main")
	}
}
