package dto

import "movie-app/model"

type UserDto struct {
	Name      string `json:"name"`
	Telephone string `json:"telephone"`
}
type RegisterDto struct {
	Name      string `json:"name"`
	Email     string `json:"email"`
	Telephone string `json:"telephone"`
	Password  string `json:"password"`
}

func ToUserDto(user model.User) UserDto {
	return UserDto{
		Name:      user.Name,
		Telephone: user.Telephone,
	}
}
