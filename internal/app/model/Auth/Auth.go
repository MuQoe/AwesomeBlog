package Auth

import (
	"AwesomeBlog/internal/app/model"
	"AwesomeBlog/tools"
	"errors"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

// Check check the password and username and return the user model.
func Check(password string, username string) (user model.UserModel, ok bool) {

	user, err := model.User().FindByUserName(username)

	if err != nil {
		// 确认数据库没有此用户
		ok = false
		if errors.Is(err, gorm.ErrRecordNotFound) {
			zap.S().Debug("User Name not Exist")
			return
		}
		zap.S().Warn(err)
		return
	}

	if user.IsEmpty() {
		ok = false
	} else {
		if comparePassword(password, user.Password) {
			ok = true
			// TODO: 增加用户权限
			// user = user.WithRoles().WithPermissions().WithMenus()
			// user.UpdatePwd(EncodePassword([]byte(password)))
		} else {
			ok = false
		}
	}
	return
}

func comparePassword(comPwd, pwdHash string) bool {
	encpwd := tools.Encrypt(comPwd)
	return encpwd == pwdHash
}
