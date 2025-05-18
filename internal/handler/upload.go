package handler

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
)

func UploadCOBOLFile(c *fiber.Ctx) error {
	file, err := c.FormFile("file")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("No file uploaded")
	}

	// Save the file to a temporary location
	savePath := fmt.Sprintf("./tmp/%s", file.Filename)
	if err := c.SaveFile(file, savePath); err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Failed to save file")
	}

	// TODO: Call COBOL analyzer here
	return c.JSON(fiber.Map{
		"message": "File uploaded successfully",
		"file":    file.Filename,
	})
}