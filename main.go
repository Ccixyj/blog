package main

import (
	_ "BlogApi/models"
	_ "BlogApi/routers"
	"fmt"

	"BlogApi/models"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

func init() {
	orm.RegisterDataBase("default", "mysql", "root:root@tcp(127.0.0.1:3306)/blog")
	orm.Debug = true
	// create table
	err := orm.RunSyncdb("default", false, true)
	if err != nil {
		fmt.Println(err)
	}
}

func initAPP() {
	var admin = models.GenerateAdmin()
	o := orm.NewOrm()
	// 三个返回参数依次为：是否新创建的，对象 Id 值，错误
	//先 profile
	if created, id, err := o.ReadOrCreate(admin, "Name"); err == nil {
		if created {
			fmt.Println("New Insert an object. Id:", id)
			o.Insert(admin.Profile)
		} else {
			fmt.Println("Get an object. Id:", id)
		}
	}

}

func main() {
	initAPP()
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	beego.Run()
}
