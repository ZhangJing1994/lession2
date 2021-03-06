// biz/user.go

package biz

type User struct {
	ID        int32
	Name      string
	Age       int32
	Telephone int16
}

type UserRepo interface {
	Save(*User) int32
}

func NewUserUsecase(repo UserRepo) *UserUsecase {
	return &UserUsecase{repo: repo}
}

type UserUsecase struct {
	repo UserRepo
}

func (s *UserUsecase) SaveUser(u *User) {
	id := s.repo.Save(u)
	u.ID = id
}
