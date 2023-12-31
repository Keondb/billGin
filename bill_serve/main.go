package main

import (
	"bill/bill_serve/core"
	"bill/bill_serve/flag"
	"bill/bill_serve/global"
	"bill/bill_serve/routers"
)

func main() {
	// 读取配置文件
	core.InitConf()
	// 初始化日志
	global.Log = core.InitLogger()
	//连接数据库
	global.DB = core.InitGorm()

	//命令行参数绑定
	option := flag.Parse()
	if flag.IsWebStop(option) {
		flag.SwitchOption(option)
		return
	}

	//路由初始化
	router := routers.InitRouter()
	addr := global.Config.System.Addr()
	global.Log.Infof("程序运行在: %s", addr)

	err := router.Run(addr)
	if err != nil {
		global.Log.Fatalf(err.Error())
	}
}
