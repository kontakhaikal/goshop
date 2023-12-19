package main

import (
	"log"

	"github.com/a-h/templ"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/adaptor"
	"github.com/kontakhaikal/goshop/model"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	db, err := gorm.Open(sqlite.Open("data/test.db"), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}

	if err := db.AutoMigrate(
		&model.User{},
		&model.Address{},
		&model.Category{},
		&model.Product{},
		&model.ProductImage{},
		&model.Section{},
	); err != nil {
		log.Fatalln(err)
	}

	app := fiber.New()

	if err := app.Listen(":8000"); err != nil {
		log.Fatalln(err)
	}
}

func Render(c *fiber.Ctx, t templ.Component) error {
	return adaptor.HTTPHandler(templ.Handler(t))(c)
}
