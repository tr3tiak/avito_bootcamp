package entity

type User struct {
	Name     string
	Password string
	Role     string
}

type UserUsecase interface {
	Register(*User) error
	Login(*User) (string, error)
}

type UserRepo interface {
	Create(*User) error
	Get(string) (*User, error)
}
