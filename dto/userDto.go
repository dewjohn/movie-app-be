package dto

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
	Avatar   string `json:"avatar"` // 暂定为头像外链链接
}

type UserModifyPasswordDto struct {
	OldPassword string `json:"oldPassword"`
	Password    string `json:"password"`
}
