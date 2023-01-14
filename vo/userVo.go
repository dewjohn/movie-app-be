package vo

import (
	"movie-app/model"
	"time"
)

type UserVo struct {
	ID       uint      `json:"uid"`
	Email    string    `json:"email"`
	Name     string    `json:"name"`
	Sign     string    `json:"sign"`
	Avatar   string    `json:"avatar"`
	Gender   int       `json:"gender"`
	Birthday time.Time `json:"birthday"`
}

func ToUserVo(user model.User) UserVo {
	return UserVo{
		ID:       user.ID,
		Name:     user.Name,
		Email:    user.Email,
		Sign:     user.Sign,
		Avatar:   user.Avatar,
		Gender:   user.Gender,
		Birthday: user.Birthday,
	}
}
