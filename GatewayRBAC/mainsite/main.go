package main

import (
	"mainsite/router"

	"github.com/TechMaster/core/config"
	"github.com/TechMaster/core/logger"
	"github.com/TechMaster/core/rbac"
	"github.com/TechMaster/core/session"
	"github.com/TechMaster/core/template"

	"github.com/kataras/iris/v12"
	"github.com/spf13/viper"
)

func main() {
	app := iris.New()
	config.ReadConfig()

	logFile := logger.Init() //Cần phải có 2 file error.html và info.html ở /views
	if logFile != nil {
		defer logFile.Close()
	}

	redisDb := session.InitRedisSession()
	defer redisDb.Close()
	app.Use(session.Sess.Handler())

	rbacConfig := rbac.NewConfig()
	rbacConfig.MakeUnassignedRoutePublic = true
	rbac.Init(rbacConfig) //Khởi động với cấu hình mặc định

	//đặt hàm này trên các hàm đăng ký route - controller
	app.Use(rbac.CheckRoutePermission)

	app.HandleDir("/", iris.Dir("./static"))

	router.RegisterRoute(app)

	template.InitViewEngine(app)

	//Luôn để hàm này sau tất cả lệnh cấu hình đường dẫn với RBAC
	rbac.BuildPublicRoute(app)
	//rbac.DebugRouteRole()
	_ = app.Listen(viper.GetString("port"))
}
