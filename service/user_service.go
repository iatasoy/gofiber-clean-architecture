package service

import (
	"context"
	"gofiber-clean-architecture/model"
	"gofiber-clean-architecture/repository"
)

type UserService struct {
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) UserService {
	return UserService{repo: repo}
}

// GetUserByID retrieves the user from the database by their ID
func (s *UserService) GetUserByID(ctx context.Context, userID string) (*model.User, error) {
	// You may need to cast userID to an appropriate type if it’s not a string
	return s.repo.GetUserById(ctx, userID)
}
