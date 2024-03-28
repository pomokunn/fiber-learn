package main

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type Product struct {
	ID    int     `json:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

var products []Product

func main() {
	app := fiber.New()

	products = append(products, Product{ID: 1, Name: "Water", Price: 10.00})
	products = append(products, Product{2, "Pepsi", 20.512})

	app.Get("/product", getProducts)
	app.Get("/product/:id", getProduct)
	app.Post("/product", postProducts)

	app.Listen(":8080")
}

func getProducts(c *fiber.Ctx) error {
	return c.JSON(products)
}

func getProduct(c *fiber.Ctx) error {
	productID, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	for _, product := range products {
		if product.ID == productID {
			return c.JSON(product)
		}
	}

	return c.SendStatus(fiber.StatusNotFound)
}

func postProducts(c *fiber.Ctx) error {
	product := new(Product)
	if err := c.BodyParser(product); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}
	products = append(products, *product)
	return c.JSON(product)
}
