package server

import (
	"elasticsearch/fiber-elasticsearch/database"
	"elasticsearch/fiber-elasticsearch/infrastructure/routers"

	"github.com/gofiber/fiber/v2"
)

func Server() *fiber.App {
	app := fiber.New()
	conf := database.Config{}

	routers.RouterUser(app, conf)

	return app
}
