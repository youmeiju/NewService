package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

	beego.GlobalControllerRouter["NewService/controllers:GoodsControllers"] = append(beego.GlobalControllerRouter["NewService/controllers:GoodsControllers"],
		beego.ControllerComments{
			Method: "GetFirstMenu",
			Router: `/GetFirstMenu`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["NewService/controllers:GoodsControllers"] = append(beego.GlobalControllerRouter["NewService/controllers:GoodsControllers"],
		beego.ControllerComments{
			Method: "GetGoodsDetail",
			Router: `/GetGoodsDetail`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["NewService/controllers:GoodsControllers"] = append(beego.GlobalControllerRouter["NewService/controllers:GoodsControllers"],
		beego.ControllerComments{
			Method: "GetGoodsList",
			Router: `/GetGoodsList`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["NewService/controllers:GoodsControllers"] = append(beego.GlobalControllerRouter["NewService/controllers:GoodsControllers"],
		beego.ControllerComments{
			Method: "GetIndexGoods",
			Router: `/indexGoods`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["NewService/controllers:OrderControllers"] = append(beego.GlobalControllerRouter["NewService/controllers:OrderControllers"],
		beego.ControllerComments{
			Method: "NewOrder",
			Router: `/NewOrder`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["NewService/controllers:UserControllers"] = append(beego.GlobalControllerRouter["NewService/controllers:UserControllers"],
		beego.ControllerComments{
			Method: "GetUserInfo",
			Router: `/GetUserToken`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["NewService/controllers:UserControllers"] = append(beego.GlobalControllerRouter["NewService/controllers:UserControllers"],
		beego.ControllerComments{
			Method: "IsUserToken",
			Router: `/IsUserToken`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

}
