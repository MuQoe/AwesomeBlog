package v1

import (
	"AwesomeBlog/internal/app/model"
	"AwesomeBlog/internal/app/model/Auth"
	"AwesomeBlog/pkg/Response"
	"AwesomeBlog/tools"
	"errors"

	"gorm.io/gorm"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type UserAPI struct {
}

func NewUserAPI() UserAPI {
	return UserAPI{}
}
func (u UserAPI) Login(c *gin.Context) {
	// 参数验证
	param := new(model.UserLoginParam)
	if tools.ValidatorParam(c, param) {
		zap.S().Debug("Param Invalid")
		Response.GlobalResponse.ResponseBadRequest(c, "Param Invalid")
		return
	}
	user, ok := Auth.Check(param.Password, param.Username)

	if !ok {
		Response.GlobalResponse.ResponseUnauthorized(c, "用户名或密码错误")
		return
	}

	Token, err := user.CreateToken()
	if err != nil {
		zap.S().Error(err)
		Response.GlobalResponse.ResponseServerError(c, "user.CreateToken() 失败")
		return
	}

	UserSession := model.UserSession()

	UserSession.UserID = user.ID
	UserSession.Token = Token

	err = UserSession.Save()
	if err != nil {
		zap.S().Error(err)
		Response.GlobalResponse.ResponseServerError(c, "UserSession 保存失败")
		return
	}
	err = user.Update(map[string]interface{}{"token": Token})
	if err != nil {
		zap.S().Error(err)
		Response.GlobalResponse.ResponseServerError(c, "User Token 保存失败")
		return
	}
	Response.GlobalResponse.ResponseCreated(c, UserSession)

	// c.SetCookie("token", Token, 60*60*24, "/", global.ServerSetting.Domain, false, false)
	return

	// TODO add User activate check

}

func (u UserAPI) Create(c *gin.Context) {
	// 参数验证
	param := new(model.UserRegisterParam)
	if tools.ValidatorParam(c, param) {
		zap.S().Debug("Param Invalid")
		Response.GlobalResponse.ResponseBadRequest(c, "Param Invalid")
		return
	}

	_, err := model.User().FindByUserName(param.Username)
	if err != nil {
		// 确认数据库没有此用户
		if errors.Is(err, gorm.ErrRecordNotFound) {
			err = model.User().CreateUser(param.Password, param.Password, param.RePassword, param.Email)
			if err != nil {
				zap.S().Errorf("svc.CreateUser err: %v", err)
				Response.GlobalResponse.ResponseBadRequest(c, "svc.CreateUser Failed")
				return
			}
			Response.GlobalResponse.ResponseCreated(c, gin.H{"msg": "Create User Success"})
			return
		}
	}
	Response.GlobalResponse.ResponseBadRequest(c, "User Name Already Exist")
}
