package handler

import (
	"github.com/gofiber/fiber/v2"
	"fmt"
	"os"
)

func GetDocumentation(c *fiber.Ctx) error {
	filename := c.Params("filename")
	path := fmt.Sprintf("./docs/%s.md", filename)

	content, err := os.ReadFile(path)
	if err != nil {
		return c.Status(fiber.StatusNotFound).SendString("Documentation not found")
	}

	return c.SendString(string(content))
}