package users

import (
	"errors"
	"net/http"
	"simple_crud/common"

	"github.com/gin-gonic/gin"
)

func UsersRegister(router *gin.RouterGroup) {
	router.POST("/login", userLogin)
	// router.Use(AuthMiddleware(true, true))
	router.POST("/", userRegistration)
}

func userRegistration(c *gin.Context) {
	userModelValidator := NewUserModelValidator()

	// check field entity
	if err := userModelValidator.Bind(c); err != nil {
		c.JSON(http.StatusUnprocessableEntity, common.NewValidatorError(err))
		return
	}

	isUsernameAvail := isUsernameAvailable(userModelValidator)

	if isUsernameAvail {
		c.JSON(http.StatusConflict, common.NewError("error", errors.New("Username Already Registered!")))
		return
	}

	if err := SaveOne(&userModelValidator.user); err != nil {
		c.JSON(http.StatusUnprocessableEntity, common.NewError("error", err))
		return
	}

	c.Set("user", userModelValidator.user)
	serializer := UserSerializer{c}
	// c.JSON(http.StatusCreated, gin.H{"status": "success", "message": "Success", "data": serializer.Response()})
	c.JSON(http.StatusCreated, common.ResponseApi("Success", serializer.Response()))
}

func isUsernameAvailable(inputUser UserModelValidator) bool {
	user, err := FindOneUser(&User{Username: inputUser.UserValidator.Username})
	if err != nil {
		return false
	}

	if user.ID != 0 {
		return true
	}

	return false
}

func userLogin(c *gin.Context) {
	loginValidator := NewLoginValidator()

	// check field entity
	if err := loginValidator.Bind(c); err != nil {
		c.JSON(http.StatusUnprocessableEntity, common.NewValidatorError(err))
		return
	}

	// Check username
	user, err := FindOneUser(&User{Username: loginValidator.Login.Username})
	if err != nil {
		c.JSON(http.StatusForbidden, common.NewError("error", errors.New("Not Registered username or invalid password")))
		return
	}

	// check password
	if user.checkPassword(loginValidator.Login.Password) != nil {
		c.JSON(http.StatusForbidden, common.NewError("error", errors.New("Not Registered username or invalid password")))
		return
	}

	UpdateContextUserModel(c, user.ID)
	serializer := UserSerializer{c}
	c.JSON(http.StatusOK, common.ResponseApi("Success", serializer.Response()))
}
