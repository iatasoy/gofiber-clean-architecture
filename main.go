package main

import (
	"context"
	"log"
	"os"

	"gofiber-clean-architecture/configuration"
	"gofiber-clean-architecture/container"
	"gofiber-clean-architecture/database"
	"gofiber-clean-architecture/handler"
	"gofiber-clean-architecture/routes"

	"github.com/gofiber/fiber/v2"
)

func main() {
	// Load environment configuration
	configuration.LoadConfig()

	// Connect to MongoDB
	if err := database.Connect(); err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err := database.Mg.Client.Disconnect(context.TODO()); err != nil {
			log.Fatalf("DB disconnect error: %v", err)
		}
	}()

	// Create Fiber app
	app := fiber.New()

	// Load JWT secret
	jwtSecret := os.Getenv("JWT_SECRET")
	if jwtSecret == "" {
		log.Fatal("JWT_SECRET is not set")
	}

	// Initialize DI container
	c, err := container.NewUserContainer()
	if err != nil {
		log.Fatal("Container init failed:", err)
	}

	// Register root route
	rootHandler := handler.NewRootHandler()
	rootHandler.RegisterRoutes(app)

	// Register all application routes (auth + protected user routes)
	routes.RegisterAllRoutes(app, jwtSecret, c.AuthService, c.UserService)

	// Start server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Printf("Starting server on port %s", port)
	log.Fatal(app.Listen(":" + port))
}
