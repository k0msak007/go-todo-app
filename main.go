package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/k0msak007/go-todo-app/configs"
	"github.com/k0msak007/go-todo-app/routes"
)

func main() {
	configs.BootApp()
	app := fiber.New()

	routes.InitRoute(app)

	app.Listen(":8000")
}
