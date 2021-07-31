package users

import (
	"simple_crud/common"

	"github.com/gin-gonic/gin"
)

type UserModelValidator struct {
	UserValidator struct {
		Name     string `form:"name" json:"name" binding:"required"`
		Username string `form:"username" json:"username" binding:"required"`
		Password string `form:"password" json:"password" binding:"required,min=8,max=255"`
	} `json:"user"`
	user User `json:"-"`
}

func (userModelValidator *UserModelValidator) Bind(c *gin.Context) error {
	err := common.Bind(c, userModelValidator)
	if err != nil {
		return err
	}

	userModelValidator.user.Username = userModelValidator.UserValidator.Username
	userModelValidator.user.Name = userModelValidator.UserValidator.Name
	userModelValidator.user.Role = "user"

	if userModelValidator.UserValidator.Password != common.NBRandomPassword {
		userModelValidator.user.setPassword(userModelValidator.UserValidator.Password)
	}
	return nil
}

func NewUserModelValidator() UserModelValidator {
	userModelValidator := UserModelValidator{}
	return userModelValidator
}

func NewUserModelValidatorFillWith(user User) UserModelValidator {
	userModelValidator := NewUserModelValidator()
	userModelValidator.UserValidator.Username = user.Username
	userModelValidator.UserValidator.Name = user.Name
	userModelValidator.UserValidator.Password = common.NBRandomPassword

	return userModelValidator
}

type LoginValidator struct {
	Login struct {
		Username string `form:"username" json:"username" binding:"required"`
		Password string `form:"password" json:"password" binding:"required,min=8,max=255"`
	} `json:"user"`
	user User
}

func (loginValidator *LoginValidator) Bind(c *gin.Context) error {
	err := common.Bind(c, loginValidator)
	if err != nil {
		return err
	}
	loginValidator.Login.Username = loginValidator.user.Username
	return nil
}

func NewLoginValidator() LoginValidator {
	loginValidator := LoginValidator{}
	return loginValidator
}
