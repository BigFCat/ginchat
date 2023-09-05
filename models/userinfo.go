package models

import (
	"fmt"
	"ginchat/utils"
	"time"

	"gorm.io/gorm"
)

//user_infos table
type UserInfo struct {
	gorm.Model
	Name          string
	PassWord      string
	Phone         string
	Email         string
	Identify      string
	ClientIp      string
	ClientPort    string
	LoginTime     time.Time `gorm:"default:2006-05-04 15:02:01"`
	LoginOutTime  time.Time `gorm:"default:2006-05-04 15:02:01"`
	HeartbeatTime time.Time `gorm:"default:2006-05-04 15:02:01"`
	IsLogout      bool
	DeviceInfo    string
}

func (t *UserInfo) TableName() string {
	return "user_infos"

}

func GetUserList() []*UserInfo {
	data := make([]*UserInfo, 10)
	utils.DB.Find(&data)

	for _, v := range data {
		fmt.Println(v)
	}
	return data
}

func CreateUser(user UserInfo) *gorm.DB {
	return utils.DB.Create(&user)

}

func GetUser(user UserInfo) int {
	var users int64
	c := len(utils.DB.Where("name=?", user.Name).Count(&users).Name())
	return c
}

func DeleteUser(user UserInfo) *gorm.DB {
	return utils.DB.Where("name = ?", user.Name).Delete(&user)
}

func UpdateUser(user UserInfo) *gorm.DB {
	return utils.DB.Model(&user).UpdateColumns(&UserInfo{Name: user.Name, PassWord: user.PassWord})
}
