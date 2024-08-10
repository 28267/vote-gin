package controllers

import (
	"strconv"

	"vote-gin/models"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

type UserController struct{}

type UserApi struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
}

func (u UserController) Register(c *gin.Context) {
	//接收用户名 密码 确认密码
	username := c.DefaultPostForm("username", "")
	password := c.DefaultPostForm("password", "")
	confirmPassword := c.DefaultPostForm("confirmPassword", "")

	if username == "" || password == "" || confirmPassword == "" {
		ReturnError(c, 4001, "请输入正确的信息")
		return
	}

	if password != confirmPassword {
		ReturnError(c, 4001, "密码和确认密码不一致")
		return
	}

	user, err := models.GetUserInfoUsername(username)
	if user.Id != 0 {
		ReturnError(c, 4001, "用户名已存在")
	}

	_, err = models.AddUser(username, EncryMd5(password))
	if err != nil {
		ReturnError(c, 4001, "保存失败，请联系管理员")
		return
	}
	ReturnSuccess(c, 0, "success", user, 1)
}

func (u UserController) Login(c *gin.Context) {
	//接收用户名 密码
	username := c.DefaultPostForm("username", "")
	password := c.DefaultPostForm("password", "")

	if username == "" || password == "" {
		ReturnError(c, 4004, "请输入正确的信息")
		return
	}

	user, _ := models.GetUserInfoUsername(username)
	if user.Id == 0 {
		ReturnError(c, 4004, "用户名或密码不正确")
		return
	}
	if user.Password != EncryMd5(password) {
		ReturnError(c, 4004, "用户名或密码不正确")
		return
	}

	session := sessions.Default(c)
	session.Set("login:"+strconv.Itoa(user.Id), user.Id)
	session.Save()
	data := UserApi{Id: user.Id, Username: user.Username}
	ReturnSuccess(c, 0, "登陆成功", data, 1)
}

// func (u UserController) AddUser(c *gin.Context) {
// 	username := c.DefaultPostForm("username", "")
// 	id, err := models.AddUser(username)
// 	if err != nil {
// 		ReturnError(c, 4002, "save error")
// 		return
// 	}
// 	ReturnSuccess(c, 0, "save success", id, 1)

// }

// func (u UserController) UpdateUser(c *gin.Context) {
// 	username := c.DefaultPostForm("username", "")
// 	idStr := c.DefaultPostForm("id", "")
// 	id, _ := strconv.Atoi(idStr)
// 	models.UpdateUser(id, username)
// 	ReturnSuccess(c, 0, "update success", true, 1)
// }

// func (u UserController) DeleteUser(c *gin.Context) {
// 	idStr := c.DefaultPostForm("id", "")
// 	id, _ := strconv.Atoi(idStr)
// 	err := models.DeleteUser(id)
// 	if err != nil {
// 		ReturnError(c, 4002, "delete error")
// 		return
// 	}
// 	ReturnSuccess(c, 0, "delete success", id, 1)
// }

// func (u UserController) GetUser(c *gin.Context) {
// 	idStr := c.DefaultPostForm("id", "")
// 	id, _ := strconv.Atoi(idStr)
// 	user, err := models.GetUser(id)
// 	if err != nil {
// 		ReturnError(c, 4002, "none info List")
// 		return
// 	}
// 	ReturnSuccess(c, 0, "success", user, 1)
// }
