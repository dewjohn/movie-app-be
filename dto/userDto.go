package dto

import "time"

type RegisterDto struct {
	Name      string `json:"name"`
	Email     string `json:"email"`
	Telephone string `json:"telephone"`
	Password  string `json:"password"`
}

type LoginDto struct {
	Telephone string `json:"telephone"`
	Password  string `json:"password"`
}

type UserModifyDto struct {
	Name     string `json:"name"`
	Birthday string `json:"birthday"`
	Gender   int    `json:"gender"`
	Sign     string `json:"sign"`
}

type UserModifyPasswordDto struct {
	OldPassword string `json:"oldPassword"`
	Password    string `json:"password"`
}

type UserInfoDto struct {
	Id   uint   `json:"id"`
	Name string `json:"name"`
}

type UserInfoToAdminDto struct {
	Id        uint      `json:"id"`
	Avatar    string    `json:"avatar"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Telephone string    `json:"telephone"`
	Gender    int       `json:"gender"`
	Birthday  time.Time `json:"birthday"`
	Sign      string    `json:"sign"`
	State     int       `json:"state"`
}
