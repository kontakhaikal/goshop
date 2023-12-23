package main

import (
	"log"
	"os"

	"github.com/kontakhaikal/goshop/handler"
	"github.com/kontakhaikal/goshop/repository"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
	"github.com/kontakhaikal/goshop/faker"
	"github.com/kontakhaikal/goshop/model"
	"github.com/urfave/cli/v2"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {

	db, err := gorm.Open(sqlite.Open("data/test.db"), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}

	productRepository := repository.NewGormProductRepository(db)

	c := &cli.App{
		Action: func(ctx *cli.Context) error {
			return nil
		},
		Commands: []*cli.Command{
			{
				Name: "db:migrate",
				Action: func(ctx *cli.Context) error {
					return RunMigration(db)
				},
			},
			{
				Name: "db:seed",
				Action: func(ctx *cli.Context) error {
					return RunSeeder(db)
				},
			},
		},
	}

	if err := c.Run(os.Args); err != nil {
		log.Fatalln(err)
	}

	engine := html.New("views", ".html")

	app := fiber.New(fiber.Config{
		Views: engine,
	})

	app.Static("/public/", "public")

	app.Get("/", handler.RenderHomeHandler())

	app.Get("/products", handler.RenderProductsHandler(productRepository))

	app.Get("/products/:slug", handler.RenderProductDetailHandler(productRepository))

	if err := app.Listen(":8000"); err != nil {
		log.Fatalln(err)
	}
}

func RunMigration(db *gorm.DB) error {
	return db.AutoMigrate(
		&model.User{},
		&model.Address{},
		&model.Category{},
		&model.Product{},
		&model.ProductImage{},
		&model.Section{},
		&model.Order{},
		&model.OrderCustomer{},
		&model.OrderItem{},
		&model.Shipment{},
	)
}

func RunSeeder(db *gorm.DB) error {

	user := faker.UserFaker()
	product := faker.ProductFaker(user)

	seeders := []any{
		user,
		product,
	}

	for _, seeder := range seeders {
		if err := db.Debug().Create(seeder).Error; err != nil {
			return err
		}
	}

	return nil
}
