package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/kontakhaikal/goshop/repository"
)

func RenderProductsHandler(repository repository.ProductRepository) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		products, err := repository.FindAll()
		if err != nil {
			return ctx.SendStatus(fiber.StatusInternalServerError)
		}
		return ctx.Render("products", fiber.Map{
			"Title":    "Products",
			"Products": products,
		}, "layouts/main")
	}
}

func RenderProductDetailHandler(repository repository.ProductRepository) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		slug := ctx.Params("slug")
		product, err := repository.FindBySlug(slug)
		if err != nil {
			return ctx.SendStatus(fiber.StatusNotFound)
		}
		return ctx.Render("product-detail", fiber.Map{
			"Title":   "product-detail",
			"Product": product,
		}, "layouts/main")
	}
}
