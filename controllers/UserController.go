package controllers

import (
	"BlogApi/models"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

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
	defer ctx.ServeJSON()
	var users []*models.User
	num, err := orm.NewOrm().QueryTable("user").Limit(-1).All(&users)
	if err != nil {
		beego.Error("查询所有用户错误")
		return
	}
	beego.Debug(fmt.Sprintf("查询成功; 共%d条数据 ", num))
	ctx.Data["json"] = models.NewSuccess("查询成功", users)

}
