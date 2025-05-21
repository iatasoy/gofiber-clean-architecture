package mapper

import (
	"gofiber-clean-architecture/dto"
	"gofiber-clean-architecture/model"
)

func ToUserResponse(user model.User) dto.UserResponse {
	return dto.UserResponse{
		ID:       user.ID.Hex(),
		Username: user.Username,
		Email:    user.Email,
	}
}
