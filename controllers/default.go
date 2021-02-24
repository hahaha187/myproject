package controllers

import (
	"fmt"
	"myproject/models"

	"github.com/astaxie/beego/client/orm"

	beego "github.com/beego/beego/v2/server/web"
)

//MainController 拥有beego.controller的方法
type MainController struct {
	beego.Controller
}

//Get 重写之后的
func (c *MainController) Get() {
	// c.Data["Website"] = "beego.me"
	// c.Data["Email"] = "astaxie@gmail.com"
	// c.TplName = "index.tpl" //设置显示的模版文件

	/*
	   //charu
	   	//1.有orm对象
	   	o := orm.NewOrm()
	   	//2.又一个要插入数据的结构体对象
	   	user := models.User{}
	   	//3.对结构体肤质
	   	user.Name="娄大皮炎"
	   	user.Pwd="1234"
	   	//4.插入
	   	_,err := o.Insert(&user)
	   	if err!= nil{
	   		fmt.Printf("insert failed ,err:%v",err)
	   		return
	   	}
	*/
	/*
	   //chaxun
	   	o := orm.NewOrm()
	   	user := models.User{}
	   	user.ID = 1
	   	//user.name="louhuiying"
	   	//err := o.Read(&user,"Name")
	   	err := o.Read(&user)
	   	if err !=nil {
	   		fmt.Printf("read failed , err:%v",err)
	   		return
	   	}
	   	fmt.Println(user.Name,user.Pwd)
	*/
	/*
	   //gengxin
	   	o := orm.NewOrm()
	   	user := models.User{}
	   	user.ID = 1
	   	err := o.Read(&user)
	   	if err == nil {
	   		user.Name="娄慧莹"
	   		user.Pwd="123"
	   		_,err :=o.Update(&user)
	   		if err!= nil {
	   			fmt.Printf("update failed,err:%v",err)
	   			return
	   		}
	   	}
	*/
	/*
		//shanchu
		o := orm.NewOrm()
		user := models.User{}
		user.ID=2
		_,err := o.Delete(&user)
		if err != nil {
			fmt.Printf("delete failed ,err: %v",err)
			return
		}

		c.Data["data"] = "娄慧莹臭智障"
		c.TplName = "test.html"
	*/

	//注册
	c.TplName = "register.html"
}

//Post fangfa
func (c *MainController) Post() {
	// c.Data["data"] = "今天吃娄慧莹"
	// c.TplName = "test.html"

	//点击注册提交form表单
	//1。拿到数据
	userName := c.GetString("userName")
	pwd := c.GetString("pwd")
	//2.对数据进行校验
	if userName == "" || pwd == "" {
		fmt.Printf("数据不能为空")
		c.Redirect("/register", 302) //跳转界面不能传递数据
		return
	}
	//3.插入数据库
	o := orm.NewOrm()
	user := models.User{}
	user.Name = userName
	user.Pwd = pwd
	_,err := o.Insert(&user)
	if err != nil {
		fmt.Printf("插入失败")
		c.Redirect("/register", 302)
		return
	}
	//4.返回登陆界面
	c.Ctx.WriteString("注册成功") //应该是返回登陆界面
}
//ShowLogin 自定义的get方法
func (c *MainController) ShowLogin() {
	c.TplName = "login.html" //指定视图文件同时可以给这个视图传递一些数据
}
//HandleLogin 自定义的post方法
func (c *MainController) HandleLogin() {
	//c.Ctx.WriteString("this is login post")
	//1.拿到数据
	userName := c.GetString("userName")
	pwd := c.GetString("pwd")
	//2.判断数据是否合法
	if userName == "" || pwd == "" {
		fmt.Printf("输入不合法")
		c.TplName = "login.html"
		return
	}
	//3.查询账号密码是否正确
	o := orm.NewOrm()
	user := models.User{}
	user.Name = userName
	err := o.Read(&user,"Name")
	if err!=nil {
		fmt.Printf("查询失败")
		c.TplName="login.html"
		return
	}
	//4.跳转界面
	c.Ctx.WriteString("登陆成功")
}