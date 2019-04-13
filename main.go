package main

import (
	_ "NewService/routers"

	"github.com/astaxie/beego"
	"NewService/models"
)

func init(){
	models.RegisterDB()
}

func main() {
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	//日志级别 1.LevelEmergency 2.LevelAlert 3.LevelCritical 4.LevelError 5.LevelWarning 6.LevelNotice 7.LevelInformational 8.LevelDebug
	//beego.SetLevel(beego.LevelEmergency)
	//是否输出行号
	beego.SetLogFuncCall(true)
	//输出方式
	beego.SetLogger("file", `{"filename":"logs/xxx.log"}`)
	//静态文件处理
	beego.SetStaticPath("/img", "img")
	beego.Run()
}