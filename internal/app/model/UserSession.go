package model

import global "AwesomeBlog/globals"

// UserLogin 用户的登录数据
type UserSessionModel struct {
	UserID int64  `json:"user_id" gorm:"primarykey"` //用户id
	Token  string `json:"token"`                     //用户token数据
}

func UserSession() UserSessionModel {
	return UserSessionModel{}
}

func (u UserSessionModel) FindByUserID(userID int64) (UserSessionModel, error) {
	var founded UserSessionModel
	err := global.DBEngine.First(&founded, "user_id = ?", userID).Error
	return founded, err
}

func (u UserSessionModel) FindByToken(Token string) (UserSessionModel, error) {
	var founded UserSessionModel
	err := global.DBEngine.First(&founded, "token = ?", Token).Error
	return founded, err
}

func (u UserSessionModel) Save() error {
	return global.DBEngine.Save(&u).Error
}
