package main

import (
	"simple_crud/common"
	"simple_crud/peoples"
	"simple_crud/users"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) {
	users.AutoMigrate()
	peoples.AutoMigrate()
}

func main() {
	common.Init()
	// Migrate(db)

	r := gin.Default()

	v1 := r.Group("/api")
	users.UsersRegister(v1.Group("/users"))
	v1.Use(users.AuthMiddleware(true, false))
	peoples.PeopleRegister(v1.Group("/peoples"))

	r.Run(":8888")
}
