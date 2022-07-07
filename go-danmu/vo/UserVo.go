package vo

import (
	"time"

	"kuukaa.fun/danmu-v5/model"
	"kuukaa.fun/danmu-v5/util"
)

//用户信息
type UserVo struct {
	ID       uint      `json:"uid"`
	Email    string    `json:"email"`
	Name     string    `json:"name"`
	Sign     string    `json:"sign"`
	Avatar   string    `json:"avatar"`
	Gender   int       `json:"gender"`
	Role     int       `json:"role"`
	Birthday time.Time `json:"birthday"`
}

//用户列表
type UserList struct {
	ID       uint      `json:"uid"`
	Name     string    `json:"name"`
	Sign     string    `json:"sign"`
	Avatar   string    `json:"avatar"`
	Gender   int       `json:"gender"`
	Birthday time.Time `json:"birthday"`
}

//管理员获取的用户列表
type AdminUserListVo struct {
	ID        uint      `json:"uid"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Sign      string    `json:"sign"`
	Avatar    string    `json:"avatar"`
	Gender    int       `json:"gender"`
	Role      int       `json:"role"`
	CreatedAt time.Time `json:"created_at"`
}

func ToUserVo(user model.User) UserVo {
	return UserVo{
		ID:       user.ID,
		Name:     user.Name,
		Email:    util.HideEmail(user.Email),
		Sign:     user.Sign,
		Avatar:   user.Avatar,
		Gender:   user.Gender,
		Role:     user.Role,
		Birthday: user.Birthday,
	}
}
