package routes

import (
	"sandbox-api/services"

	"github.com/gofiber/fiber/v2"
)

func Handle(app *fiber.App) {
	app.Get("/", services.GetHome)

	//User route
	app.Get("/users", services.ListUsers)
	app.Post("/users", services.CreateUser)

	//Auth route
	app.Post("/auth", Auth)

	//protect route;
	//AuthorizationRequired Action
	app.Use(AuthorizationRequired())
	//Log route
	app.Get("/logs", services.ListLogs)
}
