package main

import (
	"log"
	"search-cep/src/repositories"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	app.Get("/:cep", func(c *fiber.Ctx) error {
		cep := c.Params("cep")
		data, err := repositories.GetCep(cep)
		if err != nil {
			log.Fatal(err)
			return c.SendStatus(500)
		}

		if data.Uf == "" {
			return c.SendStatus(404)
		}
		return c.JSON(data)
	})

	app.Listen(":5000")

}
