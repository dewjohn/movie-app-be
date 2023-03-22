package vo

import "movie-app/model"

type AdminVo struct {
	ID        uint   `json:"id"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	Telephone string `json:"telephone"`
	Authority int    `json:"authority"`
}

func ToAdminVo(admin model.Admin) AdminVo {
	return AdminVo{
		ID:        admin.ID,
		Name:      admin.Name,
		Email:     admin.Email,
		Telephone: admin.Telephone,
		Authority: admin.Authority,
	}
}
