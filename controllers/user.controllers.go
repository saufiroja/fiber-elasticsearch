package controllers

import (
	"elasticsearch/fiber-elasticsearch/entity"
	"elasticsearch/fiber-elasticsearch/interfaces"

	"github.com/gofiber/fiber/v2"
)

type UserControllers struct {
	UserService interfaces.UserService
}

func NewUserControllers(userService interfaces.UserService) interfaces.UserControllers {
	return &UserControllers{
		UserService: userService,
	}
}

func (u *UserControllers) CreateUser(ctx *fiber.Ctx) error {
	user := entity.User{}
	data := ctx.BodyParser(&user)

	if data != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Bad Request",
		})
	}

	user, err := u.UserService.CreateUser(&user)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Internal Server Error",
		})
	}

	return ctx.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "User Created",
		"data":    user,
	})
}

func (u *UserControllers) FindAllUser(ctx *fiber.Ctx) error {
	users, err := u.UserService.FindAllUser()
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Internal Server Error",
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Success",
		"data":    users,
	})
}

func (u *UserControllers) DeleteUser(ctx *fiber.Ctx) error {
	id := ctx.Params("id")

	err := u.UserService.DeleteUser(id)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "User Deleted",
	})
}

func (u *UserControllers) FindUserById(ctx *fiber.Ctx) error {
	id := ctx.Params("id")

	user, err := u.UserService.FindUserById(id)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Success",
		"data":    user,
	})
}

func (u *UserControllers) UpdateUser(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	user := entity.User{}
	data := ctx.BodyParser(&user)

	if data != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Bad Request",
		})
	}

	err := u.UserService.UpdateUser(id, &user)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "User Updated",
	})
}

func (u *UserControllers) SearchUser(ctx *fiber.Ctx) error {
	users, err := u.UserService.SearchUser()
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Success",
		"data":    users,
	})
}

func (u *UserControllers) AggregationUser(ctx *fiber.Ctx) error {
	users, err := u.UserService.AggregationUser()
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Success",
		"data":    users,
	})
}
