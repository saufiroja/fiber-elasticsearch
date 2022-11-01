package interfaces

import (
	"elasticsearch/fiber-elasticsearch/entity"

	"github.com/gofiber/fiber/v2"
)

type UserRepository interface {
	CreateUser(user *entity.User) (entity.User, error)
	FindAllUser() (map[string]any, error)
	FindUserById(id string) (map[string]any, error)
	UpdateUser(id string, user *entity.User) error
	DeleteUser(id string) error
	SearchUser(query string) (map[string]any, error)
}

type UserService interface {
	CreateUser(user *entity.User) (entity.User, error)
	FindAllUser() (map[string]any, error)
	FindUserById(id string) (map[string]any, error)
	UpdateUser(id string, user *entity.User) error
	DeleteUser(id string) error
	SearchUser(query string) (map[string]any, error)
}

type UserControllers interface {
	CreateUser(ctx *fiber.Ctx) error
	FindAllUser(ctx *fiber.Ctx) error
	FindUserById(ctx *fiber.Ctx) error
	UpdateUser(ctx *fiber.Ctx) error
	DeleteUser(ctx *fiber.Ctx) error
	SearchUser(ctx *fiber.Ctx) error
}
