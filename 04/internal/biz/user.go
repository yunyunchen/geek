package biz

type User struct {
	Id    int64
	Name  string
	Email string
}

//var ProviderSet = wire.NewSet(NewUserBiz)



type UserRepo interface {
	InsertUser(user *User) error
}

type UserUsecase struct {
	repo UserRepo
}

func NewUserUsecase(repo UserRepo) *UserUsecase {
	return &UserUsecase{repo: repo}
}

func (uu *UserUsecase) InsertUserBiz(do *User) error {
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
