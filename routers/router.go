// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"BlogApi/controllers"

	"github.com/astaxie/beego"
)

func init() {
	ns :=
		beego.NewNamespace("/v1",
			/*beego.NSCond(func(context *context.Context) bool {
				token := context.Request.Header.Get("token")
				beego.Info(fmt.Sprintf("token is  %v", token))
				if len(strings.Fields(token)) == 0 {
					return false
				}
				return true
			}),*/
			beego.NSNamespace("/user",
				beego.NSInclude(
					&controllers.UserController{},
				),
			),
		)
	beego.AddNamespace(ns)

	beego.Router("/api/list", &controllers.TestController{}, "*:List")
	beego.Router("/person/:last/:first", &controllers.TestController{})
	beego.AutoRouter(&controllers.TestController{})
}
