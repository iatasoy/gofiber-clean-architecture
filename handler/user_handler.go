package handler

import (
	"gofiber-clean-architecture/model"
	"gofiber-clean-architecture/service"

	"github.com/gofiber/fiber/v2"
)

type UserHandler struct {
	service service.UserService
}

func NewUserHandler(service service.UserService) *UserHandler {
	return &UserHandler{service: service}
}

func (h *UserHandler) RegisterProtectedRoutes(router fiber.Router) {
	router.Post("/:id", h.GetUserByID)
}

func (h *UserHandler) GetUserByID(c *fiber.Ctx) error {
	requestedID := c.Params("id")
	authenticatedID, ok := c.Locals("userid").(string)

	if !ok || authenticatedID != requestedID {
		return c.Status(fiber.StatusForbidden).JSON(model.APIResponse{
			Status:  "error",
			Message: "Access denied",
			Error:   "You are not authorized to access this resource",
		})
	}

	user, err := h.service.GetUserByID(c.Context(), requestedID)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(model.APIResponse{
			Status:  "error",
			Message: "User not found",
			Error:   err.Error(),
		})
	}

	return c.JSON(model.APIResponse{
		Status:  "success",
		Message: "User fetched successfully",
		Data: map[string]interface{}{
			"userid":   user.ID,
			"username": user.Username,
			"email":    user.Email,
		},
	})
}
