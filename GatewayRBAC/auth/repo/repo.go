package repo

import (
	"errors"
	"strings"

	"github.com/TechMaster/core/pass"
	"github.com/TechMaster/core/pmodel"
	"github.com/TechMaster/core/rbac"
	"github.com/segmentio/ksuid"
)

var users = make(map[string]*pmodel.User)

func init() {
	CreateNewUser("Bùi Văn Hiên", "1", "hien@gmail.com", "0123456789", rbac.TRAINER, rbac.MAINTAINER)
	CreateNewUser("Nguyễn Hàn Duy", "1", "duy@gmail.com", "0123456786", rbac.TRAINER, rbac.STUDENT)
	CreateNewUser("Phạm Thị Mẫn", "1", "man@gmail.com", "0123456780", rbac.SALE, rbac.STUDENT)
	CreateNewUser("Trịnh Minh Cường", "1", "cuong@gmail.com", "0123456000", rbac.ADMIN, rbac.TRAINER)
	CreateNewUser("Nguyễn Thành Long", "1", "long@gmail.com", "0123456001", rbac.STUDENT)
}

func CreateNewUser(fullName string, password string, email string, phone string, roles ...int) {
	hassedpass, _ := pass.HashBcryptPass(password)

	user := pmodel.User{
		Id:       ksuid.New().String(),
		FullName: fullName,
		Password: hassedpass,
		Email:    strings.ToLower(email),
		Phone:    phone,
		Roles:    roles,
	}

	users[user.Email] = &user //Thêm user vào users
}
func QueryByEmail(email string) (user *pmodel.User, err error) {
	user = users[strings.ToLower(email)]
	if user == nil {
		return nil, errors.New("User not found")
	} else {
		return user, nil
	}
}
