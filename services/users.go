package services

import (
	"sandbox-api/models"

	"github.com/gofiber/fiber/v2"
)

func CreateUser(c *fiber.Ctx) error {
	return models.CreateUser(c)
}

func ListUsers(c *fiber.Ctx) error {
	return models.ListUsers(c)
}
