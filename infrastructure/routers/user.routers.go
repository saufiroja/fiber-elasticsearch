package routers

import (
	"elasticsearch/fiber-elasticsearch/controllers"
	"elasticsearch/fiber-elasticsearch/database"
	"elasticsearch/fiber-elasticsearch/repository"
	"elasticsearch/fiber-elasticsearch/service"

	"github.com/gofiber/fiber/v2"
)

func RouterUser(fiber *fiber.App, conf database.Config) {
	db := database.InitDatabase(conf)

	userRepo := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepo)
	userController := controllers.NewUserControllers(userService)

	fiber.Post("/user", userController.CreateUser)
	fiber.Get("/user", userController.FindAllUser)
	fiber.Get("/user/:id", userController.FindUserById)
	fiber.Get("/user/search/:query", userController.SearchUser)
	fiber.Get("/user/aggregation/:query", userController.AggregationUser)
	fiber.Put("/user/:id", userController.UpdateUser)
	fiber.Delete("/user/:id", userController.DeleteUser)
}
