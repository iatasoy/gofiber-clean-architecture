package service

import (
	"context"
	"errors"
	"gofiber-clean-architecture/model"
	"gofiber-clean-architecture/repository"

	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

type AuthService interface {
	RegisterUser(ctx context.Context, username, email, password string) error
	LoginUser(ctx context.Context, email, password string) (*model.User, error)
}

type authService struct {
	userRepository repository.UserRepository
}

func NewAuthService(repo repository.UserRepository) AuthService {
	return &authService{
		userRepository: repo,
	}
}

// RegisterUser registers a new user after checking for existing emails and usernames
func (s *authService) RegisterUser(ctx context.Context, username, email, password string) error {
	// Check if the email is already registered
	existingUserByEmail, err := s.userRepository.FindByEmail(ctx, email)
	if err != nil && err != mongo.ErrNoDocuments {
		return err // Unexpected DB error
	}
	if existingUserByEmail != nil {
		return errors.New("duplicate_email")
	}

	// Check if the username is already taken
	existingUserByUsername, err := s.userRepository.FindByUsername(ctx, username)
	if err != nil && err != mongo.ErrNoDocuments {
		return err // Unexpected DB error
	}
	if existingUserByUsername != nil {
		return errors.New("duplicate_username")
	}

	// Hash the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	// Create and save the new user
	user := &model.User{
		Username: username,
		Email:    email,
		Password: string(hashedPassword),
	}

	return s.userRepository.Create(ctx, user)
}

// LoginUser checks the email and password, returning the user if successful
func (s *authService) LoginUser(ctx context.Context, email, password string) (*model.User, error) {
	user, err := s.userRepository.FindByEmail(ctx, email)
	if err != nil || user == nil {
		if err == mongo.ErrNoDocuments {
			return nil, errors.New("invalid email or password")
		}
		return nil, err
	}

	// Compare passwords
	if bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)) != nil {
		return nil, errors.New("invalid email or password")
	}

	return user, nil
}
