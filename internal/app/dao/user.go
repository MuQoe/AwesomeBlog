package dao

import (
	"AwesomeBlog/internal/app/model"
	"AwesomeBlog/tools"

	"go.uber.org/zap"
)

func (d *Dao) CountUser() {

}
func (d *Dao) GetUser(UserId uint) {

}

func (d *Dao) CreateUser(username, password, email string) error {
	var user model.User
	// 生成随机的Token
	user.Token = tools.GetRandomNum(20)
	// 密码加密
	user.Password = tools.Encrypt(password)
	zap.S().Debugf("Original Password %s\nEncrypt Password %s\n", password, user.Password)

	user.UserName = username
	user.Email = email
	user.Identity = model.NORMAL_USER

	//TODO: add default attribute
	return user.CreateUser(d.engine)
}
