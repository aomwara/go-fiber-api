package models

import (
	"sandbox-api/database"

	"github.com/gofiber/fiber/v2"
)

type Logs struct {
	ID     uint   `gorm:"primarykey" json:"id"`
	Action string `json:"action"`
	UserId int64  `json:"user_id"`
}

func ListLogs(c *fiber.Ctx) error {
	db := database.DBConn
	var logs []Logs
	db.Find(&logs)
	return c.JSON(&logs)
}
