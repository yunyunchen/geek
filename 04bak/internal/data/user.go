package data

import (
	"context"
	"geek/04/internal/biz"
)

type User struct {
	ID       uint `gorm:"primary_key"`
	Username string
	Password string
}

type UserRepo struct {
	//data *Data
}

func (ar *UserRepo) GetArticle(ctx context.Context, id int64) (*biz.User, error) {
	/*p, err := ar.data.db.Article.Get(ctx, id)
	if err != nil {
		return nil, err
	}*/
	return &biz.User{
		Id:        1,
		Name:     "p.Title",
		Email:   "p.Content",
	}, nil
}