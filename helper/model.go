package helper

import "github.com/faridlan/auth-go/model"

func ToUserResponse(user *model.UserHash) *model.UserResponse {
	return &model.UserResponse{
		ID:       user.ID,
		Username: user.Username,
	}
}

func ToUserResponseLogin(user *model.UserHash) *model.UserResponseLogin {
	return &model.UserResponseLogin{
		User: &model.UserResponse{
			ID:       user.ID,
			Username: user.Username,
		},
	}
}

func ToUserResponses(users []*model.UserHash) []*model.UserResponse {
	userResponses := []*model.UserResponse{}

	for _, user := range users {
		userResponses = append(userResponses, ToUserResponse(user))
	}

	return userResponses

}
