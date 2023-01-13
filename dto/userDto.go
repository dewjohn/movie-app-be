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
