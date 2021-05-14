package main

import (
	"io"
	"log"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	// GET test for Jikan API -> https://jikan.docs.apiary.io/#
	app.Get("/:test", func(c *fiber.Ctx) error {
		var test string
		resp, err := http.Get("https://api.jikan.moe/v3")
		if err != nil {
			panic(err)
		}
		defer resp.Body.Close()
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			panic(err)
		}
		test = string(body)
		return c.SendString(test) // => Hello john 👋!
	})

	log.Fatal(app.Listen(":3000"))
}
