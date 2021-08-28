package controller

import (
	"github.com/TechMaster/core/session"

	"github.com/TechMaster/core/rbac"

	"github.com/kataras/iris/v12"
)

func ShowHomePage(ctx iris.Context) {
	if authinfo := session.GetAuthInfo(ctx); authinfo != nil {
		ctx.ViewData("roles", rbac.RolesNames(authinfo.Roles))
	}

	_ = ctx.View("index")
}
