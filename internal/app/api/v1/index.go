package v1

import (
	"AwesomeBlog/pkg/Response"
	"github.com/gin-gonic/gin"
)

type Index struct{}

func NewIndex() Index {
	return Index{}
}

func (i Index) Get(ctx *gin.Context) {
	rsp := Response.NewResponse(ctx)
	data := gin.H{
		"message": "index api",
	}
	rsp.ToResponse(data)
	return
}
