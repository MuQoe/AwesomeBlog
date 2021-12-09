package model

import (
	global "AwesomeBlog/globals"
	"AwesomeBlog/tools"
	"errors"
	"time"

	"go.uber.org/zap"

	"github.com/golang-jwt/jwt"

	"gorm.io/gorm"
)

type UserModel struct {
	ID        int64          `json:"user_id" gorm:"primarykey"`                        // 用户ID
	Identity  int8           `json:"user_Identity"`                                    // 用户权限组
	UserName  string         `json:"username" gorm:"type:varchar(100);unique"`         // 用户名
	Password  string         `json:"user_password"`                                    // 用户密码
	Avatar    string         `json:"user_avatar"`                                      // 用户头像地址
	Email     string         `json:"user_email" gorm:"type:varchar(100);unique_index"` // 用户Email 唯一值
	Sign      string         `json:"user_sign"`                                        // 用户个性签名
	Status    int8           `json:"user_status"`                                      // 用户状态(0未激活 1已激活 2已封禁)
	CreatedAt time.Time      `json:"user_createat"`                                    // 用户创建时间
	DeletedAt gorm.DeletedAt `json:"user_deleteat" gorm:"index"`                       // 用户删除时间
	Agent     string         `json:"agent"`                                            // 浏览器UA
	Ip        string         `json:"ip"`                                               // 登录ip地址
	Token     string         `json:"token"`                                            // 用户的token数据
	LoginTime time.Time      `json:"login_time"`                                       // 用户登录的时间                               //登录信息（用于身份验证）
}

/*
// MapToModel get the user model from given map.
func (u UserModel) MapToModel(m map[string]interface{}) UserModel {
	u.ID, _ = m["id"].(int64)
	u.Identity, _ = m["identity"].(int8)
	u.UserName, _ = m["user_name"].(string)
	u.Password, _ = m["password"].(string)
	u.Avatar, _ = m["avatar"].(string)
	u.Email, _ = m["email"].(string)
	u.Sign, _ = m["sign"].(string)
	u.Status, _ = m["status"].(int8)
	u.Email, _ = m["email"].(string)
	u.CreatedAt, _ = m["created_at"].(time.Time)
	u.DeletedAt, _ = m["deleted_at"].(gorm.DeletedAt)
	u.Agent, _ = m["agent"].(string)
	u.Ip, _ = m["ip"].(string)
	u.Token, _ = m["token"].(string)
	u.LoginTime, _ = m["login_time"].(time.Time)
	return u
}
*/

// User return a default user model.
func User() UserModel {
	return UserModel{}
}

func (u UserModel) FindByID(ID interface{}) (UserModel, error) {
	var find_user UserModel
	err := global.DBEngine.First(&find_user, "id = ?", ID).Error
	return find_user, err
}

func (u UserModel) FindByUserName(UserName interface{}) (UserModel, error) {
	var find_user UserModel
	err := global.DBEngine.First(&find_user, "user_name = ?", UserName).Error
	return find_user, err
}

func (u UserModel) CreateUser(username, password, Repassword, email string) error {

	if password != Repassword {
		return errors.New("Password and Repassword not same")
	}
	// 生成随机的Token
	// user.Token = tools.GetRandomNum(20)
	// 密码加密
	u.Password = tools.Encrypt(password)
	zap.S().Debugf("Original Password %s Encrypt Password %s", password, u.Password)

	u.UserName = username
	u.Email = email
	u.Identity = NORMAL_USER

	return global.DBEngine.Create(&u).Error
}

func (u UserModel) Update(values interface{}) error {
	return global.DBEngine.Model(&u).Where("id = ?", u.ID).Updates(values).Error
}

func (u UserModel) Delete() error {
	return global.DBEngine.Model(&u).Where("id = ?", u.ID).Delete(&u).Error
}

func (u UserModel) Save() error {
	return global.DBEngine.Model(&u).Save(&u).Error
}

func (u UserModel) IsAdmin() bool {
	return u.Identity == ADMINISTRATOR
}

func (u UserModel) IsEmpty() bool {
	return u.ID == 0
}

// JwtCreateToken jwt生成token
func (u UserModel) CreateToken() (string, error) {
	// 创建token
	token := jwt.New(jwt.SigningMethodHS256)

	//设置属性
	claims := token.Claims.(jwt.MapClaims)
	claims["ID"] = u.ID
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()
	claims["admin"] = u.IsAdmin()

	//生成token字符串
	t, err := token.SignedString([]byte(global.AppSetting.TokenSecret))
	if err != nil {
		return "", err
	}
	return t, nil
}

// UserRegisterParam 用户注册
type UserRegisterParam struct {
	Username   string `json:"username" validate:"required"`    // 用户名
	Password   string `json:"password" validate:"required"`    // 密码
	RePassword string `json:"re_password" validate:"required"` // 重复密码
	Email      string `json:"email" validate:"required"`       // 邮箱
}

// UserLoginParam 用户登陆
type UserLoginParam struct {
	Username string `json:"username" validate:"required"` // 用户名
	Password string `json:"password" validate:"required"` // 密码
}
