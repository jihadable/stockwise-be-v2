package utils

import (
	"github.com/jihadable/stockwise-be/model/entity"
	"github.com/jihadable/stockwise-be/model/request"
	"github.com/jihadable/stockwise-be/model/response"
)

func RequestToUser(request *request.UserRequest) *entity.User {
	return &entity.User{
		Username: request.Username,
		Email:    request.Email,
		Password: request.Password,
		Bio:      request.Bio,
	}
}

func UserToResponse(user *entity.User) *response.UserResponse {
	return &response.UserResponse{
		Id:       user.Id,
		Username: user.Username,
		Email:    user.Email,
		Bio:      user.Bio,
	}
}

func UpdateUserRequestToUser(updateUserRequest *request.UpdateUserRequest) *entity.User {
	return &entity.User{
		Username: updateUserRequest.Username,
		Bio:      updateUserRequest.Bio,
	}
}
