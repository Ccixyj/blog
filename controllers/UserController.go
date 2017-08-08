package controllers

import "github.com/astaxie/beego"

// User API
type UserController struct {
	beego.Controller
}

//提升性能
func (ctx *UserController) URLMapping() {
	ctx.Mapping("UserData", ctx.UserData)
}

// @Title Get UserData
// @Description Get Product list by some info
// @Param   token	header	string	true	"The token"
// @Success 200 { object } models.User.User
// @router /all/:key [get]
func (ctx *UserController) UserData() {
	beego.Debug("Params", ctx.Ctx.Input.Params())
	ctx.Data["json"] = map[string]string{"one": "abc", "two": "qwer"}
	ctx.ServeJSON()
}
