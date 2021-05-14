package main

import (
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func main() {
	var term string
	fmt.Scanln(&term)
	app := fiber.New()
	// GET test for Jikan API -> https://jikan.docs.apiary.io/#
	app.Get("/:test", func(c *fiber.Ctx) error {
		var result string
		resp, err := http.Get("https://api.jikan.moe/v3/search/anime?q=" + term)
		if err != nil {
			panic(err)
		}
		defer resp.Body.Close()
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			panic(err)
		}
		result = string(body)
		return c.SendString(result) // => Hello john 👋!
	})

	log.Fatal(app.Listen(":3000"))
}
