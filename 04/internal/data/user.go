package data

import (
	"fmt"
	"geek/04/internal/biz"
)

//var ProviderSet = wire.NewSet(NewUserRepo)

//type UserRepo interface {
//	InsertUser(*UserPo) error
//	//data *Data
//}

type userRepo struct{}

func NewUserRepo(repo userRepo) biz.UserRepo {
	return nil
}

//
//func (ar *userRepo) GetUser(id int64) (*UserPo, error) {
//	// todo
//	return &UserPo{},nil
//}
// (ur userRepo)
func InsertUser(po *biz.User) error {
	//todo
	fmt.Println(po.Username + " / " + po.Email + " 持久化。")
	return nil
}
