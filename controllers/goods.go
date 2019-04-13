package controllers

import (
	"NewService/models"
	"encoding/json"

	"github.com/astaxie/beego"
	"fmt"
)

type GoodsControllers struct {
	beego.Controller
}

// @Title GetIndexInfo
// @Description GetIndexInfo
// @Param token query string true "token"
// @Param utoken query string true  "utoken"
// @Success 200 {map}
// @Failure 403 body is empty
// @router /indexGoods [get]
func (GC *GoodsControllers) GetIndexGoods() {
	token := GC.GetString("token")
	if IsToken(token) {
		utoken := GC.GetString("utoken")
		num := models.IsUserToken(token, utoken)
		if num == 0 {
			GC.Data["json"] = models.SetResponse("", 40000, "token不存在正在重新登录")
		} else {
			uid := models.GetUserId(utoken)[0]["id"].(string)
			indexInfo, err := models.GetIndexInfo(uid)
			fmt.Println(err)
			Msg, Code := models.RealMsg(err)
			GC.Data["json"] = models.SetResponse(indexInfo, Code, Msg)
		}
	} else {
		GC.Data["json"] = models.SetResponse("", -1, "非法的访问用户")
	}
	GC.ServeJSON()
}

// @Title GetGoodsList
// @Description SearchGoodsList
// @Param token query  string true "token"
// @Param utoken query string true "utoken"
// @Param searchData body map true "body for SearchGoodsList"
// @Success 200 {map}
// @Failure 403 body is empty
// @router /GetGoodsList [post]
func (GC *GoodsControllers) GetGoodsList() {
	token := GC.GetString("token")
	if IsToken(token) {
		utoken := GC.GetString("utoken")
		num := models.IsUserToken(token, utoken)
		if num == 0 {
			GC.Data["json"] = models.SetResponse("", 40000, "token不存在正在重新登录")
		} else {
			var dat map[string]string
			json.Unmarshal(GC.Ctx.Input.RequestBody, &dat)
			id := dat["id"]
			pagesize := dat["pagesize"]
			pagenum := dat["pagenum"]
			if pagenum == "" {
				pagenum = "10"
			}
			keywords := dat["keywords"]
			stype := dat["stype"]
			stype_jiage := dat["stype_jiage"]
			switch stype {
			case "":
				{
					stype = "id"
					stype_jiage = "ASC"
				}
			case "xiaoliang":
				{
					stype = "sale_valum"
					stype_jiage = "DESC"
				}
			case "jiage":
				{
					stype = "shop_price"
					if stype_jiage == "jiage_sheng" {
						stype_jiage = "ASC"
					} else {
						stype_jiage = "DESC"
					}
				}
			}
			if id == "" {
				content, err := models.GetGoodsListByKeyWord(pagesize, pagenum, keywords, stype, stype_jiage)
				Msg, Code := models.RealMsg(err)
				GC.Data["json"] = models.SetResponse(content, Code, Msg)
			} else {
				content, err := models.GetGoodsListByKeyId(pagesize, pagenum, id, stype, stype_jiage)
				Msg, Code := models.RealMsg(err)
				GC.Data["json"] = models.SetResponse(content, Code, Msg)
			}
		}
	} else {
		GC.Data["json"] = models.SetResponse("", -1, "非法的访问用户")
	}
	GC.ServeJSON()
}

// @Title FirstMenu
// @Description SearchGoodsList
// @Param token query string true "token"
// @Param utoken query string true "utoken"
// @Success 200 {map}
// @Failure 403 body is empty
// @router /GetFirstMenu [get]
func (GC *GoodsControllers) GetFirstMenu() {
	token := GC.GetString("token")
	if IsToken(token) {
		utoken := GC.GetString("utoken")
		num := models.IsUserToken(token, utoken)
		if num == 0 {
			GC.Data["json"] = models.SetResponse("", 40000, "token不存在正在重新登录")
		} else {
			content, err := models.GetFirstMenu()
			Msg, Code := models.RealMsg(err)
			GC.Data["json"] = models.SetResponse(content, Code, Msg)
		}
	} else {
		GC.Data["json"] = models.SetResponse("", -1, "非法的访问用户")
	}
	GC.ServeJSON()
}

// @Title GetGoodsDetail
// @Description GetGoodsDetail
// @Param	id	query 	int	 true	"id for GoodsDetail"
// @Param token query string  "token"
// @Param utoken query string  "utoken"
// @Success 200 {map}
// @Failure 403 body is empty
// @router /GetGoodsDetail [get]
func (GC *GoodsControllers) GetGoodsDetail() {
	token := GC.GetString("token")
	if IsToken(token) {
		utoken := GC.GetString("utoken")
		num := models.IsUserToken(token, utoken)
		if num == 0 {
			GC.Data["json"] = models.SetResponse("", 40000, "token不存在正在重新登录")
		} else {
			id := GC.GetString("id")
			content, err := models.GetGoodsDetail(id)
			Msg, Code := models.RealMsg(err)
			GC.Data["json"] = models.SetResponse(content, Code, Msg)
		}
	} else {
		GC.Data["json"] = models.SetResponse("", -1, "非法的访问用户")
	}
	GC.ServeJSON()
}
