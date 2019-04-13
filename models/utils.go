package models

import (
	"github.com/astaxie/beego/orm"
	_"github.com/go-sql-driver/mysql"
	"fmt"
	"time"
	"strconv"
	"math/rand"
)
//注冊数据库
func RegisterDB() {
	//注册驱动
	orm.RegisterDriver("mysql", orm.DRMySQL)
	//注册默认数据库
	orm.RegisterDataBase("default", "mysql", "root:123456@tcp(118.31.72.110:3306)/YM?charset=utf8")//密码为空格式
}
//Response setting
func SetResponse(data interface{},code int,content string)(res map[string]interface{}){
	res = make(map[string]interface{})
	res["Data"] = data
	res["Code"] = code
	res["Msg"]= content
	return
}
//Error control
func RealMsg(err error)(Msg string, Code int)  {
	if err != nil{
		fmt.Println(err)
		Msg = "服务器异常"
		Code = 400
		return
	}else {
		Msg = ""
		Code = 200
		return
	}
}
//生成用户Token
func SetUserToken(id string)(token string)  {
	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))
	vcode := fmt.Sprintf("%04v", rnd.Int31n(1000))
	token = strconv.FormatInt(time.Now().Unix(),10)+id+vcode
	return
}