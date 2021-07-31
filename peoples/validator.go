package peoples

import (
	"simple_crud/common"
	"time"

	"github.com/gin-gonic/gin"
)

type PeopleModelValidator struct {
	PeopleValidator struct {
		NIK         string `form:"nik" json:"nik" binding:"required"`
		Name        string `form:"name" json:"name" binding:"required"`
		Birthday    string `form:"birthday" json:"birthday"`
		Gender      string `form:"gender" json:"gender" binding:"required"`
		Address     string `form:"address" json:"address"`
		Photo       string `form:"photo" json:"photo" binding:"omitempty,url"`
		RT          string `form:"rt" json:"rt"`
		RW          string `form:"rw" json:"rw"`
		Subdistrict string `form:"subdistrict" json:"subdistrict"`
		District    string `form:"district" json:"district"`
		City        string `form:"city" json:"city"`
		Province    string `form:"province" json:"province"`
		Age         int    `form:"age" json:"age"`
	} `json:"people"`
	people People
}

func (peopleModelValidator *PeopleModelValidator) Bind(c *gin.Context) error {
	err := common.Bind(c, peopleModelValidator)
	if err != nil {
		return err
	}

	peopleModelValidator.people.NIK = peopleModelValidator.PeopleValidator.NIK
	peopleModelValidator.people.Name = peopleModelValidator.PeopleValidator.Name
	peopleModelValidator.people.Birthday = peopleModelValidator.PeopleValidator.Birthday
	peopleModelValidator.people.Gender = peopleModelValidator.PeopleValidator.Gender
	peopleModelValidator.people.Address = peopleModelValidator.PeopleValidator.Address
	peopleModelValidator.people.RT = peopleModelValidator.PeopleValidator.RT
	peopleModelValidator.people.RW = peopleModelValidator.PeopleValidator.RW
	peopleModelValidator.people.Subdistrict = peopleModelValidator.PeopleValidator.Subdistrict
	peopleModelValidator.people.District = peopleModelValidator.PeopleValidator.District
	peopleModelValidator.people.City = peopleModelValidator.PeopleValidator.City
	peopleModelValidator.people.Province = peopleModelValidator.PeopleValidator.Province
	peopleModelValidator.people.Age = peopleModelValidator.PeopleValidator.Age
	peopleModelValidator.people.Photo = peopleModelValidator.PeopleValidator.Photo

	return nil
}

func NewPeopleModelValidator() PeopleModelValidator {
	peopleModelValidator := PeopleModelValidator{}
	return peopleModelValidator
}

func NewPeopleModelValidatorFillWith(people People) PeopleModelValidator {
	peopleModelValidator := NewPeopleModelValidator()

	dateParse, _ := time.Parse(time.RFC3339, people.Birthday)
	people.Birthday = dateParse.Format("2006-01-02")

	peopleModelValidator.PeopleValidator.NIK = people.NIK
	peopleModelValidator.PeopleValidator.Name = people.Name
	peopleModelValidator.PeopleValidator.Birthday = people.Birthday
	peopleModelValidator.PeopleValidator.Gender = people.Gender
	peopleModelValidator.PeopleValidator.Address = people.Address
	peopleModelValidator.PeopleValidator.RT = people.RT
	peopleModelValidator.PeopleValidator.RW = people.RW
	peopleModelValidator.PeopleValidator.Subdistrict = people.Subdistrict
	peopleModelValidator.PeopleValidator.District = people.District
	peopleModelValidator.PeopleValidator.City = people.City
	peopleModelValidator.PeopleValidator.Province = people.Province
	peopleModelValidator.PeopleValidator.Age = people.Age
	peopleModelValidator.PeopleValidator.Photo = people.Photo

	return peopleModelValidator
}
