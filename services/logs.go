package services

import (
	"sandbox-api/models"

	"github.com/gofiber/fiber/v2"
)

func ListLogs(c *fiber.Ctx) error {
	return models.ListLogs(c)
}
