package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"ucd-platform/internal/handler"
)

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Welcome to UCD Platform - COBOL Analyzer")
	})

	app.Post("/upload", handler.UploadCOBOLFile)
	app.Get("/docs/:filename", handler.GetDocumentation)

	log.Fatal(app.Listen(":3000"))
}