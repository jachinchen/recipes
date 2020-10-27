package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	// Routes
	app.Get("/foo", handler)
	app.Get("/", hello)

	// app.Use(func(c *fiber.Ctx) error {
	// 	result := utils.ImmutableString(c.Params("foo"))
	// 	return c.SendString(result)
	// })

	// app.Get("/:value", func(c *fiber.Ctx) error {
	// 	return c.SendString("value: " + c.Params("value"))
	// })
	// app.Get("/:name/:age", func(c *fiber.Ctx) error {
	// 	if c.Params("name") != "" {
	// 		return c.SendString("Hello " + c.Params("name") + ", age: " + c.Params("age"))
	// 	}
	// 	return c.SendString("Where is raymond?")
	// })
	// app.Get("/api/*", func(c *fiber.Ctx) error {
	// 	return c.SendString("API path: " + c.Params("*"))
	// })

	app.Static("/", "./public")

	// Start server
	log.Fatal(app.Listen(":3000"))
}

func hello(c *fiber.Ctx) error {
	return c.SendString("Hello, World!")
}

func handler(c *fiber.Ctx) error {
	result := c.Params("foo")
	fmt.Println("1" + result)
	buffer := make([]byte, len(result))
	copy(buffer, result)
	resultCopy := string(buffer)
	test := "what"

	return c.SendString(test + resultCopy)
}
