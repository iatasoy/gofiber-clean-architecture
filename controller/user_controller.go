package controller

import (
	"gofiber-clean-architecture/model"
	"gofiber-clean-architecture/service"
	"gofiber-clean-architecture/validators"

	"github.com/gofiber/fiber/v2"
)

type UserController struct {
	authService service.AuthService
}

func NewUserController(authService service.AuthService) *UserController {
	return &UserController{authService: authService}
}

func (controller *UserController) Route(app *fiber.App) {
	app.Post("/auth/register", controller.Register)
	app.Post("/auth/login", controller.Login)
}

func (controller *UserController) Register(c *fiber.Ctx) error {
	var request model.User

	if err := c.BodyParser(&request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid JSON body",
		})
	}

	if err := validators.ValidateStruct(request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	if err := controller.authService.RegisterUser(c.Context(), request.Username, request.Email, request.Password); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "User registered successfully",
	})
}

func (controller *UserController) Login(c *fiber.Ctx) error {
	var request model.User

	if err := c.BodyParser(&request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid JSON body",
		})
	}

	if err := validators.ValidateStruct(request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	user, err := controller.authService.LoginUser(c.Context(), request.Email, request.Password)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Invalid email or password",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Login successful",
		"user":    user,
	})
}
