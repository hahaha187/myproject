package main

import (
	_ "myproject/routers"
	_"myproject/models" //默认运行时会先调用models包中的init函数
	beego "github.com/beego/beego/v2/server/web"
)

func main() {
	beego.Run()
}

