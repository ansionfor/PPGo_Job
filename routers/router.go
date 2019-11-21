package routers

import (
	"github.com/astaxie/beego"
	"github.com/george518/PPGo_Job/controllers"
)

func init() {
	// 默认登录
	siteDir := beego.AppConfig.String("site.dir")
	beego.Router(siteDir + "/", &controllers.LoginController{}, "*:Login")
	beego.Router(siteDir + "/login_in", &controllers.LoginController{}, "*:LoginIn")
	beego.Router(siteDir + "/login_out", &controllers.LoginController{}, "*:LoginOut")
	beego.Router(siteDir + "/help", &controllers.HomeController{}, "*:Help")
	beego.Router(siteDir + "/home", &controllers.HomeController{}, "*:Index")
	beego.Router(siteDir + "/home/start", &controllers.HomeController{}, "*:Start")

	beego.AutoRouter(&controllers.TaskController{})
	beego.AutoRouter(&controllers.GroupController{})
	beego.AutoRouter(&controllers.TaskLogController{})
	//资源分组管理
	beego.AutoRouter(&controllers.ServerGroupController{})
	beego.AutoRouter(&controllers.ServerController{})
	beego.AutoRouter(&controllers.BanController{})

	//权限用户相关
	beego.AutoRouter(&controllers.AuthController{})
	beego.AutoRouter(&controllers.RoleController{})
	beego.AutoRouter(&controllers.AdminController{})
	beego.AutoRouter(&controllers.UserController{})

	beego.AutoRouter(&controllers.NotifyTplController{})
}
