package peoples

import (
	"fmt"
	"simple_crud/common"
	"strconv"
)

type People struct {
	NIK         string
	Name        string
	Birthday    string
	Gender      string
	Address     string
	Photo       string
	RT          string
	RW          string
	Subdistrict string
	District    string
	City        string
	Province    string
	Age         int
}

func AutoMigrate() {
	db := common.GetDB()
	err := db.AutoMigrate(&People{})
	if err != nil {
		fmt.Println(err.Error())
	}
}

func FindOne(condition interface{}) (People, error) {
	db := common.GetDB()
	var people People
	err := db.Where(condition).First(&people).Error
	if err != nil {
		return people, err
	}
	return people, nil
}

func GetPeoples(condition interface{}, limit, offset string) ([]People, int, error) {
	db := common.GetDB()

	offset_int, err := strconv.Atoi(offset)
	if err != nil {
		offset_int = 0
	}

	limit_int, err := strconv.Atoi(limit)
	if err != nil {
		limit_int = 10
	}
	var count int64

	var peoples []People
	err = db.Where(condition).Limit(limit_int).Offset(offset_int).Find(&peoples).Error
	db.Model(&[]People{}).Count(&count)
	return peoples, int(count), err
}

func Create(data interface{}) error {
	db := common.GetDB()
	err := db.Create(data).Error
	return err
}

func (people *People) Update(data interface{}) error {
	db := common.GetDB()
	err := db.Where("nik = ?", people.NIK).Save(data).Error
	return err
}

func (people *People) Delete() error {
	db := common.GetDB()
	err := db.Where("nik = ?", people.NIK).Delete(&People{}).Error
	return err
}
