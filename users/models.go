package users

import (
	"errors"
	"fmt"
	"simple_crud/common"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	// ID       int    `gorm:"primaryKey;column:id"`
	// Name     string `gorm:"column:name"`
	// Username string `gorm:"column:username"`
	// Password string `gorm:"column:password"`
	// Role     string `gorm:"column:role"`
	ID       int
	Name     string
	Username string
	Password string
	Role     string
}

func AutoMigrate() {
	db := common.GetDB()

	err := db.AutoMigrate(&User{})
	if err != nil {
		fmt.Printf(err.Error())
	}
}

func (user *User) setPassword(password string) error {
	if len(password) == 0 {
		return errors.New("password should not be empty")
	}
	bytePassword := []byte(password)
	passwordHash, _ := bcrypt.GenerateFromPassword(bytePassword, bcrypt.DefaultCost)
	user.Password = string(passwordHash)
	return nil
}

func (user *User) checkPassword(password string) error {
	bytePassword := []byte(password)
	byteHashedPassword := []byte(user.Password)

	return bcrypt.CompareHashAndPassword(byteHashedPassword, bytePassword)
}

func FindOneUser(condition interface{}) (User, error) {
	db := common.GetDB()
	var user User
	err := db.Where(condition).First(&user).Error
	if err != nil {
		return user, err
	}
	return user, nil
}

func SaveOne(data interface{}) error {
	db := common.GetDB()
	err := db.Create(data).Error
	return err
}

func (user *User) Update(data interface{}) error {
	db := common.GetDB()
	err := db.Where("id = ?", user.ID).Save(data).Error
	return err
}

func (user *User) Delete() error {
	db := common.GetDB()
	err := db.Where("id = ?", user.ID).Delete(&User{}).Error
	return err
}
