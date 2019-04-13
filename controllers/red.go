package controllers

import (
	"NewService/models"

	"github.com/astaxie/beego"
)

type RedController struct {
	beego.Controller
}

// @Title GetRedDetail
// @Description GetRedDetail
// @Param token query  string true "token"
// @Param rid query string true "rid"
// @Success 200 {map}
// @Failure 403 body is empty
// @router /GetRedDetail [get]
func (RC *RedController) GetRedDetail() {
	token := RC.GetString("token")
	if IsToken(token) {
		rid := RC.GetString("rid")
		err, redDetail := models.GetRedDetail(rid)
		if err != nil {
			panic(err)
		} else {
			RC.Data["json"] = models.SetResponse(redDetail[0], 1, "")
		}
	}
	RC.ServeJSON()
}

// @Title GetRedDetail1
// @Description GetRedDetail1
// @Param token query  string true "token"
// @Param rid query string true "rid"
// @Success 200 {map}
// @Failure 403 body is empty
// @router /GetRedDetail1 [get]
func (RC *RedController) GetRedDetail1() {
	token := RC.GetString("token")
	if IsToken(token) {
		rid := RC.GetString("rid")
		err, redDetail := models.GetRedDetail1(rid)
		if err != nil {
			panic(err)
		} else {
			RC.Data["json"] = models.SetResponse(redDetail[0], 1, "")
		}
	}
	RC.ServeJSON()
}


// @Title AddRed
// @Description AddRed
// @Param token query  string true "token"
// @Param utoken query  string true "utoken"
// @Param rid query string true "rid"
// @Success 200 {map}
// @Failure 403 body is empty
// @router /AddRed [get]
func (RC *RedController) AddRed() {
	token := RC.GetString("token")
	if IsToken(token) {
		utoken := RC.GetString("utoken")
		num := models.IsUserToken(token, utoken)
		if num == 0 {
			RC.Data["json"] = models.SetResponse("", 2, "token不存在正在重新登录")
		} else {
			rid := RC.GetString("rid")
			uid := models.GetUserId(utoken)[0]["id"].(string)
			code := models.AddRed(uid, rid)
			if code == 1 {
				RC.Data["json"] = models.SetResponse("", 1, "红包领取成功")
			} else if code == 4 {
				RC.Data["json"] = models.SetResponse("", 4, "已领取过该红包请勿重新领取")
			} else {
				RC.Data["json"] = models.SetResponse("", 5, "服务器飞往火星，请稍后重试")
			}
		}
	}
	RC.ServeJSON()
}

// @Title GetRedList
// @Description GetRedList
// @Param token query  string true "token"
// @Param utoken query  string true "utoken"
// @Success 200 {map}
// @Failure 403 body is empty
// @router /GetRedList [get]
func (RC *RedController) GetRedList() {
	token := RC.GetString("token")
	if IsToken(token) {
		utoken := RC.GetString("utoken")
		num := models.IsUserToken(token, utoken)
		if num == 0 {
			RC.Data["json"] = models.SetResponse("", 2, "token不存在正在重新登录")
		} else {
			uid := models.GetUserId(utoken)[0]["id"].(string)
			err, redList := models.GetRedList(uid)
			if err != nil {
				panic(err)
			} else {
				RC.Data["json"] = models.SetResponse(redList, 1, "")
			}
		}
	}
	RC.ServeJSON()
}

// @Title GetUseRed
// @Description GetUseRed
// @Param token query  string true "token"
// @Param utoken query  string true "utoken"
// @Param allPrice query  string true "allPrice"
// @Success 200 {map}
// @Failure 403 body is empty
// @router /GetUseRed [get]
func (RC *RedController) GetUseRed() {
	token := RC.GetString("token")
	if IsToken(token) {
		utoken := RC.GetString("utoken")
		num := models.IsUserToken(token, utoken)
		if num == 0 {
			RC.Data["json"] = models.SetResponse("", 2, "token不存在正在重新登录")
		} else {
			uid := models.GetUserId(utoken)[0]["id"].(string)
			allPrice := RC.GetString("allPrice")
			num, red := models.GetUseRed(uid, allPrice)
			if num == 0 {
				RC.Data["json"] = models.SetResponse("", 4, "")
			} else {
				RC.Data["json"] = models.SetResponse(red[0], 1, "")
			}
		}
	}
	RC.ServeJSON()
}

// @Title GetUseRedList
// @Description GetUseRedList
// @Param token query  string true "token"
// @Param utoken query  string true "utoken"
// @Param allprice query  string true "allprice"
// @Success 200 {map}
// @Failure 403 body is empty
// @router /GetUseRedList [get]
func (RC *RedController) GetUseRedList() {
	token := RC.GetString("token")
	if IsToken(token) {
		utoken := RC.GetString("utoken")
		num := models.IsUserToken(token, utoken)
		if num == 0 {
			RC.Data["json"] = models.SetResponse("", 2, "token不存在正在重新登录")
		} else {
			allprice := RC.GetString("allprice")
			uid := models.GetUserId(utoken)[0]["id"].(string)
			err, redList := models.GetUseRedList(uid, allprice)
			if err != nil {
				panic(err)
			} else {
				RC.Data["json"] = models.SetResponse(redList, 1, "")
			}
		}
	}
	RC.ServeJSON()
}
