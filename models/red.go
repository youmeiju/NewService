package models

import (
	"github.com/astaxie/beego/orm"
	"time"
)

func GetRedDetail(rid string)(err error,maps []orm.Params)  {
	o:=orm.NewOrm()
	_,err=o.Raw("SELECT * FROM `red_paper` rp,`user_red_paper` urp WHERE rp.id=urp.rid AND urp.`id` = ?",rid).Values(&maps)
	return
}


func GetRedDetail1(rid string)(err error,maps []orm.Params)  {
	o:=orm.NewOrm()
	_,err=o.Raw("SELECT * FROM `red_paper` rp WHERE rp.`id` = ?",rid).Values(&maps)
	return
}

func AddRed(uid string,rid string)(code int)  {
	o:=orm.NewOrm()
	var maps []orm.Params
	num1,_:=o.Raw("SELECT * FROM `user_red_paper` urp WHERE urp.uid =? AND urp.rid = ?",uid,rid).Values(&maps)
	if num1 == 0{
		_,err:=o.Raw("INSERT INTO `user_red_paper` VALUE(NULL,?,?,0)",uid,rid).Exec()
		if err==nil{
			code = 1
		}else{
			code = 5
		}
	}else{
		code = 4
	}
	return
}

func AddRed2(uid string,rid string){
	o:=orm.NewOrm()
	o.Raw("INSERT INTO `user_red_paper` VALUE(NULL,?,?,0)",uid,rid).Exec()
	return
}

func GetRedList(uid string)(err error,maps []orm.Params)  {
	o:=orm.NewOrm()
	_,err=o.Raw("select rp.* from `user_red_paper` urp,`red_paper` rp where urp.`rid` = rp.`id` and urp.uid=? and urp.has_used=0 and rp.`start_time`<=? AND rp.`end_time`>=? and rp.state=0",uid,time.Now().Format("2006-01-02"),time.Now().Format("2006-01-02 15:04:05")).Values(&maps)
	return
}

func GetUseRed(uid string,allprice string)(num int64,maps []orm.Params)  {
	o:=orm.NewOrm()
	num,_=o.Raw("SELECT rp.*,urp.`id` FROM `user_red_paper` urp,`red_paper` rp WHERE urp.`rid` = rp.`id` AND urp.uid=? AND rp.`limit_price`<? and urp.has_used=0 ORDER BY rp.`q_dai_jiner` DESC LIMIT 0,1",uid,allprice).Values(&maps)
	return
}

func UpUserRedCode(rid string,code string)()  {
	o:=orm.NewOrm()
	o.Raw("update `user_red_paper` set has_used = ? where id = ?",code,rid).Exec()
}

func GetUseRedList(uid string,allprice string)(err error,maps []orm.Params)  {
	o:=orm.NewOrm()
	_,err=o.Raw("SELECT rp.*,urp.id FROM `user_red_paper` urp,`red_paper` rp WHERE urp.`rid` = rp.`id` AND urp.uid=? AND urp.has_used=0 AND rp.`limit_price`<? and rp.`start_time`<=? AND rp.`end_time`>=?",uid,allprice,time.Now().Format("2006-01-02"),time.Now().Format("2006-01-02 15:04:05")).Values(&maps)
	return
}

func ReBackRed(oid string){
	o:=orm.NewOrm()
	var maps []orm.Params
	o.Raw("SELECT rp_id,u_id FROM `order` WHERE id = ?",oid).Values(&maps)
	rid :=maps[0]["rp_id"]
	if rid != 0{
		o.Raw("UPDATE `user_red_paper` SET has_used = 0 WHERE id=?",rid).Exec()
	}
	return
}