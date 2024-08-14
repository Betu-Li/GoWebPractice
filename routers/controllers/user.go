package controllers

import (
	"GinExample2/models"
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"strconv"
)

type UserController struct{}

// Register 注册
func (uc UserController) Register(c *gin.Context) {
	username := c.DefaultPostForm("username", "")
	password := c.DefaultPostForm("password", "")
	confirmPassword := c.DefaultPostForm("confirmPassword", "")

	// 验证参数
	if username == "" || password == "" || confirmPassword == "" {
		ReturnError(c, 4001, "请输入正确的信息")
		return
	}
	if password != confirmPassword {
		ReturnError(c, 4001, "两次密码不一致")
		return
	}

	//检查用户是否已存在
	if user, _ := models.GetUserByUsername(username); user != nil {
		ReturnError(c, 4004, "用户已存在")
		return
	}

	// 创建用户
	user := &models.User{
		Username: username,
		Password: EncryMd5(password),
	}
	if err := models.CreateUser(user); err != nil {
		ReturnError(c, 4004, "注册失败")
		return
	}

	ReturnSuccess(c, 200, "注册成功", nil, 0)

}

func (uc UserController) Login(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")
	// 验证参数
	if username == "" || password == "" {
		ReturnError(c, 4001, "请输入正确的信息")
		return
	}

	// 检查用户是否存在
	user, err := models.GetUserByUsername(username)
	if err != nil {
		ReturnError(c, 4004, "用户不存在")
		return
	}
	// 检查密码是否正确
	if user.Password != EncryMd5(password) {
		ReturnError(c, 4004, "密码错误")
		return
	}

	session := sessions.Default(c)
	if session == nil {
		ReturnError(c, 5000, "session初始化失败")
		return
	}

	session.Set("login:"+strconv.Itoa(user.Id), user.Id)
	if err := session.Save(); err != nil {
		fmt.Println("session存储失败:", err)
		ReturnError(c, 5000, "session存储失败")
		return
	}

	ReturnSuccess(c, 200, "登录成功", user.Username, 0)

}
