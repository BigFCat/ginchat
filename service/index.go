package service

import "github.com/gin-gonic/gin"

// GetPage 首页
// @Summary 首页
// @Tags 首页信息
// @Description 首页
// @Success 200 {string} json:"{"code": 200, "message": [...]}"
// @Router /index [get]
func GetIndex(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "hello gin",
	})

}
