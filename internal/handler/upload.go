package handler

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"os"
	"path/filepath"
)


func UploadHandler(c *fiber.Ctx) error {
	// Parse the form and get the uploaded file
	fileHeader, err := c.FormFile("file")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "File not found in request",
		})
	}

	// Validate file extension (e.g., only .cbl or .cob allowed)
	if filepath.Ext(fileHeader.Filename) != ".cbl" && filepath.Ext(fileHeader.Filename) != ".cob" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid file type. Only .cbl and .cob files are allowed.",
		})
	}

	// Ensure tmp directory exists
	if err := os.MkdirAll("./tmp", os.ModePerm); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Could not create tmp directory",
		})
	}

	// Save the uploaded file to ./tmp/
	dstPath := fmt.Sprintf("./tmp/%s", fileHeader.Filename)
	if err := c.SaveFile(fileHeader, dstPath); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Could not save file",
		})
	}

	// Return success with file name
	return c.JSON(fiber.Map{
		"message": "File uploaded successfully",
		"filename": fileHeader.Filename,
	})
}