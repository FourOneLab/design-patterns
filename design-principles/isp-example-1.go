package design_principles

type UserService interface {
	Register(cellphone, password string)
	Login(cellphone, password string)
	GetUserInfoById(id int64) UserInfo
	GetUserInfoByCellphone(cellphone string) UserInfo
}

type UserInfo struct {
	Id        int64
	Name      string
	Password  string
	Cellphone string
}

type UserServiceImpl struct{}

func (u *UserServiceImpl) Register(cellphone, password string) {
	panic("implement me")
}

func (u *UserServiceImpl) Login(cellphone, password string) {
	panic("implement me")
}

func (u *UserServiceImpl) GetUserInfoById(id int64) UserInfo {
	panic("implement me")
}

func (u *UserServiceImpl) GetUserInfoByCellphone(cellphone string) UserInfo {
	panic("implement me")
}

// ------------------------------
// 增加删除用户功能

type RestrictedUserService interface {
	DeleteUserByCellphone(cellphone string)
	DeleteUserById(id int64)
}

type RestrictedUserServiceImpl struct{}

func (r *RestrictedUserServiceImpl) DeleteUserByCellphone(cellphone string) {
	panic("implement me")
}

func (r *RestrictedUserServiceImpl) DeleteUserById(id int64) {
	panic("implement me")
}
