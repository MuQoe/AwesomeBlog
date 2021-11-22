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
	data := gin.H{
		"message": "index api",
	}
	Response.GlobalResponse.ResponseOk(ctx, data)
	return
}
