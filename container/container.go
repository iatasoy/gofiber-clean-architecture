package container

import (
	"gofiber-clean-architecture/database"
	"gofiber-clean-architecture/handler"
	"gofiber-clean-architecture/repository"
	"gofiber-clean-architecture/service"
)

// Container holds dependencies
type Container struct {
	UserRepository repository.UserRepository
	UserService    service.UserService
	AuthService    service.AuthService
}

// NewUserContainer initializes and returns all dependencies
func NewUserContainer() (*Container, error) {
	repo := repository.NewUserRepository(database.UserCollection.Database())
	authSvc := service.NewAuthService(repo)
	userSvc := service.NewUserService(repo)

	return &Container{
		UserRepository: repo,
		UserService:    userSvc,
		AuthService:    authSvc,
	}, nil
}

// NewAuthHandler returns an AuthHandler
func NewAuthHandler(authService service.AuthService) *handler.AuthHandler {
	return handler.NewAuthHandler(authService)
}

// NewUserHandler returns a UserHandler
func NewUserHandler(userService service.UserService) *handler.UserHandler {
	return handler.NewUserHandler(userService)
}
