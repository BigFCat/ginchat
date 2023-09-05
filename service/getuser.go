package service

import (
	"fmt"
	"ginchat/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

// GetPage 获取用户信息
// @Summary 获取用户信息
// @Tags 用户模块
// @Description 获取用户信息
// @Success 200 {string} json:"{"code": 200, "message": [...]}"
// @Router /user/getUserList [get]
func GetUserList(c *gin.Context) {
	datas := make([]*models.UserInfo, 10)
	datas = models.GetUserList()
	c.JSON(200, gin.H{
		"message": datas,
	})

}

// @Summary 创建用户
// @Tags 用户模块
// @Description 创建用户
// @param name query string false "用户名"
// @param password query string false "密码"
// @param repassword query string false "确认密码"
// @Success 200 {string} json:"{"code": 200, "message": [...]}"
// @Router /user/createUser [get]
func CreateUser(c *gin.Context) {
	user := models.UserInfo{}
	user.Name = c.Query("name")
	password := c.Query("password")
	reassword := c.Query("repassword")
	if password != reassword {
		c.JSON(-1, gin.H{
			"message": "密码不一致",
		})
		return
	}
	user.PassWord = password
	if models.GetUser(user) > 0 {
		fmt.Println("name重复数量", models.GetUser(user))
		c.JSON(-1, gin.H{
			"message": "用户已存在",
		})
		return
	}

	models.CreateUser(user)
	c.JSON(200, gin.H{
		"message": "新增用户成功",
	})

}

// @Summary 删除用户
// @Tags 用户模块
// @Description 删除用户
// @param name query string false "用户名"
// @Success 200 {string} json:"{"code": 200, "message": [...]}"
// @Router /user/deleteUser [get]
func DeleteUser(c *gin.Context) {
	user := models.UserInfo{}
	user.Name = c.Query("name")
	if models.GetUser(user) <= 0 {
		fmt.Println(models.GetUser(user))
		c.JSON(-1, gin.H{
			"message": "用户不存在!",
		})
		return
	}
	models.DeleteUser(user)
	c.JSON(200, gin.H{
		"message": "删除用户成功",
	})

}

// @Summary 修改用户
// @Tags 用户模块
// @Description 修改用户
// @param id formData string false "id"
// @param name formData string false "用户名"
// @param password formData string false "密码"
// @Success 200 {string} json:"{"code": 200, "message": [...]}"
// @Router /user/updateUser [post]
func UpdateUser(c *gin.Context) {
	user := models.UserInfo{}
	id, _ := strconv.Atoi(c.PostForm("id"))
	user.ID = uint(id)
	user.Name = c.PostForm("name")
	user.PassWord = c.PostForm("password")

	models.UpdateUser(user)
	c.JSON(200, gin.H{
		"message": "修改用户成功",
	})

}
