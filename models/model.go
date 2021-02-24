package models

import (
	"github.com/astaxie/beego/client/orm"
//
	_ "github.com/go-sql-driver/mysql"
)
//User 定义一个结构体
type User struct {
	ID int
	Name string
	Pwd string
}

func init(){
	// 设置数据库基本信息
    orm.RegisterDataBase("default", "mysql", "root:lhy871601318@tcp(127.0.0.1:3306)/test1?charset=utf8")
  
    // register model映射model数据
    orm.RegisterModel(new(User))

    // create table 生成表
    orm.RunSyncdb("default", false, true)
}