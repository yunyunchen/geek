//+build wireinject
package main

import (
	"geek/04/internal/biz"
	"geek/04/internal/data"
	"geek/04/internal/service"
	"github.com/google/wire"
)

func InitializeAllInstance() service.UserService {
	panic(wire.Build(data.NewUserRepo, biz.NewUserBiz, service.NewUserService))
	return service.UserService{}
}

/*var UserSet = wire.NewSet(
	data.NewUserRepo,
	biz.NewUserBiz,
	service.NewUserService,
)
*/
