package controllers

import (
	"github.com/astaxie/beego"
	"encoding/json"
	"NewService/models"
	"strings"
	"fmt"
	"time"
	"math/rand"
	"net/http"
	"io/ioutil"
	"encoding/xml"
	"bytes"
	"sort"
	"crypto/md5"
	"encoding/hex"
	"strconv"
)


type OrderControllers struct {
	beego.Controller
}

// @Title NewOrder
// @Description 新建订单
// @Param token query  string true "token"
// @Param utoken query string true "utoken"
// @Param orderInfo body map true "body for orderInfo"
// @Success 200 {map}
// @Failure 403 body is empty
// @router /NewOrder [post]
func (OC *OrderControllers) NewOrder() {
	token := OC.GetString("token")
	if IsToken(token) {
		utoken := OC.GetString("utoken")
		num := models.IsUserToken(token, utoken)
		if num == 0 {
			OC.Data["json"] = models.SetResponse("", 2, "token不存在正在重新登录")
		} else {
			uid := models.GetUserId(utoken)[0]["id"].(string)
			var dat map[string]interface{}
			json.Unmarshal(OC.Ctx.Input.RequestBody, &dat)
			rid := dat["rid"].(string)
			address := dat["address_info"].(map[string]interface{})
			goods := dat["goods"].( []interface{})
			allPrice := dat["all_price"]
			postInfo := dat["postInfo"]
			cityName := address["cityName"]
			countyName := address["countyName"]
			detailInfo := address["detailInfo"]
			provinceName := address["provinceName"]
			telNumber := address["telNumber"]
			userName := address["userName"]
			err,oId := models.AddOrder(uid,userName,telNumber,provinceName,cityName,countyName,detailInfo,allPrice,rid,postInfo)
			for _, gInfo := range goods {
				gi:=gInfo.(map[string]interface{})
				id := gi["id"]
				shopPrice, _ := gi["shop_price"].(string)
				buyCount := gi["buy_count"]
				models.AddOrderDetail(oId[0]["MAX(id)"].(string),id,shopPrice,buyCount)
			}
			if err==nil{
				num := models.GetOrderNum(uid)
				if(num[0]["num"]=="1"){
					iid:=models.GetInviteId(uid)
					if iid[0]["invite_id"] != nil{
						models.AddRed2(iid[0]["invite_id"].(string),"3")
					}
				}
				models.UpUserRedCode(rid,"1")
				OC.Data["json"]=models.SetResponse(oId[0]["MAX(id)"], 1, "")
			}else {
				OC.Data["json"]=models.SetResponse("", 5, "网络飞往火星，请稍后重试")
			}
		}
	}
	OC.ServeJSON()
}
// @Title GetOrderInfo
// @Description GetOrderInfo
// @Param token query  string true "token"
// @Param utoken query string true "utoken"
// @Param oid query string true "utoken"
// @Success 200 {map}
// @Failure 403 body is empty
// @router /GetOrderInfo [get]
func (OC *OrderControllers)GetOrderInfo()  {
	token := OC.GetString("token")
	if IsToken(token) {
		utoken := OC.GetString("utoken")
		num := models.IsUserToken(token, utoken)
		if num == 0 {
			OC.Data["json"] = models.SetResponse("", 2, "token不存在正在重新登录")
		} else {
			err,orderInfo:=models.GetOrderInfo(OC.GetString("oid"))
			if err == nil{
				OC.Data["json"] = models.SetResponse(orderInfo[0], 1, "")
			}
		}
	}
	OC.ServeJSON()
}
// @Title OrderPay
// @Description OrderPay
// @Param token query  string true "token"
// @Param utoken query string true "utoken"
// @Param oid query string true "oid"
// @Success 200 {map}
// @Failure 403 body is empty
// @router /OrderPay [get]
func (OC *OrderControllers)OrderPay()  {
	req := OC.Ctx.Request
	addr := req.RemoteAddr // "IP:port" "192.168.1.150:8889"
	ip:=strings.Split(addr,":")[0]
	token := OC.GetString("token")
	if IsToken(token) {
		utoken := OC.GetString("utoken")
		num := models.IsUserToken(token, utoken)
		if num == 0 {
			OC.Data["json"] = models.SetResponse("", 2, "token不存在正在重新登录")
		} else {
			user:=models.GetOpenId(utoken)
			oid:=OC.GetString("oid")
			err,orderInfo:=models.GetOrderInfo(oid)
			var reqMap = make(map[string]string, 0)
			reqMap["appid"] = "wx331b51ae3849285d"//微信小程序appid
			reqMap["body"] = "订单号"+oid          //商品描述
			reqMap["mch_id"] = "1513958441"     //商户号
			rnd := rand.New(rand.NewSource(time.Now().UnixNano()))
			reqMap["nonce_str"] = time.Now().Format("20060102150405") + fmt.Sprintf("%04v", rnd.Int31n(1000))    //随机数
			reqMap["notify_url"] = "https://service.shanghaiyoumeiju2018.com/v1/order/UpdateOrderStatue"   //通知地址
			reqMap["openid"] = user[0]["openid"].(string)       //商户唯一标识 openid
			reqMap["out_trade_no"] = oid   //订单号
			reqMap["spbill_create_ip"] = ip     //用户端ip   //订单生成的机器 IP
			totalPrice,_ :=strconv.Atoi(orderInfo[0]["total_price"].(string))
			reqMap["total_fee"] = strconv.Itoa(totalPrice*100)  //订单总金额，单位为分
			reqMap["trade_type"] = "JSAPI"      //trade_type=JSAPI时（即公众号支付），此参数必传，此参数为微信用户在商户对应appid下的唯一标识
			reqMap["sign"] = WxPayCalcSign(reqMap,"zw620521198608071116000000000000")
			reqStr := Map2Xml(reqMap)
			// 调用支付统一下单API
			req, err := http.NewRequest("POST", "https://api.mch.weixin.qq.com/pay/unifiedorder", strings.NewReader(reqStr))
			if err != nil {
				// handle error
			}
			req.Header.Set("Content-Type", "text/xml;charset=utf-8")
			client := &http.Client{}
			resp, err := client.Do(req)
				defer resp.Body.Close()
			body, err := ioutil.ReadAll(resp.Body)
			resp1 := models.UnifyOrderResp{}
			err = xml.Unmarshal(body, &resp1)
			if strings.ToUpper(resp1.Return_code) == "SUCCESS"{
				// 再次签名
				var resMap = make(map[string]string)
				resMap["appId"] = "wx331b51ae3849285d"
				resMap["nonceStr"] = resp1.Nonce_str         //商品描述
				resMap["package"] = "prepay_id=" + resp1.Prepay_id //商户号
				resMap["signType"] = "MD5"              //签名类型
				resMap["timeStamp"] = strconv.FormatInt(time.Now().Unix(),10)    //当前时间戳
				resMap["paySign"] = WxPayCalcSign(resMap,"zw620521198608071116000000000000") // 返回5个支付参数及sign 用户进行确认支付
				OC.Data["json"] = models.SetResponse(resMap,1,"")
			}else{
				OC.Data["json"] = models.SetResponse("",5,"微信接口请求失败")
			}
		}
	}
	OC.ServeJSON()
}

type WXPayNotifyResp struct {
	Return_code string `xml:"return_code"`
	Return_msg  string `xml:"return_msg"`
}
// @Title UpdateOrderStatue
// @Description UpdateOrderStatue
// @Success 200 {map}
// @Failure 403 body is empty
// @router /UpdateOrderStatue [post]
func(OC *OrderControllers)UpdateOrderStatue(){
	req := models.WXPayNotifyReq{}
	err := xml.Unmarshal(OC.Ctx.Input.RequestBody, &req)
	if err != nil{
		panic(err)
	}
	//{wx331b51ae3849285d ABC_DEBIT 100 CNY N 1513958441 201809131441270059 o6GKW5Cpv7HhTbw2g9nPZ4IK3N90 80 SUCCESS SUCCESS 283C2E372DDA3B52A32C26772BC81E43 20180913144152 100 JSAPI 4200000202201809131526570757}
	if req.Return_code == "SUCCESS"{
		if req.Result_code== "SUCCESS"{
			oid:=req.Out_trade_no
			err:=models.UpdateOrderStatue(oid,"001")
			if err == nil{
				adminUser,_:=models.GetAdminUser()
				for _,v := range adminUser{
					pushMsg(oid,v["push_openid"])
				}
				OC.Data["xml"] = WXPayNotifyResp{"SUCCESS",""}
			}
		}
	}
	OC.ServeXML()
}
//微信支付计算签名的函数
func Map2Xml(mReq map[string]string) (xml string) {
	sb := bytes.Buffer{}
	sb.WriteString("<xml>")
	for k,v := range mReq{
		sb.WriteString("<"+k+">"+v+"</"+k+">")
	}
	sb.WriteString("</xml>")
	return sb.String()
}
//微信支付计算签名的函数
func WxPayCalcSign(mReq map[string]string, key string) (sign string) {
	//STEP 1, 对key进行升序排序.
	sorted_keys := make([]string, 0)
	for k, _ := range mReq {
		sorted_keys = append(sorted_keys, k)
	}
	sort.Strings(sorted_keys)
	//STEP2, 对key=value的键值对用&连接起来，略过空值
	var signStrings string
	for _, k := range sorted_keys {
		value := fmt.Sprintf("%v", mReq[k])
		if value != "" {
			signStrings = signStrings + k + "=" + value + "&"
		}
	}
	//STEP3, 在键值对的最后加上key=API_KEY
	if key != "" {
		signStrings = signStrings + "key=" + key
	}
	//STEP4, 进行MD5签名并且将所有字符转为大写.
	md5Ctx := md5.New()
	md5Ctx.Write([]byte(signStrings)) //
	cipherStr := md5Ctx.Sum(nil)
	upperSign := strings.ToUpper(hex.EncodeToString(cipherStr))
	return upperSign
}

// @Title GetOrderList
// @Description GetOrderList
// @Param token query  string true "token"
// @Param utoken query string true "utoken"
// @Success 200 {map}
// @Failure 403 body is empty
// @router /GetOrderList [get]
func (OC *OrderControllers)GetOrderList(){
	token := OC.GetString("token")
	if IsToken(token) {
		utoken := OC.GetString("utoken")
		num := models.IsUserToken(token, utoken)
		if num == 0 {
			OC.Data["json"] = models.SetResponse("", 2, "token不存在正在重新登录")
		} else {
			id:=models.GetUserId(utoken)[0]["id"].(string)
			_,orders:=models.GetOrderList(id)
			for i,v := range orders{
				_,orderGoodsInfo:=models.GetOrderGoodsInfo(v["id"].(string))
				realGoodsMsg :=  make([]map[string]interface{},len(orderGoodsInfo))
				for i,v := range orderGoodsInfo{
					a := make(map[string]interface{})
					a["buy_count"] = v["buy_count"]
					a["shop_price"] = v["shop_price"]
					goodsInfo,_:=models.GetGoodsDetail(v["g_id"].(string))
					a["goodsInfo"] = goodsInfo[0]
					realGoodsMsg[i] = a
				}
				orders[i]["glist"] = realGoodsMsg
			}
			OC.Data["json"] = models.SetResponse(orders, 1, "")
		}
	}
	OC.ServeJSON()
}
// @Title GetOrderDetail
// @Description GetOrderDetail
// @Param token query  string true "token"
// @Param utoken query string true "utoken"
// @Param oid query string true "oid"
// @Success 200 {map}
// @Failure 403 body is empty
// @router /GetOrderDetail [get]
func (OC *OrderControllers)GetOrderDetail()  {
	token := OC.GetString("token")
	if IsToken(token) {
		utoken := OC.GetString("utoken")
		num := models.IsUserToken(token, utoken)
		if num == 0 {
			OC.Data["json"] = models.SetResponse("", 2, "token不存在正在重新登录")
		} else {
			oid := OC.GetString("oid")
			err,orderInfo:=models.GetOrderInfo(oid)
			if err==nil {
				_,orderGoodsInfo:=models.GetOrderGoodsInfo(oid)
				realGoodsMsg :=  make([]map[string]interface{},len(orderGoodsInfo))
				for i,v := range orderGoodsInfo{
					a := make(map[string]interface{})
					a["buy_count"] = v["buy_count"]
					a["shop_price"] = v["shop_price"]
					goodsInfo,_:=models.GetGoodsDetail(v["g_id"].(string))
					a["goodsInfo"] = goodsInfo[0]
					realGoodsMsg[i] = a
				}
				data := make(map[string]interface{})
				data["orderInfo"] = orderInfo
				data["lGoodsInfo"] = realGoodsMsg
				OC.Data["json"] = models.SetResponse(data,1,"")
			}else{
				panic(err)
			}
		}
	}
	OC.ServeJSON()
}
// @Title UpdateOrderStatues
// @Description UpdateOrderStatues
// @Param token query  string true "token"
// @Param utoken query string true "utoken"
// @Param status query string true "status"
// @Param oid query string true "oid"
// @Success 200 {map}
// @Failure 403 body is empty
// @router /UpdateOrderStatues [get]
func (OC *OrderControllers)UpdateOrderStatues()  {
	token := OC.GetString("token")
	if IsToken(token) {
		utoken := OC.GetString("utoken")
		num := models.IsUserToken(token, utoken)
		if num == 0 {
			OC.Data["json"] = models.SetResponse("", 2, "token不存在正在重新登录")
		} else {
			status:=OC.GetString("status")
			err:=models.UpdateOrderStatue(OC.GetString("oid"),status)
			models.ReBackRed(OC.GetString("oid"))
			if err!=nil{
				panic(err)
			}
			OC.Data["json"] = models.SetResponse("", 1 ,"")
		}
	}
	OC.ServeJSON()
}

