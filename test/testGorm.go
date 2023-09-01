package main

import (
	"fmt"
	"ginchat/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dsn := "root:123456@tcp(127.0.0.1:3306)/ginchat?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("failed to connect db")
	}

	db.AutoMigrate(&models.UserInfo{})

	//create
	user := &models.UserInfo{Name: "hqy1", PassWord: "111", Phone: "1233"}
	// user := &models.UserInfo{}
	// user.Name = "hqy"
	db.Create(user)

	//read
	fmt.Println(db.First(user, 1))
}
