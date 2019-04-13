package models

import (
	"github.com/astaxie/beego/orm"
	"fmt"
)
//判断token是否存在
func IsUserToken(token string,utoken string)(num int64){
	var list []orm.ParamsList
	o := orm.NewOrm()
	num,_ = o.Raw("select * from `user` where utoken = ?",utoken).ValuesList(&list)
	return
}
//获取用户信息
func GetUserInfo(openid string)(num int64,maps []orm.Params)  {
	o := orm.NewOrm()
	num,_ = o.Raw("SELECT * FROM `user` WHERE openid = ? ",openid).Values(&maps)
	return
}
//添加用户
func AddUser(openid string)(err error)  {
	o := orm.NewOrm()
	_,err = o.Raw("INSERT INTO `user` VALUES(NULL,?,NULL,NULL,NULL)",openid).Exec()
	return
}
//更新用户Token
func UpdateUserToken(id string,token string)(err error){
	o := orm.NewOrm()
	_,err = o.Raw("UPDATE `user` SET utoken = ? WHERE id = ? ",token,id).Exec()
	fmt.Println(err,1)
	return
}
//get user id
func GetUserId(token string)(maps []orm.Params)  {
	o := orm.NewOrm()
	o.Raw("select id from user where utoken = ? ", token).Values(&maps)
	return
}

func GetOpenId(token string)(maps []orm.Params)  {
	o := orm.NewOrm()
	o.Raw("select openid from user where utoken = ? ", token).Values(&maps)
	return
}

func GetInviteId(uid string)(maps []orm.Params){
	o := orm.NewOrm()
	o.Raw("SELECT invite_id FROM `user` WHERE id = ?", uid).Values(&maps)
	return
}

func InsertInviteId(uid string,iid string)(err error)  {
	o := orm.NewOrm()
	_,err = o.Raw("UPDATE `user` SET invite_id = ? WHERE id =? ",iid, uid).Exec()
	return
}