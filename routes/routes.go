package routes

import (
	"gofiber-clean-architecture/handler"
	"gofiber-clean-architecture/middleware"
	"gofiber-clean-architecture/service"

	"github.com/gofiber/fiber/v2"
)

// Register all routes
func RegisterAllRoutes(app *fiber.App, jwtSecret string, authService service.AuthService, userService service.UserService) {
	// Create the handlers with the necessary services
	rootHandler := handler.NewRootHandler()
	authHandler := handler.NewAuthHandler(authService)
	userHandler := handler.NewUserHandler(userService)

	// Register the root route
	rootHandler.RegisterRoutes(app)

	// Register the auth routes
	authHandler.RegisterRoutes(app)

	// Register the user routes (protected with JWT)
	protected := app.Group("/user/", middleware.JWTMiddleware(jwtSecret))
	userHandler.RegisterProtectedRoutes(protected)
}
