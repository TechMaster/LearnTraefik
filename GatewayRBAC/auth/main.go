package main

import (
	"auth/router"

	"github.com/TechMaster/core/config"
	"github.com/TechMaster/core/rbac"
	"github.com/TechMaster/core/session"
	"github.com/TechMaster/core/template"
	"github.com/spf13/viper"

	"github.com/TechMaster/logger"
	"github.com/iris-contrib/middleware/cors"
	"github.com/kataras/iris/v12"
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

	crs := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowCredentials: true,
	})
	app.UseRouter(crs)

	app.Use(session.Sess.Handler())

	rbacConfig := rbac.NewConfig()
	rbacConfig.RootAllow = true
	rbacConfig.MakeUnassignedRoutePublic = true
	rbac.Init(rbacConfig) //Khởi động với cấu hình mặc định
	//đặt hàm này trên các hàm đăng ký route - controller
	app.Use(rbac.CheckRoutePermission)
	router.RegisterRoute(app)

	template.InitViewEngine(app)

	//Luôn để hàm này sau tất cả lệnh cấu hình đường dẫn với RBAC
	rbac.BuildPublicRoute(app)
	_ = app.Listen(viper.GetString("port"))
}
