package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/django/v3"
	"log"
	"time"
)

func main() {
	engine := django.New("./templates", ".django")
	app := fiber.New(fiber.Config{Views: engine})

	app.Static("/static", "./static")
	app.Get("/", indexGet)
	app.Get("/fun", funGet)
	app.Get("/infinite", infiniteGet)

	log.Fatal(app.Listen(":3000"))
}

func indexGet(c *fiber.Ctx) error {
	var roark string
	if time.Now().Hour() < 12 {
		roark = "before roark"
	} else {
		roark = "after roark"
	}

	return c.Render("index", fiber.Map{
		"time": roark,
	})
}

func infiniteGet(c *fiber.Ctx) error {
	return c.Render("infinite", nil)
}

func funGet(c *fiber.Ctx) error {
	return c.Render("fun", nil)
}
