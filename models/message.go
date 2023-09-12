package models

import (
	"gorm.io/gorm"
)

//user_infos table
type Messagge struct {
	gorm.Model
	FormID   string //发送者
	TargetID string //接收者
	Type     string //消息类型 群聊私聊广播
	Media    int    //消息内容
	Pic      string //图片
	Url      string //url
	Desc     string //描述
	Amount   int
}

func (t *Messagge) TableName() string {
	return "messages"

}
