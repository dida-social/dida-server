package user

type Service interface {
	Login(username, password string) (token string, id int64, err error)
}
