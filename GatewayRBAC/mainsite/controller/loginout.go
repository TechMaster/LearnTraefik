package controller

import (
	"mainsite/pmodel"
	"mainsite/resto"
	"mainsite/session"
	"net/http"

	json "github.com/goccy/go-json"
	"github.com/spf13/viper"

	"github.com/TechMaster/eris"
	"github.com/TechMaster/logger"
	"github.com/kataras/iris/v12"
)

/*
Login thông qua form. Dành cho ứng dụng web server side renderings
*/
func Login(ctx iris.Context) {
	type LoginRequest struct {
		Email string
		Pass  string
	}
	var loginReq LoginRequest

	if err := ctx.ReadForm(&loginReq); err != nil {
		logger.Log(ctx, eris.NewFrom(err).SetType(eris.WARNING).BadRequest())
		return
	}

	response, err := resto.Post(viper.GetString("authservice.host")+"/api/login", loginReq)
	if err != nil {
		logger.Log(ctx, eris.NewFromMsg(err, "Lỗi khi gọi Auth service").InternalServerError())
		return
	}

	if response.StatusCode != http.StatusOK { //Đăng nhập lỗi
		var res struct {
			Error string `json:"error"`
		}
		_ = json.NewDecoder(response.Body).Decode(&res)
		logger.Log(ctx, eris.Warning(res.Error).UnAuthorized())
		return
	}

	//Đăng nhập thành công thì nhận AuthenInfo trả về từ auth service
	var authInfo pmodel.AuthenInfo
	if err := json.NewDecoder(response.Body).Decode(&authInfo); err != nil {
		logger.Log(ctx, eris.NewFromMsg(err, "Lỗi phân tích kết quả đăng nhập").InternalServerError())
		return
	}

	//Phải lưu authentication ở đây chứ không phải ở auth service !
	session.SetAuthenticated(ctx, authInfo)

	//Login thành công thì quay về trang chủ
	ctx.Redirect("/")
}

func LogoutFromWeb(ctx iris.Context) {
	logout(ctx)
	ctx.Redirect("/")
}

func logout(ctx iris.Context) {
	session.Sess.Destroy(ctx)
}
