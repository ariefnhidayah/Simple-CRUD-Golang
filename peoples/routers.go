package peoples

import (
	"errors"
	"net/http"
	"simple_crud/common"

	"github.com/gin-gonic/gin"
)

func PeopleRegister(router *gin.RouterGroup) {
	router.GET("/", getPeoples)
	router.POST("/", createPeople)
	router.GET("/:nik", getPeople)
	router.PUT("/:nik", updatePeople)
	router.DELETE("/:nik", deletePeople)
}

func createPeople(c *gin.Context) {
	// validator
	peopleModelValidator := NewPeopleModelValidator()
	if err := peopleModelValidator.Bind(c); err != nil {
		c.JSON(http.StatusUnprocessableEntity, common.NewValidatorError(err))
		return
	}

	// check available nik
	if check := isAvailableNIK(peopleModelValidator.people.NIK); check {
		c.JSON(http.StatusConflict, common.NewError("error", errors.New("NIK Already Registered!")))
		return
	}

	if err := Create(&peopleModelValidator.people); err != nil {
		c.JSON(http.StatusUnprocessableEntity, common.NewError("error", err))
		return
	}

	serializer := PeopleSerializer{peopleModelValidator.people}
	c.JSON(http.StatusCreated, common.ResponseApi("Success", serializer.Response()))
}

func isAvailableNIK(NIK string) bool {
	people, err := FindOne(&People{NIK: NIK})
	if err != nil {
		return false
	}

	if people.NIK != "0" {
		return true
	}

	return false
}

func getPeoples(c *gin.Context) {
	limit := c.Query("limit")
	offset := c.Query("offset")
	peoples, peopleCount, err := GetPeoples(&People{}, limit, offset)
	if err != nil {
		c.JSON(http.StatusNotFound, common.NewError("error", errors.New("Not Found!")))
		return
	}
	serializer := PeoplesSerializer{peoples}
	c.JSON(http.StatusOK, common.ResponseApi("Success", gin.H{"peoples": serializer.Response(), "count": peopleCount}))
}

func getPeople(c *gin.Context) {
	nik := c.Param("nik")

	people, err := FindOne(&People{NIK: nik})
	if err != nil {
		c.JSON(http.StatusNotFound, common.NewError("error", errors.New("People Not Found!")))
		return
	}

	serializer := PeopleSerializer{people}
	c.JSON(http.StatusOK, common.ResponseApi("Success", serializer.Response()))
}

func updatePeople(c *gin.Context) {
	nik := c.Param("nik")
	people, err := FindOne(&People{NIK: nik})
	if err != nil {
		c.JSON(http.StatusNotFound, common.NewError("error", errors.New("People Not Found!")))
		return
	}

	peopleModelValidator := NewPeopleModelValidatorFillWith(people)
	if err := peopleModelValidator.Bind(c); err != nil {
		c.JSON(http.StatusUnprocessableEntity, common.NewValidatorError(err))
		return
	}

	peopleModelValidator.people.NIK = people.NIK
	if err := people.Update(&peopleModelValidator.people); err != nil {
		c.JSON(http.StatusUnprocessableEntity, common.NewError("error", err))
		return
	}

	serializer := PeopleSerializer{people}
	c.JSON(http.StatusOK, common.ResponseApi("Success", serializer.Response()))
}

func deletePeople(c *gin.Context) {
	nik := c.Param("nik")
	people, err := FindOne(&People{NIK: nik})
	if err != nil {
		c.JSON(http.StatusNotFound, common.NewError("error", errors.New("People Not Found!")))
		return
	}

	if err := people.Delete(); err != nil {
		c.JSON(http.StatusUnprocessableEntity, common.NewError("error", err))
	}

	c.JSON(http.StatusOK, common.ResponseApi("Success", gin.H{}))
}
