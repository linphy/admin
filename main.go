package main

import (
	"fmt"
	"github.com/linphy/beego"
	//"github.com/linphy/beego/orm"
	"github.com/linphy/admin/controllers"
	"github.com/linphy/admin/controllers/rbac"
	"github.com/linphy/admin/lib"
	"github.com/linphy/admin/models/rbacmodels"
	_ "github.com/linphy/go-oci8"
	"os"
)

func main() {

	//orm.Debug = true
	fmt.Println("Starting....")
	//orm.RegisterDataBase("default", "mysql", "root:root@/admin?charset=utf8")
	beego.Router("/", &controllers.MainController{})
	beego.Router("/public/index", &controllers.MainController{}, "*:Index")

	beego.Router("/rbac/user/AddUser", &rbac.UserController{}, "*:AddUser")
	beego.Router("/rbac/user/UpdateUser", &rbac.UserController{}, "*:UpdateUser")
	beego.Router("/rbac/user/DelUser", &rbac.UserController{}, "*:DelUser")
	beego.Router("/rbac/user/index", &rbac.UserController{}, "*:Index")

	beego.Router("/rbac/node/AddAndEdit", &rbac.NodeController{}, "*:AddAndEdit")
	beego.Router("/rbac/node/DelNode", &rbac.NodeController{}, "*:DelNode")
	beego.Router("/rbac/node/index", &rbac.NodeController{}, "*:Index")

	beego.Router("/rbac/group/AddGroup", &rbac.GroupController{}, "*:AddGroup")
	beego.Router("/rbac/group/UpdateGroup", &rbac.GroupController{}, "*:UpdateGroup")
	beego.Router("/rbac/group/DelGroup", &rbac.GroupController{}, "*:DelGroup")
	beego.Router("/rbac/group/index", &rbac.GroupController{}, "*:Index")

	beego.Router("/rbac/role/AddAndEdit", &rbac.RoleController{}, "*:AddAndEdit")
	beego.Router("/rbac/role/DelRole", &rbac.RoleController{}, "*:DelRole")
	beego.Router("/rbac/role/AccessToNode", &rbac.RoleController{}, "*:AccessToNode")
	beego.Router("/rbac/role/AddAccess", &rbac.RoleController{}, "*:AddAccess")
	beego.Router("/rbac/role/RoleToUserList", &rbac.RoleController{}, "*:RoleToUserList")
	beego.Router("/rbac/role/AddRoleToUser", &rbac.RoleController{}, "*:AddRoleToUser")
	beego.Router("/rbac/role/Getlist", &rbac.RoleController{}, "*:Getlist")
	beego.Router("/rbac/role/index", &rbac.RoleController{}, "*:Index")

	fmt.Println("Start ok")
	//判断初始化参数
	initArgs()
	beego.AddFuncMap("stringsToJson", lib.StringsToJson)
	beego.Run()

}

func initArgs() {
	args := os.Args
	for _, v := range args {
		if v == "-syncdb" {
			rbacmodels.Syncdb()
		}
	}
}
