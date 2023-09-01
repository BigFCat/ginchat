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
	LoginTime     time.Time
	LoginOutTime  time.Time `gorm:"cloumn:login_out_time default:null json:'login_out_time'"`
	HeartbeatTime time.Time
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
