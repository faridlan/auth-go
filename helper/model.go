package helper

import (
	"github.com/faridlan/auth-go/model/domain"
	"github.com/faridlan/auth-go/model/web"
)

func ToUserResponse(user *domain.User) *web.UserResponse {
	return &web.UserResponse{
		ID:        user.ID,
		Username:  user.Username,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
}

func ToUserResponseLogin(user *domain.User) *web.UserResponseLogin {
	return &web.UserResponseLogin{
		User: &web.UserResponse{
			ID:        user.ID,
			Username:  user.Username,
			CreatedAt: user.CreatedAt,
			UpdatedAt: user.UpdatedAt,
		},
	}
}

func ToUserResponses(users []*domain.User) []*web.UserResponse {
	userResponses := []*web.UserResponse{}

	for _, user := range users {
		userResponses = append(userResponses, ToUserResponse(user))
	}

	return userResponses

}
