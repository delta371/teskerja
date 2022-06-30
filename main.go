package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	app.Get("/", func(c *fiber.Ctx) error {
		var response []string
		n := 1
		i := 1
		for {
			if i > 99 {
				break
			}

			response = append(response, fmt.Sprintf("%d  X  %d  = %d", n, i, n*i))
			i++
		}

		return c.JSON(response)
	})

	log.Fatal(app.Listen(":3000"))
}
