package models

import (
	"database/sql"
	"fmt"
	"ginchat/utils"

	"gorm.io/gorm"
)

//user_infos table
type UserInfo struct {
	gorm.Model
	Name          string
	PassWord      string
	Phone         string `valid:"matches(^1[3-9]{1}\\d{9}$)"`
	Email         string `valid:"email"`
	Identify      string
	ClientIp      string
	ClientPort    string
	Salt          string
	LoginTime     sql.NullTime
	LoginOutTime  sql.NullTime
	HeartbeatTime sql.NullTime
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

func FindUserByPhone(phone string) *gorm.DB {
	user := UserInfo{}
	return utils.DB.Where("Phone=?", phone).First(&user)
}
func FindUserByName(name string) UserInfo {
	user := UserInfo{}
	utils.DB.Where("Name=?", name).First(&user)
	return user
}
func FindUserByNameAndPwd(name, password string) UserInfo {
	user := UserInfo{}
	utils.DB.Where("Name=? and pass_word=?", name, password).First(&user)
	return user
}

// 根据ID、Name、Phone、email字段，查询数据
func GetUserInfo(user *UserInfo) []*UserInfo {
	u_db := utils.DB.Model(&user)
	var infos []*UserInfo
	if user.ID > 0 {
		u_db.Where("ID=?", user.ID)
	}
	if user.Name != "" {
		u_db.Where("name=?", user.Name)
	}
	if user.Phone != "" {
		u_db.Where("phone=?", user.Phone)
	}
	if user.Email != "" {
		u_db.Where("email=?", user.Email)
	}
	u_db.Find(&infos)
	return infos

}

func DeleteUser(user UserInfo) *gorm.DB {
	return utils.DB.Where("name = ?", user.Name).Delete(&user)
}

func UpdateUser(user UserInfo) *gorm.DB {
	return utils.DB.Model(&user).UpdateColumns(&UserInfo{Name: user.Name, PassWord: user.PassWord, Phone: user.Phone, Email: user.Email})
}
