package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

	beego.GlobalControllerRouter["BlogApi/controllers:UserController"] = append(beego.GlobalControllerRouter["BlogApi/controllers:UserController"],
		beego.ControllerComments{
			Method: "UserData",
			Router: `/all/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

}
