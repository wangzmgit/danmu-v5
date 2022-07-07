package dto

//注册
type RegisterDto struct {
	Email    string
	Password string
	Code     string
}

//登录
type LoginDto struct {
	Email    string
	Password string
}

//修改用户信息
type ModifyUserDto struct {
	Name     string
	Gender   int
	Birthday string
	Sign     string
}

//管理员修改用户信息
type AdminModifyUserDto struct {
	ID    uint
	Name  string
	Email string
	Sign  string
}

type ModifyRoleDto struct {
	UID  uint
	Role int
}
