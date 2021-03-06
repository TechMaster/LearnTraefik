package router

import (
	"mainsite/controller"

	"github.com/TechMaster/core/rbac"

	"github.com/kataras/iris/v12"
)

func RegisterRoute(app *iris.Application) {

	app.Get("/", controller.ShowHomePage) //Không dùng rbac có nghĩa là public method
	app.Post("/login", controller.Login)
	rbac.Get(app, "/logout", rbac.AllowAll(), controller.LogoutFromWeb)

	blog := app.Party("/blog")
	{
		blog.Get("/", controller.GetAllPosts) //Không dùng rbac có nghĩa là public method
		rbac.Get(blog, "/all", rbac.AllowAll(), controller.GetAllPosts)
		rbac.Get(blog, "/create", rbac.Forbid(rbac.MAINTAINER), controller.GetAllPosts)
		rbac.Get(blog, "/{id:int}", rbac.Allow(rbac.AUTHOR, rbac.EDITOR), controller.GetPostByID)
		rbac.Get(blog, "/delete/{id:int}", rbac.Allow(rbac.ADMIN, rbac.AUTHOR, rbac.EDITOR), controller.DeletePostByID)
		rbac.Any(blog, "/any", rbac.Allow(rbac.MAINTAINER), controller.PostMiddleware)
	}

	student := app.Party("/student")
	{
		rbac.Get(student, "/submithomework", rbac.Allow(rbac.STUDENT), controller.SubmitHomework)
	}

	trainer := app.Party("/trainer")
	{
		rbac.Get(trainer, "/createlesson", rbac.Allow(rbac.TRAINER), controller.CreateLesson)
	}

	sysop := app.Party("/sysop")
	{
		rbac.Get(sysop, "/backupdb", rbac.Allow(rbac.MAINTAINER), controller.BackupDB)
		rbac.Get(sysop, "/upload", rbac.Allow(rbac.MAINTAINER), controller.ShowUploadForm)
		rbac.Get(sysop, "/err", rbac.AllowAll(), controller.ShowErr)
		rbac.Post(sysop, "/upload", rbac.Allow(rbac.MAINTAINER, rbac.SALE), iris.LimitRequestBodySize(300000), controller.UploadPhoto)
	}

	sales := app.Party("/sale")
	{
		rbac.Get(sales, "/runads", rbac.Allow(rbac.SALE), controller.RunAdvertise)
	}

}
