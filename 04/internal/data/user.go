package data

import (
	"fmt"
	"time"
)

type UserPo struct {
	Id        int64 `gorm:"primary_key"`
	Username  string
	Email     string
	CreatedAt *time.Time
	DeletedAt *time.Time
	UpdatedAt *time.Time
}

type UserRepo interface {
	Insert(*UserPo) error
	//data *Data
}

type userRepo struct{}

func NewUserRepo(repo UserRepo) userRepo {
	return userRepo{}
}

//
//func (ar *userRepo) GetUser(id int64) (*UserPo, error) {
//	// todo
//	return &UserPo{},nil
//}
// (ur userRepo)
func Insert(po *UserPo) error {
	//todo
	fmt.Println(po.Username + " / " + po.Email + " 持久化。")
	return nil
}
