package user

type ServiceStruct struct{}

func (u *ServiceStruct) Login(username, password string) (token string, id int64, err error) {
	return "测试token", 1, nil
}

func NewUserService() *ServiceStruct {
	return &ServiceStruct{}
}

var _ Service = (*ServiceStruct)(nil)
