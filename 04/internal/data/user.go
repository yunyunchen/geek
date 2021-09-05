package data

import (
	"geek/04/internal/biz"
	"gorm.io/gorm"
	"time"
)

type userPo struct {
	Id        int64 `gorm:"primary_key"`
	Username  string
	Email     string
	CreatedAt *time.Time
	DeletedAt *time.Time
	UpdatedAt *time.Time
}


type userRepo struct{
	db *gorm.DB
}

func NewUserRepo(db *gorm.DB) biz.UserRepo {
	return &userRepo{
		db: db,
	}
}

func (up *userRepo)InsertUser(po *biz.User) error {
	result := up.db.First(&po, po.Name)
	if result.Error != nil {
		return nil
	}

	// PO -> DO
	return  nil
}
