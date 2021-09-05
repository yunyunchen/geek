//+build wireinject

package main

import (
	"geek/04/configs"
	"geek/04/internal/biz"
	"geek/04/internal/data"
	"geek/04/internal/server"
	"geek/04/internal/service"
	"github.com/google/wire"
)

func initApp(config *configs.Config) (*userApp,error) {
	//panic(wire.Build(data.NewUserRepo, biz.NewUserBiz, service.NewUserService))
	panic(wire.Build(server.Provider,service.Provider,biz.Provider,data.Provider,newUserApp))
}

