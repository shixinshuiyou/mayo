package controllers

import (
	"mayo/models"
	"encoding/json"
	"strings"

	"github.com/astaxie/beego"
)

// 用户个人信息
type MaYoController struct {
	beego.Controller
}

// @Title 注册用户
// @用户注册信息
// @Param	body		body 	models.Mayo	true		"body for user content"
// @Success 200 {int} models.Mayo.Id
// @Failure 403 body is empty
// @router / [post]
func (mayo *MaYoController) Post() {
	var user models.Mayo
	err := json.Unmarshal(mayo.Ctx.Input.RequestBody, &user)
	if err != nil {
		panic("注册信息有误，请检查")
	}
	uid := models.AddUser(user)
	mayo.Data["json"] = models.Success(uid)
	mayo.ServeJSON()
}

// @Title 用户登录
// @用戶使用账号密码登录
// @Param	name	query 	string	true	"用户名"
// @Param	pass	query	string	true	"密码"
// @Success 200 {object} models.Mayo
// @Failure 403 :uid is empty
// @router /name [get]
func (mayo *MaYoController) AdminByName() {
	name := mayo.GetString("name")
	// TODO PASSWORD 加解密
	pass := mayo.GetString("pass")
	if name == "" && pass == "" && strings.Contains(name, " ") && strings.Contains(pass, " ") {
		mayo.Data["json"] = models.CharFail("")
	}
	mayo.ServeJSON()
}

// @Title Delete
// @Description delete the user
// @Param	uid		path 	string	true		"The uid you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 uid is empty
// @router /:uid [delete]
func (mayo *MaYoController) Delete() {
	uid := mayo.GetString(":uid")
	models.DeleteUser(uid)
	mayo.Data["json"] = "delete success!"
	mayo.ServeJSON()
}

// @Title Login
// @Description Logs user into the system
// @Param	username		query 	string	true		"The username for login"
// @Param	password		query 	string	true		"The password for login"
// @Success 200 {string} login success
// @Failure 403 user not exist
// @router /login [get]
func (mayo *MaYoController) Login() {
	username := mayo.GetString("username")
	password := mayo.GetString("password")
	if models.Login(username, password) {
		mayo.Data["json"] = "login success"
	} else {
		mayo.Data["json"] = "user not exist"
	}
	mayo.ServeJSON()
}

// @Title logout
// @Description Logs out current logged in user session
// @Success 200 {string} logout success
// @router /logout [get]
func (mayo *MaYoController) Logout() {
	mayo.Data["json"] = "logout success"
	mayo.ServeJSON()
}
