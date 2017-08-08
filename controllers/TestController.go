package controllers

import "github.com/astaxie/beego"

type TestController struct {
	beego.Controller
}

func (this *TestController) Get() {
	beego.Debug("Get")
	this.Data["Username"] = "astaxie"
	this.Ctx.Output.Body([]byte("ok"))
}

func (this *TestController) List() {
	beego.Debug("List")
	this.Ctx.Output.Body([]byte("i am list"))
}

func (this *TestController) Params() {
	beego.Debug("Params", this.Ctx.Input.Params())
	this.Ctx.Output.Body([]byte(this.Ctx.Input.Params()["0"] + this.Ctx.Input.Params()["1"] + this.Ctx.Input.Params()["2"]))
}

func (this *TestController) Myext() {
	beego.Debug("Myext", this.Ctx.Input.Params())
	this.Ctx.Output.Body([]byte(this.Ctx.Input.Param(":ext")))
}

func (this *TestController) GetUrl() {
	beego.Debug("GetUrl")
	this.Ctx.Output.Body([]byte(this.URLFor(".Myext")))
}
