package main

import (
	"blogx_server/core"
	"blogx_server/flags"
	"blogx_server/global"
	"blogx_server/router"
)

func main() {
	flags.Parse()
	//这里是要先读取配置参数，然后才能做各种初始化
	global.Config = core.ReadConf()
	core.InitLogrus()
	global.DB = core.InitDB()
	flags.Run()

	// 启动web程序
	router.Run()

}
