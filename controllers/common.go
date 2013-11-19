package controllers

import (
	"fmt"
	"github.com/linphy/admin/lib/rbac"
	"github.com/linphy/beego"
	"github.com/linphy/beego/orm"
)

type CommonController struct {
	beego.Controller
}

func (this *CommonController) Rsp(status bool, str string) {
	this.Data["json"] = &map[string]interface{}{"status": status, "info": str}
	this.ServeJson()
}

func init() {
	var dns string
	db_type := beego.AppConfig.String("db_type")
	db_host := beego.AppConfig.String("db_host")
	db_port := beego.AppConfig.String("db_port")
	db_user := beego.AppConfig.String("db_user")
	db_pass := beego.AppConfig.String("db_pass")
	db_name := beego.AppConfig.String("db_name")
	db_str := beego.AppConfig.String("db_str")
	switch db_type {
	case "mysql":
		dns = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8", db_user, db_pass, db_host, db_port, db_name)
		break
	case "oracle":
		dns = fmt.Sprintf("%s/%s@%s", db_user, db_pass, db_str)
		break
	}
	orm.RegisterDataBase("default", db_type, dns)
	//验证权限
	rbac.AccessRegister()
}
