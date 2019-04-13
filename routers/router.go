// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"NewService/controllers"

	"github.com/astaxie/beego"
)

func init() {
	ns := beego.NewNamespace("/v1",
		beego.NSNamespace("/goods",
			beego.NSInclude(
				&controllers.GoodsControllers{},
			),
		),
		beego.NSNamespace("/user",
			beego.NSInclude(
				&controllers.UserControllers{},
			),
		),
		beego.NSNamespace("/order",
			beego.NSInclude(
				&controllers.OrderControllers{},
			),
		),
		beego.NSNamespace("/red",
			beego.NSInclude(
				&controllers.RedController{},
			),
		),
		beego.NSNamespace("/adminUser",
			beego.NSInclude(
				&controllers.AdminUserControllers{},
			),
		),
	)
	beego.AddNamespace(ns)
}
