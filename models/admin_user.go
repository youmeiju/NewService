package models

import (
	"github.com/astaxie/beego/orm"
)

func ConfirmUser(user string,password string)(num int64,maps []orm.Params)  {
	o:= orm.NewOrm()
	num,_=o.Raw("select * from `admin_user` au where au.`tel_num` = ? and au.`pass_word` = ?",user,password).Values(&maps)
	return
}

func UpAdminUser(id string,token string)(err error)  {
	o:=orm.NewOrm()
	_,err = o.Raw("  UPDATE `admin_user` au SET token =? WHERE id = ? ",token,id).Exec()
	return
}

func IsAdminUser(id string,token string)(num int64)  {
	o:=orm.NewOrm()
	var maps []orm.Params
	num,_ = o.Raw("select * from `admin_user` where id =? and token = ?",id,token).Values(&maps)
	return
}

func GetOrderNumByState()(num int64)  {
	o:=orm.NewOrm()
	var maps []orm.Params
	num,_ = o.Raw("SELECT * FROM `order` WHERE order_status = 001 OR order_status=002").Values(&maps)
	return
}

func GetOrderNumById(id string)(num int64)  {
	o:=orm.NewOrm()
	var maps []orm.Params
	num,_ = o.Raw("SELECT * FROM `order` WHERE a_user = ?",id).Values(&maps)
	return
}

func GetAdminOrderList()(err error,maps []orm.Params)  {
	o:=orm.NewOrm()
	_,err=o.Raw("SELECT * FROM `order` WHERE order_status = 001 OR order_status=002 ORDER BY id DESC").Values(&maps)
	return
}

func GetNotAdminOrderList(id string)(err error,maps []orm.Params) {
	o:=orm.NewOrm()
	_,err=o.Raw("SELECT * FROM `order` WHERE a_user=? ORDER BY id DESC",id).Values(&maps)
	return
}

func GetUser(code1,code2 string)(maps []orm.Params)  {
	o:=orm.NewOrm()
	o.Raw("SELECT id,tel_num,real_name FROM `admin_user` WHERE is_surper =? OR is_surper =?",code1,code2).Values(&maps)
	return
}

func UpdateOrderUser(oid string,aid string)(err error)  {
	o:=orm.NewOrm()
	_,err=o.Raw("UPDATE `order` SET a_user = ?,order_status=? WHERE id = ?",aid,"002",oid).Exec()
	return
}

func DeleteUser(did string)(err error){
	o:=orm.NewOrm()
	_,err=o.Raw("DELETE FROM `admin_user` WHERE id =?",did).Exec()
	return
}

func NewUser(dat map[string]string)(err error){
	o:=orm.NewOrm()
	_,err=o.Raw("INSERT INTO `admin_user` VALUE(NULL,?,?,NULL,?,NULL,?,NULL)",dat["tel_num"],dat["pass_word"],dat["is_super"],dat["real_name"]).Exec()
	return
}

func UpdateUser(uid string,dat map[string]string)(err error){
	o:=orm.NewOrm()
	_,err=o.Raw("UPDATE `admin_user` SET tel_num=?,pass_word=?,is_surper = ?,real_name=? WHERE id = ?",dat["tel_num"],dat["pass_word"],dat["is_super"],dat["real_name"],uid).Exec()
	return
}

func GetAdminUserInfo(uid string)(maps []orm.Params,err error){
	o:=orm.NewOrm()
	_,err=o.Raw("select * from `admin_user` where id = ?",uid).Values(&maps)
	return
}

func UpUserPushOpenid(id string,openid string)(err error){
	o:=orm.NewOrm()
	_,err=o.Raw("UPDATE `admin_user` SET push_openid=? WHERE id = ?",openid,id).Exec()
	return
}

func GetAdminUser()(maps []orm.Params,err error)  {
	o:=orm.NewOrm()
	_,err = o.Raw("SELECT push_openid FROM `admin_user` WHERE push_openid IS NOT NULL AND push_openid != ? AND is_surper=1","").Values(&maps)
	return
}

func GetUserPushOpenid(id string)(maps []orm.Params,err error)  {
	o:=orm.NewOrm()
	_,err = o.Raw("SELECT push_openid FROM `admin_user` WHERE id =?",id).Values(&maps)
	return
}