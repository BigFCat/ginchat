package main

import (
	"fmt"
	"ginchat/models"
	"ginchat/utils"
)

func main() {
	utils.InitConfig()
	utils.InitMySQL()

	utils.DB.AutoMigrate(&models.UserInfo{})

	//create
	user := &models.UserInfo{Name: "hqy3", PassWord: "111", Phone: "1233"}
	utils.DB.Create(user)

	//read
	fmt.Println(utils.DB.First(user))
}
