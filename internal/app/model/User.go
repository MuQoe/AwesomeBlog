package model

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID            uint           `json:"user_id" gorm:"primarykey"`                        // 用户ID
	Identity      int8           `json:"user_Identity"`                                    // 用户权限组
	UserName      string         `json:"username" gorm:"type:varchar(100);unique"`         // 用户名
	Password      string         `json:"user_password"`                                    // 用户密码
	Avatar        string         `json:"user_avatar"`                                      // 用户头像地址
	Email         string         `json:"user_email" gorm:"type:varchar(100);unique_index"` // 用户Email 唯一值
	Sign          string         `json:"user_sign"`                                        // 用户个性签名
	Status        int8           `json:"user_status"`                                      // 用户状态(0未激活 1已激活 2已封禁)
	LastLoginTime time.Time      `json:"user_lastlogintime"`                               // 上次登陆时间
	CreatedAt     time.Time      `json:"user_createat"`                                    // 用户创建时间
	DeletedAt     gorm.DeletedAt `json:"user_deleteat" gorm:"index"`                       // 用户删除时间
	Agent         string         `json:"agent"`                                            // 浏览器UA
	Ip            string         `json:"ip"`                                               // 登录ip地址
	Token         string         `json:"token"`                                            // 用户的token数据
	LoginTime     time.Time      `json:"login_time"`                                       // 用户登录的时间                               //登录信息（用于身份验证）
}

func (u User) GetUser(db *gorm.DB, userID uint) (*User, error) {
	var find_user User
	err := db.Find(&find_user, "id = ?", userID).Error
	return &find_user, err
}

func (u User) CreateUser(db *gorm.DB) error {
	return db.Create(&u).Error
}

func (u User) UpdateUser(db *gorm.DB, values interface{}) error {
	return db.Model(&u).Where("id = ?", u.ID).Updates(values).Error
}

func (u User) DeleteUser(db *gorm.DB) error {
	return db.Where("id = ?", u.ID).Delete(&u).Error
}
