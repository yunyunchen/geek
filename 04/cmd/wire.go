//+build wireinject
package main

import (
	"geek/04/internal/biz"
	"geek/04/internal/data"
	"geek/04/internal/service"
	"github.com/google/wire"
)

func InitializeAllInstance() *service.UserService {
	//wire.Build(v1.NewUserInterface,UserSet)
	panic(wire.Build(service.NewUserService,
		biz.NewUserBiz,
		data.NewUserRepo,
	))
}

/*var UserSet = wire.NewSet(
	service.NewUserService,
	biz.NewUserBiz,
	data.NewUserRepo,
)
*/
