package models

import (
	"github.com/astaxie/beego/orm"
	"fmt"
	"strconv"
	"time"
)

//IndexInfo
func GetIndexInfo(uid string)(Map map[string]interface{},err error)  {
	Map = make(map[string]interface{})
	var index_cate_list []orm.Params
	var index_swiper_list []orm.Params
	var index_quan_list []orm.Params
	var index_tuijian_list []orm.Params
	var index_hot_list []orm.Params
	var index_new_list []orm.Params
	o := orm.NewOrm()
	_,err = o.Raw("SELECT gi.`id` id,gi.`g_img`,gi.`g_name`,gi.`shop_price` FROM `index_goods` ig,`goods_info` gi WHERE ig.`goods_id` = gi.`id` AND ig.`type`=1").Values(&index_tuijian_list)
	_,err = o.Raw("SELECT gi.`id` id,gi.`g_img`,gi.`g_name`,gi.`shop_price` FROM `index_goods` ig,`goods_info` gi WHERE ig.`goods_id` = gi.`id` AND ig.`type`=2").Values(&index_hot_list)
	_,err = o.Raw("SELECT gi.`id` id,gi.`g_img`,gi.`g_name`,gi.`shop_price` FROM `index_goods` ig,`goods_info` gi WHERE ig.`goods_id` = gi.`id` AND ig.`type`=3").Values(&index_new_list)
	//../malldetail/malldetail 商品详情 ../malllist/malllist?cid= 分类
	_,err = o.Raw("SELECT asa.`image_url` img_url,IF(asa.`link_type`=1,CONCAT('../malldetail/malldetail?sid=',asa.`link_content`),CONCAT('../malllist/malllist?cid='+ asa.`link_content`)) url FROM `app_setting` asa WHERE asa.`type` = 1").Values(&index_swiper_list)
	_,err = o.Raw("SELECT asa.`image_url`,asa.`name`,asa.`link_content` id FROM `app_setting` asa WHERE asa.`type` = 2").Values(&index_cate_list)
	_,err = o.Raw("SELECT rp.* FROM `red_paper` rp LEFT JOIN `user_red_paper` urp ON rp.`id`=urp.`rid` AND urp.`uid`=? WHERE urp.`uid` IS NULL AND rp.`start_time`<=? AND rp.`end_time`>=? and rp.`red_type` = 1 ",uid,time.Now().Format("2006-01-02"),time.Now().Format("2006-01-02 15:04:05")).Values(&index_quan_list)
	Map["index_tuijian_list"] = index_tuijian_list
	Map["index_hot_list"] = index_hot_list
	Map["index_new_list"] = index_new_list
	Map["index_swiper_list"] = index_swiper_list
	Map["index_cate_list"] = index_cate_list
	Map["index_quan_list"] = index_quan_list
	return
}
//GetGoodsListByKeyWord
func GetGoodsListByKeyWord(pagesize string,pagenum string,keywords string,stype string,stype_jiage string)(Map []orm.Params,err error){
	o:=orm.NewOrm()
	size,_:=strconv.Atoi(pagesize)
	num,_:=strconv.Atoi(pagenum)
	index := num*(size-1)
	keywords = "%"+keywords+"%"
	order :=stype+" "+stype_jiage
	sql:=fmt.Sprintf("select * from goods_info where g_name like \"%s\"  order by %s LIMIT %d,%d",keywords,order,index,num)
	_,err= o.Raw(sql).Values(&Map)
	return
}
//GoodsList
func GetFirstMenu()(Map []orm.Params,err error){
	o := orm.NewOrm()
	_,err = o.Raw("select distinct(fm.`id`),fm.`name` from `first_menu` fm,`goods_info` gi where fm.`id` = gi.`g_mf_id` ").Values(&Map)
	return
}
//GetMenuGoodsList
func GetGoodsListByKeyId(pagesize string,pagenum string,id string,stype string,stype_jiage string)(Map []orm.Params,err error){
	o := orm.NewOrm()
	size,_:=strconv.Atoi(pagesize)
	num,_:=strconv.Atoi(pagenum)
	index := num*(size-1)
	order :=stype+" "+stype_jiage
	sql:=fmt.Sprintf("select * from goods_info where g_mf_id=%s order by %s LIMIT %d,%d",id,order,index,num)
	_,err= o.Raw(sql).Values(&Map)
	return
}
//GetGoodsDetail
func GetGoodsDetail(id string)(Map []orm.Params,err error)  {
	o := orm.NewOrm()
	_,err = o.Raw("select * from `goods_info` where id =?",id).Values(&Map)
	return
}