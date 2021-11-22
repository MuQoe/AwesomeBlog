package v1

import (
	"AwesomeBlog/internal/app/service"
	"AwesomeBlog/pkg/Response"
	"AwesomeBlog/tools"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type UserAPI struct {
}

func NewUserAPI() UserAPI {
	return UserAPI{}
}

func (u UserAPI) Create(c *gin.Context) {
	// 参数验证
	param := new(service.UserRegisterParam)
	if tools.ValidatorParam(c, param) {
		zap.S().Debug("Param Invalid")
		Response.GlobalResponse.ResponseBadRequest(c, "Param Invalid")
		return
	}
	svc := service.New(c.Request.Context())
	err := svc.CreateUser(param)
	if err != nil {
		zap.S().Errorf("svc.CreateUser err: %v", err)
		Response.GlobalResponse.ResponseBadRequest(c, "svc.CreateUser Failed")
		return
	}
	Response.GlobalResponse.ResponseOk(c, gin.H{"msg": "Create User Success"})
}
