package tools

import (
	"AwesomeBlog/pkg/Response"

	"go.uber.org/zap"

	"github.com/go-playground/validator/v10"

	"github.com/gin-gonic/gin"
)

// 实例化验证对象
var validate = validator.New()

// ValidatorParam 校验参数是否正确(true表示参数非法)
func ValidatorParam(c *gin.Context, param interface{}) bool {
	// 先验证基本的数据接口是否匹配
	if c.Bind(param) != nil && c.BindQuery(param) != nil {
		Response.GlobalResponse.ResponseBadRequest(c)
		return false
	} else if err := validate.Struct(param); err != nil { // 再验证格式是否正确
		zap.S().Info(err)
		Response.GlobalResponse.ResponseBadRequest(c, "参数不符合格式")
		return true
	} else { // 验证通过后返回空
		return false
	}
}

// JudgeParams 判断用户的参数是否都输入了
func JudgeParams(params ...string) bool {
	//遍历参数
	for _, v := range params {
		if v == "" {
			return true
		}
	}
	return false
}
