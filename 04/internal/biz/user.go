package biz

import (
	"geek/04/internal/data"
	"time"
)

type UserDo struct {
	Id    int64
	Name  string
	Email string
}
type UserBiz interface {
	//QueryById(id string) (data.UserPo, error)
	Insert(*data.UserPo) error
	// todo
}

type userBiz struct {
	userRepo data.UserRepo
}

func NewUserBiz(userRepo data.UserRepo) userBiz {
	return userBiz{userRepo: userRepo}
}

//func (uc *UserUsecase) Get(ctx context.Context, id int64) (p *User, err error) {
//	p, err = uc.repo.GetUser(ctx, id)
//	if err != nil {
//		return
//	}
//	return
//}
//
func Insert(do *UserDo) error {
	now := time.Now()
	po := data.UserPo{
		//Id: do.Id,
		Username:  do.Name,
		Email:     do.Email,
		CreatedAt: &now,
		DeletedAt: nil,
		UpdatedAt: &now,
	}
	return data.Insert(&po)
}
