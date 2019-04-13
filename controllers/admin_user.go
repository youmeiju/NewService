package controllers

import (
	"NewService/models"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"math/rand"
	"github.com/astaxie/beego"
	"time"
	"crypto/tls"
	"strings"
	"strconv"
	"encoding/xml"
	"encoding/base64"
	"github.com/nanjishidu/gomini/gocrypto"
)

type AdminUserControllers struct {
	beego.Controller
}

//@Title Login
//@Description Login
//@Param data body map true	"data for Login"
//@Success 200 {map}
//@Failure 403 body is empty
//@router /Login [post]
func (AUC *AdminUserControllers) Login() {
	var dat map[string]string
	json.Unmarshal(AUC.Ctx.Input.RequestBody, &dat)
	user, password := dat["user"], dat["password"]
	num, userInfo := models.ConfirmUser(user, password)
	if num == 0 {
		AUC.Data["json"] = models.SetResponse("", 2, "用户名或密码错误请重新登录")
	} else {
		realUserInfo:= make(map[string]string)
		id := userInfo[0]["id"].(string)
		token := models.SetUserToken(id)
		err := models.UpAdminUser(id,token)
		if err == nil {
			realUserInfo["id"] = id
			realUserInfo["is_surper"] = userInfo[0]["is_surper"].(string)
			realUserInfo["real_name"] = userInfo[0]["real_name"].(string)
			realUserInfo["token"] = token
			AUC.Data["json"] = models.SetResponse(realUserInfo, 1, "")
		} else {
			AUC.Data["json"] = models.SetResponse("", 5, "服务器飞往火星")
		}
	}
	AUC.ServeJSON()
}

//@Title GetOrderNum
//@Description GetOrderNum
// @Param token query  string true "token"
// @Param id query  string true "id"
// @Param isSuper query  string true "isSuper"
//@Success 200 {map}
//@Failure 403 body is empty
//@router /GetOrderNum [get]
func (AUC *AdminUserControllers) GetOrderNum() {
	id := AUC.GetString("id")
	token := AUC.GetString("token")
	num := models.IsAdminUser(id, token)
	if num == 0 {
		AUC.Data["json"] = models.SetResponse("", 2, "登录失效，请重新登录后重试")
	} else {
		isSuper := AUC.GetString("isSuper")
		if isSuper == "1" {
			num := models.GetOrderNumByState()
			AUC.Data["json"] = models.SetResponse(num, 1, "")
		} else {
			num := models.GetOrderNumById(id)
			AUC.Data["json"] = models.SetResponse(num, 1, "")
		}
	}
	AUC.ServeJSON()
}

//@Title GetAdminOrderList
//@Description GetAdminOrderList
// @Param token query  string true "token"
// @Param id query  string true "id"
// @Param isSuper query  string true "isSuper"
//@Success 200 {map}
//@Failure 403 body is empty
//@router /GetAdminOrderList [get]
func (AUC *AdminUserControllers) GetAdminOrderList() {
	id := AUC.GetString("id")
	token := AUC.GetString("token")
	num := models.IsAdminUser(id, token)
	if num == 0 {
		AUC.Data["json"] = models.SetResponse("", 2, "登录失效，请重新登录后重试")
	} else {
		isSuper := AUC.GetString("isSuper")
		if isSuper == "1" {
			_, orders := models.GetAdminOrderList()
			for i, v := range orders {
				_, orderGoodsInfo := models.GetOrderGoodsInfo(v["id"].(string))
				realGoodsMsg := make([]map[string]interface{}, len(orderGoodsInfo))
				for i, v := range orderGoodsInfo {
					a := make(map[string]interface{})
					a["buy_count"] = v["buy_count"]
					a["shop_price"] = v["shop_price"]
					goodsInfo, _ := models.GetGoodsDetail(v["g_id"].(string))
					a["goodsInfo"] = goodsInfo[0]
					realGoodsMsg[i] = a
				}
				orders[i]["glist"] = realGoodsMsg
			}
			AUC.Data["json"] = models.SetResponse(orders, 1, "")
		} else {
			_, orders := models.GetNotAdminOrderList(id)
			for i, v := range orders {
				_, orderGoodsInfo := models.GetOrderGoodsInfo(v["id"].(string))
				a := make(map[string]interface{})
				realGoodsMsg := make([]map[string]interface{}, len(orderGoodsInfo))
				for i, v := range orderGoodsInfo {
					a["buy_count"] = v["buy_count"]
					a["shop_price"] = v["shop_price"]
					goodsInfo, _ := models.GetGoodsDetail(v["g_id"].(string))
					a["goodsInfo"] = goodsInfo[0]
					realGoodsMsg[i] = a
				}
				orders[i]["glist"] = realGoodsMsg
			}
			AUC.Data["json"] = models.SetResponse(orders, 1, "")
		}
	}
	AUC.ServeJSON()
}

//@Title GetAdminOrderDetail
//@Description GetAdminOrderList
// @Param token query  string true "token"
// @Param id query  string true "id"
// @Param oid query  string true "oid"
//@Success 200 {map}
//@Failure 403 body is empty
//@router /GetAdminOrderDetail [get]
func (AUC *AdminUserControllers) GetAdminOrderDetail() {
	id := AUC.GetString("id")
	token := AUC.GetString("token")
	num := models.IsAdminUser(id, token)
	if num == 0 {
		AUC.Data["json"] = models.SetResponse("", 2, "登录失效，请重新登录后重试")
	} else {
		oid := AUC.GetString("oid")
		err, orderInfo := models.GetOrderInfo(oid)
		if err == nil {
			_, orderGoodsInfo := models.GetOrderGoodsInfo(oid)
			a := make(map[string]interface{})
			realGoodsMsg := make([]map[string]interface{}, len(orderGoodsInfo))
			for i, v := range orderGoodsInfo {
				a["buy_count"] = v["buy_count"]
				a["shop_price"] = v["shop_price"]
				goodsInfo, _ := models.GetGoodsDetail(v["g_id"].(string))
				a["goodsInfo"] = goodsInfo[0]
				realGoodsMsg[i] = a
			}
			data := make(map[string]interface{})
			data["orderInfo"] = orderInfo
			data["lGoodsInfo"] = realGoodsMsg
			AUC.Data["json"] = models.SetResponse(data, 1, "")
		} else {
			panic(err)
		}
	}
	AUC.ServeJSON()
}

//@Title GetUser
//@Description GetUser
// @Param token query  string true "token"
// @Param id query  string true "id"
// @Param code query  string true "code"
//@Success 200 {map}
//@Failure 403 body is empty
//@router /GetUser [get]
func (AUC *AdminUserControllers)BindUser()  {
	id := AUC.GetString("id")
	token := AUC.GetString("token")
	num := models.IsAdminUser(id, token)
	if num == 0 {
		AUC.Data["json"] = models.SetResponse("", 2, "登录失效，请重新登录后重试")
	} else {
		if code:=AUC.GetString("code");code=="1"{
			userList:=models.GetUser("0","1")
			AUC.Data["json"] = models.SetResponse(userList, 1, "")
		}else{
			userList:=models.GetUser("0","")
			AUC.Data["json"] = models.SetResponse(userList, 1, "")
		}
	}
	AUC.ServeJSON()
}

//@Title UpdateOrderUser
//@Description UpdateOrderUser
// @Param token query  string true "token"
// @Param id query  string true "id"
// @Param oid query  string true "oid"
// @Param aid query  string true "aid"
//@Success 200 {map}
//@Failure 403 body is empty
//@router /UpdateOrderUser [get]
func (AUC *AdminUserControllers)UpdateOrderUser()  {
	id := AUC.GetString("id")
	token := AUC.GetString("token")
	num := models.IsAdminUser(id, token)
	if num == 0 {
		AUC.Data["json"] = models.SetResponse("", 2, "登录失效，请重新登录后重试")
	} else {
		oid:=AUC.GetString("oid")
		aid:=AUC.GetString("aid")
		err:=models.UpdateOrderUser(oid,aid)
		if err==nil {
			PushOpenid,_:=models.GetUserPushOpenid(aid)
			pushMsg(oid,PushOpenid[0]["push_openid"])
			AUC.Data["json"] = models.SetResponse("", 1, "")
		}else{
			AUC.Data["json"] = models.SetResponse("", 5, "服务器飞往火星请稍后重试")
		}

	}
	AUC.ServeJSON()
}

//@Title Refund
//@Description Refund
// @Param token query  string true "token"
// @Param id query  string true "id"
// @Param oid query  string true "oid"
// @Param price query  string true "price"
//@Success 200 {map}
//@Failure 403 body is empty
//@router /Refund [get]
func (AUC *AdminUserControllers)Refund() {
	id := AUC.GetString("id")
	token := AUC.GetString("token")
	num := models.IsAdminUser(id, token)
	if num == 0 {
		AUC.Data["json"] = models.SetResponse("", 2, "登录失效，请重新登录后重试")
	} else {
		oid := AUC.GetString("oid")
		price,err := AUC.GetInt("price")
	var reqMap= make(map[string]string, 0)
	reqMap["appid"] = "wx331b51ae3849285d" //微信小程序appid //商品描述
	reqMap["mch_id"] = "1513958441"        //商户号
	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))
	reqMap["nonce_str"] = time.Now().Format("20060102150405") + fmt.Sprintf("%04v", rnd.Int31n(1000)) //随机数
	reqMap["notify_url"] = "https://service.shanghaiyoumeiju2018.com/v1/adminUser/RefundNotice"                                       //通知地址
	reqMap["out_trade_no"] = oid                                                                     //商户订单号
	reqMap["out_refund_no"] = oid                                                             //商户订单号
	reqMap["total_fee"] = strconv.Itoa(price*100)                                                             //订单总金额，单位为分
	reqMap["refund_fee"] = strconv.Itoa(price*100)                                                                //退款总金额，单位为分
	reqMap["sign"] = WxPayCalcSign(reqMap, "zw620521198608071116000000000000")
	reqStr := Map2Xml(reqMap)
	var (
		//wechatCertPath= `D:\soft\beego\src\NewService\test\cert/apiclient_cert.pem`
		//wechatKeyPath= `D:\soft\beego\src\NewService\test\cert/apiclient_key.pem`
		wechatCertPath= `/root/go/src/NewService/cert/apiclient_cert.pem`
		wechatKeyPath= `/root/go/src/NewService/cert/apiclient_key.pem`
	)
	cert, _ := tls.LoadX509KeyPair(wechatCertPath, wechatKeyPath)
	var _tlsConfig *tls.Config
	_tlsConfig = &tls.Config{
		Certificates: []tls.Certificate{cert},
	}
	tr := &http.Transport{TLSClientConfig: _tlsConfig}
	client := &http.Client{Transport: tr}
	resp, err := client.Post("https://api.mch.weixin.qq.com/secapi/pay/refund", "application/xml", strings.NewReader(reqStr))
	if err!=nil {
		panic(err)
	}
	//defer resp.Body.Close()
	defer resp.Body.Close()
	bodyByte, _ := ioutil.ReadAll(resp.Body)
	resp1 := models.RefundOrderResp{}
	err = xml.Unmarshal(bodyByte, &resp1)
	if resp1.Return_code=="SUCCESS"{
		if resp1.Result_code == "SUCCESS"{
			err:=models.UpdateOrderStatue(oid,"004")
			if err==nil{
				AUC.Data["json"] = models.SetResponse("", 1, "操作退款成功，正在退款中")
			}else{
				AUC.Data["json"] = models.SetResponse("", 5, "服务器飞往火星稍后重试")
			}
		}else{
			AUC.Data["json"] = models.SetResponse("", 4, resp1.Err_code_des)
		}
	}else{
		AUC.Data["json"] = models.SetResponse("", 5, "服务器飞往火星稍后重试")
	}}
	AUC.ServeJSON()
}

//@Title RefundNotice
//@Description RefundNotice
//@Success 200 {map}
//@Failure 403 body is empty
//@router /RefundNotice [post]
func(AUC *AdminUserControllers)RefundNotice(){
	req := models.WXPayNotifyReq{}
	err := xml.Unmarshal(AUC.Ctx.Input.RequestBody, &req)
	if err != nil{
		panic(err)
	}
	//{wx331b51ae3849285d ABC_DEBIT 100 CNY N 1513958441 201809131441270059 o6GKW5Cpv7HhTbw2g9nPZ4IK3N90 80 SUCCESS SUCCESS 283C2E372DDA3B52A32C26772BC81E43 20180913144152 100 JSAPI 4200000202201809131526570757}
	if req.Return_code == "SUCCESS"{
		req2 := models.WXPayNotifyReq{}
		var (
			paykey   = "zw620521198608071116000000000000"
			req_info = req.Req_info
		)
		plaintext := respInfoDecode(paykey,req_info)
		xml.Unmarshal(plaintext, &req2)
		oid:=req2.Out_trade_no
		err:=models.UpdateOrderStatue(oid,"005")
		if err == nil{
			models.ReBackRed(oid)
			AUC.Data["xml"] = WXPayNotifyResp{"SUCCESS",""}
		}
	}
	AUC.ServeXML()
}
//wx req_info解码
func respInfoDecode(paykey string,req_info string)(plaintext []byte)  {
	b, _ := base64.StdEncoding.DecodeString(req_info)
	gocrypto.SetAesKey(strings.ToLower(gocrypto.Md5(paykey)))
	plaintext, _ = gocrypto.AesECBDecrypt(b)
	return
}

//@Title DeleteUser
//@Description DeleteUser
// @Param token query  string true "token"
// @Param id query  string true "id"
// @Param did query  string true "id"
//@Success 200 {map}
//@Failure 403 body is empty
//@router /DeleteUser [get]
func (AUC *AdminUserControllers)DeleteUser()  {
	id := AUC.GetString("id")
	token := AUC.GetString("token")
	num := models.IsAdminUser(id, token)
	if num == 0 {
		AUC.Data["json"] = models.SetResponse("", 2, "登录失效，请重新登录后重试")
	} else {
		did := AUC.GetString("did")
		err:=models.DeleteUser(did)
		if err != nil{
			panic(err)
		}
		AUC.Data["json"] = models.SetResponse("", 1, "用户删除成功")
	}
	AUC.ServeJSON()
}

//@Title NewUser
//@Description NewUser
// @Param token query  string true "token"
// @Param id query  string true "id"
// @Param body query  string true "UserInfo"
//@Success 200 {map}
//@Failure 403 body is empty
//@router /NewUser [post]
func (AUC *AdminUserControllers)NewUser()  {
	id := AUC.GetString("id")
	token := AUC.GetString("token")
	num := models.IsAdminUser(id, token)
	if num == 0 {
		AUC.Data["json"] = models.SetResponse("", 2, "登录失效，请重新登录后重试")
	} else {
		var dat map[string]string
		json.Unmarshal(AUC.Ctx.Input.RequestBody, &dat)
		err:=models.NewUser(dat)
		if err != nil{
			panic(err)
		}
		AUC.Data["json"] = models.SetResponse("", 1, "用户添加成功")
	}
	AUC.ServeJSON()
}

//@Title UpdateUser
//@Description Refund
// @Param token query  string true "token"
// @Param id query  string true "id"
// @Param uid query  string true "id"
// @Param body query  string true "UserInfo"
//@Success 200 {map}
//@Failure 403 body is empty
//@router /UpdateUser [post]
func (AUC *AdminUserControllers)UpdateUser()  {
	id := AUC.GetString("id")
	token := AUC.GetString("token")
	num := models.IsAdminUser(id, token)
	if num == 0 {
		AUC.Data["json"] = models.SetResponse("", 2, "登录失效，请重新登录后重试")
	} else {
		var dat map[string]string
		uid := AUC.GetString("uid")
		json.Unmarshal(AUC.Ctx.Input.RequestBody, &dat)
		err:=models.UpdateUser(uid,dat)
		if err != nil{
			panic(err)
		}
		AUC.Data["json"] = models.SetResponse("", 1, "用户信息更改成功")
	}
	AUC.ServeJSON()
}


//@Title GetAdminUserInfo
//@Description Refund
// @Param token query  string true "token"
// @Param id query  string true "id"
// @Param uid query  string true "uid"
//@Success 200 {map}
//@Failure 403 body is empty
//@router /GetAdminUserInfo [get]
func (AUC *AdminUserControllers)GetAdminUserInfo()  {
	id := AUC.GetString("id")
	token := AUC.GetString("token")
	num := models.IsAdminUser(id, token)
	if num == 0 {
		AUC.Data["json"] = models.SetResponse("", 2, "登录失效，请重新登录后重试")
	} else {
		uid := AUC.GetString("uid")
		maps,_:=models.GetAdminUserInfo(uid)
		AUC.Data["json"] = models.SetResponse(maps[0], 1, "")
	}
	AUC.ServeJSON()
}


//@Title WXBind
//@Description WXBind
//@Param phone query  string true "phone"
//@Param phone_code query  string true "phone_code"
//@Param code query  string true "code"
//@Success 200 {map}
//@Failure 403 body is empty
//@router /WXBind [get]
func (AUC *AdminUserControllers) WXBind() {
	phone,password,code:=AUC.GetString("phone"),AUC.GetString("phone_code"),AUC.GetString("code")
	num, userInfo := models.ConfirmUser(phone, password)
	if num == 0 {
		AUC.Data["json"] = models.SetResponse("", 2, "账号或密码错误")
	} else {
		//-1请求微信接口code换取openid和SESSIONKEY
		url := fmt.Sprintf("https://api.weixin.qq.com/sns/oauth2/access_token?appid=wxb4e8fb564d2b9cca&secret=29b49b5a99dc02b7d95703896696df8c&code=%s&grant_type=authorization_code", code)
		res, err := http.Get(url)
		if err == nil {
			defer res.Body.Close()
			m := make(map[string]interface{})
			body, _ := ioutil.ReadAll(res.Body)
			json.Unmarshal(body, &m)
			if m["errmsg"] == nil {
				openid := m["openid"]
				realUserInfo := userInfo[0]
				id := realUserInfo["id"].(string)
				err:=models.UpUserPushOpenid(id,openid.(string))
				if err==nil{
					AUC.Data["json"] = models.SetResponse("", 1, "绑定成功")
				}else{
					AUC.Data["json"] = models.SetResponse("", 5, "服务器飞往火星请稍后重试")
				}
			} else {
				AUC.Data["json"] = models.SetResponse("", 4, "获取微信授权失败请稍后重试")
			}
		} else {
			AUC.Data["json"] = models.SetResponse("", 4, "获取微信授权失败请稍后重试")
		}
	}
	AUC.ServeJSON()
}
