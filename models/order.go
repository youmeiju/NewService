package models

import(
	"github.com/astaxie/beego/orm"
	"time"
)

type UnifyOrderResp struct {
	Return_code string `xml:"return_code"`
	Return_msg  string `xml:"return_msg"`
	Appid       string `xml:"appid"`
	Mch_id      string `xml:"mch_id"`
	Nonce_str   string `xml:"nonce_str"`
	Sign        string `xml:"sign"`
	Result_code string `xml:"result_code"`
	Prepay_id   string `xml:"prepay_id"`
	Trade_type  string `xml:"trade_type"`
}

type RefundOrderResp struct {
	Return_code string `xml:"return_code"`
	Return_msg  string `xml:"return_msg"`
	Result_code string `xml:"result_code"`
	Err_code       string `xml:"err_code"`
	Err_code_des      string `xml:"err_code_des"`
	Sign        string `xml:"sign"`
}

type WXPayNotifyReq struct {
	Appid          string `xml:"appid"`
	Bank_type      string `xml:"bank_type"`
	Cash_fee       float64   `xml:"cash_fee"`
	Fee_type       string `xml:"fee_type"`
	Is_subscribe   string `xml:"is_subscribe"`
	Mch_id         string `xml:"mch_id"`
	Nonce_str      string `xml:"nonce_str"`
	Openid         string `xml:"openid"`
	Out_trade_no   string `xml:"out_trade_no"`
	Result_code    string `xml:"result_code"`
	Return_code    string `xml:"return_code"`
	Sign           string `xml:"sign"`
	Time_end       string `xml:"time_end"`
	Total_fee      float64 `xml:"total_fee"`
	Trade_type     string `xml:"trade_type"`
	Transaction_id string `xml:"transaction_id"`
	Req_info       string `xml:"req_info"`
}



func AddOrder(id interface{},userName interface{},telNumber interface{},provinceName interface{},cityName interface{},countyName interface{},detailInfo interface{},allPrice interface{},rid interface{},postInfo interface{})(err error,maps []orm.Params) {
	o:=orm.NewOrm()
	_,err=o.Raw("INSERT INTO `order` VALUES(NULL,?,?,?,?,?,?,?,?,?,?,?,?,NULL,?)",rid,id,userName,telNumber,provinceName,cityName,countyName,detailInfo,time.Now().Format("2006-01-02 15:04:05.000"),time.Now().Format("2006-01-02 15:04:05.000"),allPrice,"000",postInfo).Exec()
	o.Raw("SELECT MAX(id) FROM `order` WHERE  u_id = ?",id).Values(&maps)
	return
}

func AddOrderDetail(oId interface{},gId interface{},shopPrice interface{},buyCount interface{})(err error)  {
	o:=orm.NewOrm()
	_,err = o.Raw("INSERT INTO `order_detail` VALUES(NULL,?,?,?,?)",oId,gId,shopPrice,buyCount).Exec()
	return
}

func GetOrderInfo(oid string)(err error,maps []orm.Params)  {
	o:=orm.NewOrm()
	_,err=o.Raw("select * from `order` where id = ?",oid).Values(&maps)
	return
}

func GetOrderList(id string)(err error,maps []orm.Params)  {
	o:=orm.NewOrm()
	_,err=o.Raw("select * from `order` where u_id = ? order by id desc",id).Values(&maps)
	return
}
func GetOrderGoodsInfo(oid string)(err error,maps []orm.Params)  {
	o:=orm.NewOrm()
	_,err=o.Raw("select * from `order_detail` where o_id = ?",oid).Values(&maps)
	return
}

func UpdateOrderStatue(oid string,statue string)(err error)  {
	o:=orm.NewOrm()
	_,err=o.Raw("UPDATE `order` SET order_status = ? WHERE id =  ?",statue,oid).Exec()
	return
}

func GetOrderNum(id string)(maps []orm.Params)  {
	o:=orm.NewOrm()
	o.Raw("SELECT COUNT(id) AS num FROM `order` WHERE u_id = ?",id).Values(&maps)
	return
}


