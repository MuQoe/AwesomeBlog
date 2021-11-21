package Response

import (
	"AwesomeBlog/pkg/errcode"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
)

type Response struct {
	Ctx *gin.Context
}

func NewResponse(ctx *gin.Context) *Response {
	return &Response{Ctx: ctx}
}

func (r *Response) ToResponse(data interface{}) {
	if data == nil {
		data = gin.H{}
	}
	r.Ctx.JSON(http.StatusOK, data)
}

func (r *Response) ToErrorResponse(err *errcode.Error) {
	zap.S().Errorf("code: %d, msg: %s", err.Code, err.Msg)
	response := gin.H{"code": err.Code, "msg": err.Msg}
	details := err.Detail
	if len(details) > 0 {
		response["details"] = details
	}
	r.Ctx.JSON(err.StatusCode(), response)
}
