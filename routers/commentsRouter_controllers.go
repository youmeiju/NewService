package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

	beego.GlobalControllerRouter["NewService/controllers:AdminUserControllers"] = append(beego.GlobalControllerRouter["NewService/controllers:AdminUserControllers"],
		beego.ControllerComments{
			Method: "DeleteUser",
			Router: `/DeleteUser`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["NewService/controllers:AdminUserControllers"] = append(beego.GlobalControllerRouter["NewService/controllers:AdminUserControllers"],
		beego.ControllerComments{
			Method: "GetAdminOrderDetail",
			Router: `/GetAdminOrderDetail`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["NewService/controllers:AdminUserControllers"] = append(beego.GlobalControllerRouter["NewService/controllers:AdminUserControllers"],
		beego.ControllerComments{
			Method: "GetAdminOrderList",
			Router: `/GetAdminOrderList`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["NewService/controllers:AdminUserControllers"] = append(beego.GlobalControllerRouter["NewService/controllers:AdminUserControllers"],
		beego.ControllerComments{
			Method: "GetAdminUserInfo",
			Router: `/GetAdminUserInfo`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["NewService/controllers:AdminUserControllers"] = append(beego.GlobalControllerRouter["NewService/controllers:AdminUserControllers"],
		beego.ControllerComments{
			Method: "GetOrderNum",
			Router: `/GetOrderNum`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["NewService/controllers:AdminUserControllers"] = append(beego.GlobalControllerRouter["NewService/controllers:AdminUserControllers"],
		beego.ControllerComments{
			Method: "BindUser",
			Router: `/GetUser`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["NewService/controllers:AdminUserControllers"] = append(beego.GlobalControllerRouter["NewService/controllers:AdminUserControllers"],
		beego.ControllerComments{
			Method: "Login",
			Router: `/Login`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["NewService/controllers:AdminUserControllers"] = append(beego.GlobalControllerRouter["NewService/controllers:AdminUserControllers"],
		beego.ControllerComments{
			Method: "NewUser",
			Router: `/NewUser`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["NewService/controllers:AdminUserControllers"] = append(beego.GlobalControllerRouter["NewService/controllers:AdminUserControllers"],
		beego.ControllerComments{
			Method: "Refund",
			Router: `/Refund`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["NewService/controllers:AdminUserControllers"] = append(beego.GlobalControllerRouter["NewService/controllers:AdminUserControllers"],
		beego.ControllerComments{
			Method: "RefundNotice",
			Router: `/RefundNotice`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["NewService/controllers:AdminUserControllers"] = append(beego.GlobalControllerRouter["NewService/controllers:AdminUserControllers"],
		beego.ControllerComments{
			Method: "UpdateOrderUser",
			Router: `/UpdateOrderUser`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["NewService/controllers:AdminUserControllers"] = append(beego.GlobalControllerRouter["NewService/controllers:AdminUserControllers"],
		beego.ControllerComments{
			Method: "UpdateUser",
			Router: `/UpdateUser`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["NewService/controllers:AdminUserControllers"] = append(beego.GlobalControllerRouter["NewService/controllers:AdminUserControllers"],
		beego.ControllerComments{
			Method: "WXBind",
			Router: `/WXBind`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

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
			Method: "GetOrderDetail",
			Router: `/GetOrderDetail`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["NewService/controllers:OrderControllers"] = append(beego.GlobalControllerRouter["NewService/controllers:OrderControllers"],
		beego.ControllerComments{
			Method: "GetOrderInfo",
			Router: `/GetOrderInfo`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["NewService/controllers:OrderControllers"] = append(beego.GlobalControllerRouter["NewService/controllers:OrderControllers"],
		beego.ControllerComments{
			Method: "GetOrderList",
			Router: `/GetOrderList`,
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

	beego.GlobalControllerRouter["NewService/controllers:OrderControllers"] = append(beego.GlobalControllerRouter["NewService/controllers:OrderControllers"],
		beego.ControllerComments{
			Method: "OrderPay",
			Router: `/OrderPay`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["NewService/controllers:OrderControllers"] = append(beego.GlobalControllerRouter["NewService/controllers:OrderControllers"],
		beego.ControllerComments{
			Method: "UpdateOrderStatue",
			Router: `/UpdateOrderStatue`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["NewService/controllers:OrderControllers"] = append(beego.GlobalControllerRouter["NewService/controllers:OrderControllers"],
		beego.ControllerComments{
			Method: "UpdateOrderStatues",
			Router: `/UpdateOrderStatues`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["NewService/controllers:RedController"] = append(beego.GlobalControllerRouter["NewService/controllers:RedController"],
		beego.ControllerComments{
			Method: "AddRed",
			Router: `/AddRed`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["NewService/controllers:RedController"] = append(beego.GlobalControllerRouter["NewService/controllers:RedController"],
		beego.ControllerComments{
			Method: "GetRedDetail",
			Router: `/GetRedDetail`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["NewService/controllers:RedController"] = append(beego.GlobalControllerRouter["NewService/controllers:RedController"],
		beego.ControllerComments{
			Method: "GetRedDetail1",
			Router: `/GetRedDetail1`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["NewService/controllers:RedController"] = append(beego.GlobalControllerRouter["NewService/controllers:RedController"],
		beego.ControllerComments{
			Method: "GetRedList",
			Router: `/GetRedList`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["NewService/controllers:RedController"] = append(beego.GlobalControllerRouter["NewService/controllers:RedController"],
		beego.ControllerComments{
			Method: "GetUseRed",
			Router: `/GetUseRed`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["NewService/controllers:RedController"] = append(beego.GlobalControllerRouter["NewService/controllers:RedController"],
		beego.ControllerComments{
			Method: "GetUseRedList",
			Router: `/GetUseRedList`,
			AllowHTTPMethods: []string{"get"},
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

	beego.GlobalControllerRouter["NewService/controllers:UserControllers"] = append(beego.GlobalControllerRouter["NewService/controllers:UserControllers"],
		beego.ControllerComments{
			Method: "SetInviteId",
			Router: `/SetInviteId`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

}
