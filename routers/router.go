package routers

import (
	"myproject/controllers"

	beego "github.com/beego/beego/v2/server/web"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/abc", &controllers.MainController{})
	beego.Router("/register", &controllers.MainController{})
	//注意：当是实现了自定义的get请求方法，请求将不会访问默认方法
	beego.Router("/login", &controllers.MainController{},"get:ShowLogin;post:HandleLogin")
}
