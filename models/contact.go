package models

import "gorm.io/gorm"

// 人员关系
type Contact struct {
	gorm.Model
	OwnerId  uint
	TargetId uint
	Type     int
	Describe string
}

func (t *Contact) TableName() string {
	return "contact"
}
