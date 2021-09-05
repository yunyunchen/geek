package biz

import (
	"context"
)

type User struct {
	Id    int64
	Name  string
	Email string
}

type UserRepo interface {
	// db
	GetUser(ctx context.Context, id int64) (*User, error)
	UpdateUser(ctx context.Context, id int64, user *User) error
}

type UserUsecase struct {
	repo UserRepo
}

/*func NewArticleUsecase(repo UserRepo, logger log.Logger) *UserUsecase {
	return &UserUsecase{repo: repo}
}*/

func (uc *UserUsecase) Get(ctx context.Context, id int64) (p *User, err error) {
	p, err = uc.repo.GetUser(ctx, id)
	if err != nil {
		return
	}
	return
}

func (uc *UserUsecase) Update(ctx context.Context, id int64, user *User) error {
	return uc.repo.UpdateUser(ctx, id, user)
}


