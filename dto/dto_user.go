package dto

type UserDto struct {
	Id_user  int    `json:"id_user"`
	Name     string `json:"name"`
	LastName string `json:"last_name"`
	UserName string `json:"user_name"`
	Password string `json:"password"`
}

type UsersDto []UserDto

type LoginDto struct {
	UserName string `json:"user_name"`
	Password string `json:"password"`
}
