package biz

import (
	"time"
)

/*type User struct {
	Id    int64
	Name  string
	Email string
}*/

//var ProviderSet = wire.NewSet(NewUserBiz)

type User struct {
	Id        int64 `gorm:"primary_key"`
	Username  string
	Email     string
	CreatedAt *time.Time
	DeletedAt *time.Time
	UpdatedAt *time.Time
}

type UserRepo interface {
	InsertUser(user *User) error
}

type UserUsecase struct {
	repo UserRepo
}

func NewUserBiz(repo UserRepo) *UserUsecase {
	return &UserUsecase{repo: repo}
}

//func (uc *UserUsecase) Get(ctx context.Context, id int64) (p *User, err error) {
//	p, err = uc.repo.GetUser(ctx, id)
//	if err != nil {
//		return
//	}
//	return
//}
//
func (uu UserUsecase) InsertUserBiz(do *User) error {
	/*now := time.Now()
	po := data.UserPo{
		//Id: do.Id,
		Username:  do.Name,
		Email:     do.Email,
		CreatedAt: &now,
		DeletedAt: nil,
		UpdatedAt: &now,
	}*/
	//return data.InsertUser(&po)
	return uu.repo.InsertUser(do)
}
