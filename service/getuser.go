package service

import (
	"ginchat/models"

	"github.com/gin-gonic/gin"
)

// GetPage 获取用户信息
// @Summary 获取用户信息
// @Description 获取用户信息
// @Tags 用户信息
// @Success 200 {string} json:"{"code": 200, "message": [...]}"
// @Router /user/getUserList [get]
func GetUserList(c *gin.Context) {
	data := make([]*models.UserInfo, 10)
	data = models.GetUserList()
	c.JSON(200, gin.H{
		"message": data,
	})

}
