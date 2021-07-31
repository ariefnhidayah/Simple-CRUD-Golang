package common

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Database struct {
	*gorm.DB
}

var DB *gorm.DB

func Init() *gorm.DB {
	HOST := "127.0.0.1"
	PASS := ""
	USER := "root"
	DBNAME := "crud_golang"
	PORT := "3306"
	// dsn := "root:@tcp(127.0.0.1:3306)/crud_golang?charset=utf8mb4&parseTime=True&loc=Local"
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", USER, PASS, HOST, PORT, DBNAME)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("Db err: (Init)", err)
	}

	fmt.Println("DB: Connected")
	DB = db

	return DB
}

func GetDB() *gorm.DB {
	return DB
}
