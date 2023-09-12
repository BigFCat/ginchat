package models

import "gorm.io/gorm"

//群信息
type GroupInfo struct {
	gorm.Model
	Name     string
	OwnerId  uint
	Icon     string
	Type     int
	Describe string
}

func (t *GroupInfo) TableName() string {
	return "group_info"
}
