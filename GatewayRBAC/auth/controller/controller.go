package controller

import (
	"auth/repo"
	"fmt"

	"github.com/TechMaster/core/pmodel"
	"github.com/TechMaster/core/rbac"
	"github.com/TechMaster/core/session"

	"github.com/kataras/iris/v12"
)

/*
	Lưu thông tin đăng nhập từ client gửi lên
*/
type LoginRequest struct {
	Email string
	Pass  string
}

func ShowHomePage(ctx iris.Context) {
	if raw_authinfo := ctx.GetViewData()[session.AUTHINFO]; raw_authinfo != nil {
		authinfo := raw_authinfo.(*pmodel.AuthenInfo)
		ctx.ViewData("roles", rbac.RolesNames(authinfo.Roles))
	}
	_ = ctx.View("index")
}

func ShowSecret(ctx iris.Context) {
	// Check if user is authenticated
	if !session.IsLogin(ctx) {
		ctx.StatusCode(iris.StatusForbidden)
		return
	}
	_, _ = ctx.WriteString("Secret Page")
}

/*
Login thông qua form. Dành cho ứng dụng web server side renderings
*/
func Login(ctx iris.Context) {
	var loginReq LoginRequest

	if err := ctx.ReadForm(&loginReq); err != nil {
		fmt.Println(err.Error())
		return
	}

	user, err := repo.QueryByEmail(loginReq.Email)
	if err != nil { //Không tìm thấy user
		_, _ = ctx.WriteString("Login Failed")
		return
	}

	if user.Pass != loginReq.Pass {
		_, _ = ctx.WriteString("Wrong password")
		return
	}

	session.SetAuthenticated(ctx, pmodel.AuthenInfo{
		User:  user.User,
		Email: user.Email,
		Roles: user.Roles,
	})

	//Login thành công thì quay về trang chủ
	ctx.Redirect("/")
}

func LogoutFromWeb(ctx iris.Context) {
	logout(ctx)
	ctx.Redirect("/")
}

func logout(ctx iris.Context) {
	/*	if !session.IsLogin(ctx) {
		logger.Log(ctx, eris.Warning("Bạn chưa login").UnAuthorized())
	}*/
	//Xoá toàn bộ session và xoá luôn cả Cookie sessionid ở máy người dùng
	session.Sess.Destroy(ctx)
}
