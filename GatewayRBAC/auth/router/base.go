package router

import (
	"auth/controller"
	"github.com/TechMaster/core/rbac"

	"github.com/kataras/iris/v12"
)

func RegisterRoute(app *iris.Application) {
	app.Get("/", controller.ShowHomePage)
	app.Post("/login", controller.Login)
	rbac.Get(app, "/secret", rbac.AllowAll(), controller.ShowSecret)
	rbac.Get(app, "/logout", rbac.AllowAll(), controller.LogoutFromWeb)

	api := app.Party("/api")
	{
		api.Post("/login", controller.LoginREST)
		api.Get("/logout", controller.LogoutREST)
	}
}
