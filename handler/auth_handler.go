package handler

import (
	"gofiber-clean-architecture/dto"
	"gofiber-clean-architecture/middleware"
	"gofiber-clean-architecture/model"
	"gofiber-clean-architecture/service"
	"log"

	"github.com/gofiber/fiber/v2"
)

const (
	RegisterRoute       = "/auth/register"
	LoginRoute          = "/auth/login"
	LoginSuccessMessage = "Login successful"
	RegisterSuccessMsg  = "User registered successfully"
)

type AuthHandler struct {
	service service.AuthService
}

func NewAuthHandler(service service.AuthService) *AuthHandler {
	return &AuthHandler{service: service}
}

func (h *AuthHandler) RegisterRoutes(app *fiber.App) {
	app.Post(RegisterRoute, h.RegisterUser)
	app.Post(LoginRoute, h.LoginUser)
}

func (h *AuthHandler) RegisterUser(c *fiber.Ctx) error {
	var req dto.RegisterRequest

	// Parse body
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(model.APIResponse{
			Status:  "error",
			Message: "Malformed JSON in request payload",
		})
	}

	// Basic required fields check
	if req.Username == "" || req.Email == "" || req.Password == "" {
		return c.Status(fiber.StatusBadRequest).JSON(model.APIResponse{
			Status:  "error",
			Message: "All fields (username, email, password) are required",
		})
	}

	// Service call
	if err := h.service.RegisterUser(c.Context(), req.Username, req.Email, req.Password); err != nil {
		switch err.Error() {
		case "duplicate_username":
			return c.Status(fiber.StatusConflict).JSON(model.APIResponse{
				Status: "error",
				Error:  "Username already exists",
			})
		case "duplicate_email":
			return c.Status(fiber.StatusConflict).JSON(model.APIResponse{
				Status: "error",
				Error:  "Email already exists",
			})
		default:
			return c.Status(fiber.StatusInternalServerError).JSON(model.APIResponse{
				Status: "error",
				Error:  "An unexpected error occurred",
			})
		}
	}

	// Success response
	return c.Status(fiber.StatusCreated).JSON(model.APIResponse{
		Status:  "success",
		Message: RegisterSuccessMsg,
	})
}

func (h *AuthHandler) LoginUser(c *fiber.Ctx) error {
	var req dto.LoginRequest

	// Parse body
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(model.APIResponse{
			Status:  "error",
			Message: "Malformed JSON in request payload",
		})
	}

	// Basic required fields check
	if req.Email == "" || req.Password == "" {
		return c.Status(fiber.StatusBadRequest).JSON(model.APIResponse{
			Status:  "error",
			Message: "Both email and password are required",
		})
	}

	// Call service to login user
	user, err := h.service.LoginUser(c.Context(), req.Email, req.Password)
	if err != nil {
		log.Printf("Login error: %v", err)
		return c.Status(fiber.StatusUnauthorized).JSON(model.APIResponse{
			Status: "error",
			Error:  "Invalid email or password",
		})
	}

	if user == nil {
		return c.Status(fiber.StatusInternalServerError).JSON(model.APIResponse{
			Status: "error",
			Error:  "An unexpected error occurred",
		})
	}

	jwtToken, err := middleware.GenerateJWT(user.ID.Hex(), user.Username)

	if err != nil {
		log.Printf("Token generation error: %v", err)
		return c.Status(fiber.StatusInternalServerError).JSON(model.APIResponse{
			Status: "error",
			Error:  "Failed to generate token",
		})
	}

	// Success response with user data
	return c.Status(fiber.StatusOK).JSON(model.APIResponse{
		Status:  "success",
		Message: LoginSuccessMessage,
		Data: map[string]interface{}{
			"userid":   user.ID,
			"username": user.Username,
			"email":    user.Email,
			"token":    jwtToken,
		},
	})
}
