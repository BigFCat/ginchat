package main

import (
	"ginchat/models"
	"ginchat/utils"
)

func main() {
	utils.InitConfig()
	utils.InitMySQL()

	// utils.DB.AutoMigrate(&models.UserInfo{})
	utils.DB.AutoMigrate(&models.Messagge{})
	utils.DB.AutoMigrate(&models.Contact{})
	utils.DB.AutoMigrate(&models.GroupInfo{})

	// //create
	// user := &models.UserInfo{Name: "hqy3", PassWord: "111", Phone: "1233"}
	// // utils.DB.Create(user)

	// //read
	// fmt.Println(utils.DB.Select([]string{"name"}).Find(&user))
	// fmt.Println(utils.DB.Table("user_infos").Select([]string{"name"}).Scan(&user))
	// // utils.DB.Where(models.UserInfo{Name: user.Name}).FirstOrInit(&user)
	// var count int64
	// utils.DB.Model(&user).Where("name = ?", user.Name).Count(&count)

	// fmt.Println(count)
}
