package dto

type AdminLoginDto struct {
	Telephone string `json:"telephone"`
	Password  string `json:"password"`
}

type AddAdminDto struct {
	Name      string `json:"name"`
	Email     string `json:"email"`
	Telephone string `json:"telephone"`
	Password  string `json:"password"`
	Authority int    `json:"authority"`
}

type GetUserDto struct {
	Page     int `json:"page"`
	PageSize int `json:"page_size"`
}

type UserStateDto struct {
	Uid   int `json:"uid"`
	State int `json:"state"`
}
