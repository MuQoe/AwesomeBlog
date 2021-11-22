package service

import (
	global "AwesomeBlog/globals"
	"AwesomeBlog/internal/app/dao"
	"context"
)

type Service struct {
	ctx context.Context
	dao *dao.Dao
}

func New(ctx context.Context) Service {
	service := Service{ctx: ctx}
	service.dao = dao.New(global.DBEngine)
	return service
}
