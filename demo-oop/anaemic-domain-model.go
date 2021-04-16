package demo_oop

// Controller + VO(view object)
// 负责暴露接口，将 Service 中获取的 BO 转换为 VO 并返回给前端

type UserVo struct {
	Id        string
	Name      string
	Cellphone string
}

type UserController struct {
	userService UserService
}

func (c UserController) GetUserById(userId string) UserVo {
	userBo := c.userService.GetUserById(userId)

	convertUserBoToUserVo := func(bo UserBo) UserVo {
		return UserVo{}
	}

	userVo := convertUserBoToUserVo(userBo)
	return userVo
}

// ------------------------------
// Repository + Entity
// 负责数据访问，将数据库表中的数据转换为 Entity 并返回给 Service。

type UserEntity struct {
	Id        string
	Name      string
	Cellphone string
}

type UserRepository struct{}

func (r UserRepository) GetUserById(userId string) UserEntity {
	return UserEntity{}
}

// ------------------------------
// Service + BO(business object)
// 负责业务逻辑，将 Repository 中的 Entity 转换为 BO 并返回给 Controller。

type UserBo struct {
	Id        string
	Name      string
	Cellphone string
}

type UserService struct {
	userRepository UserRepository
}

func (s UserService) GetUserById(userId string) UserBo {
	userEntity := s.userRepository.GetUserById(userId)

	convertUserEntityToUserBo := func(entity UserEntity) UserBo {
		return UserBo{}
	}

	userBo := convertUserEntityToUserBo(userEntity)
	return userBo
}
