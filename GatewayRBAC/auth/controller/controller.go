package controller

import (
	"auth/repo"
	"fmt"

	"github.com/TechMaster/core/logger"
	"github.com/TechMaster/core/pass"
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
	if authinfo := session.GetAuthInfo(ctx); authinfo != nil {
		ctx.ViewData("roles", rbac.RolesNames(authinfo.Roles))
	}

	_ = ctx.View("index")
}

func ShowSecret(ctx iris.Context) {
	logger.Info(ctx, "Đây là trang tuyệt mật")
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
		_, _ = ctx.WriteString("User Not Found")
		return
	}

	if !pass.CheckPassword(loginReq.Pass, user.Password, "") {
		_, _ = ctx.WriteString("Wrong password")
		return
	}

	session.SetAuthenticated(ctx, pmodel.AuthenInfo{
		Id:       user.Id,
		FullName: user.FullName,
		Email:    user.Email,
		Roles:    pmodel.IntArrToRoles(user.Roles), //Chuyển từ mảng []int sang map[int]bool
	})

	//Login thành công thì quay về trang chủ
	ctx.Redirect("/")
}

func LogoutFromWeb(ctx iris.Context) {
	session.Logout(ctx)
	ctx.Redirect("/")
}
