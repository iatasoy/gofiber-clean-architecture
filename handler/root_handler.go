package handler

import (
	"gofiber-clean-architecture/model"

	"github.com/gofiber/fiber/v2"
)

type RootHandler struct{}

func NewRootHandler() *RootHandler {
	return &RootHandler{}
}

func (h *RootHandler) RegisterRoutes(app *fiber.App) {
	app.Get("/", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON(model.APIResponse{
			Status:  "success",
			Message: "Welcome to the GoFiber Clean Architecture API!",
		})
	})
}
