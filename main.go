package main

import (
	"github.com/a-h/templ"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/adaptor"
	"github.com/kontakhaikal/goshop/repository"
	"github.com/kontakhaikal/goshop/views"
)

func main() {
	app := fiber.New()

	app.Static("/public", "./public")

	userRepository := repository.NewUserRepository()

	app.Get("/", func(c *fiber.Ctx) error {
		users := userRepository.GetUsers()
		return Render(c, views.Home(users))
	})

	app.Post("/", func(c *fiber.Ctx) error {
		name := c.FormValue("name")
		userRepository.AddUser(repository.User{
			Id:   name,
			Name: name,
		})
		return c.Redirect("/")
	})

	app.Listen(":8000")
}

func Render(c *fiber.Ctx, t templ.Component) error {
	return adaptor.HTTPHandler(templ.Handler(t))(c)
}
