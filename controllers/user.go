package controllers

import (
	"github.com/astaxie/beego"
	"fmt"
	"net/http"
	"io/ioutil"
	"encoding/json"
	"NewService/models"
)

type UserControllers struct {
	beego.Controller
}
 //@Title GetUserToken
 //@Description GetGoodsDetail
 //@Param data body map true	"data for GetUserInfo"
 //@Success 200 {map}
 //@Failure 403 body is empty
 //@router /GetUserToken [post]
func (UC *UserControllers)GetUserInfo()  {
	var dat map[string]string
	json.Unmarshal(UC.Ctx.Input.RequestBody, &dat)
	token := dat["token"]
	if IsToken(token){
		code :=  dat["code"]
		//-1请求微信接口code换取openid和SESSIONKEY
		url:= fmt.Sprintf("https://api.weixin.qq.com/sns/jscode2session?appid=wx331b51ae3849285d&secret=9aff02f9a7e6db67e4c429a0d2009a67&js_code=%s&grant_type=authorization_code",code)
		res, err := http.Get(url)
		if err == nil{
			defer res.Body.Close()
			m := make(map[string]string)
			body, _ := ioutil.ReadAll(res.Body)
			err:=json.Unmarshal(body,&m)
			if err == nil{
				if m["errcode"] == ""{
					openid := m["openid"]
					//openid和SESSIONKEY换取userInfo
					num,user:=models.GetUserInfo(openid)
					//新增用户并返回userInfo
					if num == 0{
						err:=models.AddUser(openid)
						if err == nil{
							_,user:=models.GetUserInfo(openid)
							id := user[0]["id"].(string)
							token := models.SetUserToken(id)
							err:=models.UpdateUserToken(id,token)
							m := make(map[string]string)
							m["id"] = id
							m["token"] = token
							Msg,Code := models.RealMsg(err)
							UC.Data["json"] = models.SetResponse(m,Code,Msg)
						}else{
							UC.Data["json"] = models.SetResponse("",40001,"微信授权失败,请稍后重试4")
						}
					}else{
						UC.Data["json"] = models.SetResponse(user,200,"")
					}
				}else{
					UC.Data["json"] = models.SetResponse("",40001,"微信授权失败,请稍后重试3")
				}
			}else{
				UC.Data["json"] = models.SetResponse("",40001,"微信授权失败,请稍后重试2")
			}
		}else{
			UC.Data["json"] = models.SetResponse("",41001,"微信授权失败,请稍后重试1")
		}
	}else{
		UC.Data["json"] = models.SetResponse("",0,"非法的访问")
	}
	UC.ServeJSON()
}
//@Title IsUserToken
//@Description IsUserToken
//@Param Token body map true "body for Token"
//@Success 200 {map}
//@Failure 403 body is empty
//@router /IsUserToken [post]
func (UC *UserControllers)IsUserToken()  {
	var dat map[string]string
	json.Unmarshal(UC.Ctx.Input.RequestBody, &dat)
	token := dat["token"]
	utoken := dat["utoken"]
	if IsToken(token){
		num:=models.IsUserToken(token,utoken)
		if num == 0 {
			UC.Data["json"] = models.SetResponse("",0,"用户token不存在")
		}else {
			UC.Data["json"] = models.SetResponse("",1,"用户token存在")
		}
	}else {
		UC.Data["json"] = models.SetResponse("",-1,"非法的访问用户")
	}
	UC.ServeJSON()
}

//token是否正确
func IsToken(token string)(bool)  {
	if token == "wujiu59" {
		return true
	}
	return false
}

//@Title SetInviteId
//@Description SetInviteId
// @Param iid query  string true "token"
// @Param uid query  string true "id"
//@router /SetInviteId [get]
func (UC *UserControllers)SetInviteId()  {
	uid := UC.GetString("uid")
	iid := models.GetInviteId(uid)
	if iid[0]["invite_id"] == nil{
		models.InsertInviteId(uid,UC.GetString("iid"))
		UC.Data["json"] = models.SetResponse("",1,"")
	}else{
		UC.Data["json"] = models.SetResponse("",1,"")
	}
	UC.ServeJSON()
}
