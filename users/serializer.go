package users

import (
	"simple_crud/common"

	"github.com/gin-gonic/gin"
)

type UserSerializer struct {
	c *gin.Context
}

type UserResponse struct {
	Username string `json:"username"`
	Name     string `json:"name"`
	Role     string `json:"role"`
	Token    string `json:"token"`
}

func (mine *UserSerializer) Response() UserResponse {
	model := mine.c.MustGet("user").(User)
	user := UserResponse{
		Username: model.Username,
		Name:     model.Name,
		Role:     model.Role,
		Token:    common.GenToken(model.ID, model.Role),
	}
	return user
}
