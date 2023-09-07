package service

import (
	"fmt"
	"ginchat/models"
	"ginchat/utils"
	"math/rand"
	"strconv"
	"time"

	"github.com/asaskevich/govalidator"
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
// @param phone query string false "手机号"
// @param email query string false "邮箱"
// @Success 200 {string} json:"{"code": 200, "message": [...]}"
// @Router /user/createUser [get]
func CreateUser(c *gin.Context) {
	user := models.UserInfo{}
	user.Name = c.Query("name")
	password := c.Query("password")
	reassword := c.Query("repassword")
	user.Phone = c.Query("phone")
	user.Email = c.Query("email")
	rand.Seed(time.Now().Unix())
	salt := fmt.Sprintf("%06d", rand.Int31())
	_, err := govalidator.ValidateStruct(user)
	if err != nil {
		fmt.Println(err)
		c.JSON(-1, gin.H{
			"message": "err",
		})
		return
	} else {
		data := models.FindUserByName(user.Name)
		if data.Name != "" {
			// if len(models.GetUserInfo(&user)) > 0 {
			c.JSON(-1, gin.H{
				"message": "用户已存在",
			})
			return
		}

		if password != reassword {
			c.JSON(-1, gin.H{
				"message": "密码不一致",
			})
			return
		}
		user.PassWord = utils.Md5AddSalt(password, salt)
		user.Salt = salt
		// user.PassWord = password

		models.CreateUser(user)
		c.JSON(200, gin.H{
			"message": "新增用户成功",
		})
	}
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
	if len(models.GetUserInfo(&user)) <= 0 {
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
// @param phone formData string false "手机号"
// @param email formData string false "邮箱"
// @Success 200 {string} json:"{"code": 200, "message": [...]}"
// @Router /user/updateUser [post]
func UpdateUser(c *gin.Context) {
	user := models.UserInfo{}
	id, _ := strconv.Atoi(c.PostForm("id"))
	user.ID = uint(id)

	if len(models.GetUserInfo(&user)) <= 0 {
		c.JSON(-1, gin.H{
			"message": "用户不存在！",
		})
		return
	}
	user.Name = c.PostForm("name")
	user.PassWord = c.PostForm("password")
	user.Phone = c.PostForm("phone")
	user.Email = c.PostForm("email")
	_, err := govalidator.ValidateStruct(user)
	if err != nil {
		fmt.Println(err)
		c.JSON(-1, gin.H{
			"message": "err",
		})
		return
	} else {
		models.UpdateUser(user)
		c.JSON(200, gin.H{
			"message": "修改用户成功",
		})
	}
}

// @Summary 用户登入
// @Tags 用户模块
// @Description 用户登入
// @param name formData string false "用户名"
// @param password formData string false "密码"
// @Success 200 {string} json:"{"code": 200, "message": [...]}"
// @Router /user/userLogin [post]
func UserLogin(c *gin.Context) {
	user := models.UserInfo{}
	user.Name = c.PostForm("name")
	password := c.PostForm("password")

	userinfo := models.FindUserByName(user.Name)
	if userinfo.Name != "" {
		if utils.ValidPassword(password, userinfo.Salt, userinfo.PassWord) {
			r := models.FindUserByNameAndPwd(user.Name, userinfo.PassWord)
			c.JSON(200, gin.H{
				"message": r,
			})
		} else {
			c.JSON(-1, gin.H{
				"message": "密码错误",
			})
		}

	} else {
		c.JSON(200, gin.H{
			"message": "用户不存在",
		})
	}
}
