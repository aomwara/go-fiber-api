package models

import (
	"sandbox-api/database"

	"github.com/gofiber/fiber/v2"
)

type Users struct {
	ID    uint   `gorm:"primarykey" json:"id"`
	Email string `json:"email"`
	Point int64  `json:"point"`
}

func ListUsers(c *fiber.Ctx) error {
	db := database.DBConn
	var users []Users
	db.Find(&users)
	return c.JSON(&users)
}

func CreateUser(c *fiber.Ctx) error {
	db := database.DBConn
	user := new(Users)
	if err := c.BodyParser(user); err != nil {
		return err
	}
	db.Create(&user)
	return c.JSON(&user)
}
