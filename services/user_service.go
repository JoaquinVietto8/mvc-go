package services

import (
	"fmt"
	userCliente "mvc-go/clients/user"
	"mvc-go/dto"
	"mvc-go/model"
	e "mvc-go/utils/errors"
)

type userService struct{}

type userServiceInterface interface {
	GetUserById(id int) (dto.UserDto, e.ApiError)
	GetUsers() (dto.UsersDto, e.ApiError)
	LoginUser(loginDto dto.LoginDto) (dto.TokenDto, e.ApiError)
}

var (
	UserService userServiceInterface
)

func init() {
	UserService = &userService{}
}

func (s *userService) GetUserById(id int) (dto.UserDto, e.ApiError) {

	var user model.User = userCliente.GetUserById(id)
	var userDto dto.UserDto

	if user.Id == 0 {
		return userDto, e.NewBadRequestApiError("user not found")
	}
	userDto.Name = user.Name
	userDto.LastName = user.LastName
	userDto.UserName = user.UserName
	userDto.Id_user = user.Id
	return userDto, nil
}

func (s *userService) GetUsers() (dto.UsersDto, e.ApiError) {

	var users model.Users = userCliente.GetUsers()
	var usersDto dto.UsersDto

	for _, user := range users {
		var userDto dto.UserDto
		userDto.Name = user.Name
		userDto.LastName = user.LastName
		userDto.UserName = user.UserName
		userDto.Id_user = user.Id

		usersDto = append(usersDto, userDto)
	}

	return usersDto, nil
}

func (s *userService) LoginUser(loginDto dto.LoginDto) (dto.TokenDto, e.ApiError) { //develve (dto.TokenDto, e.ApiError)
	var tokenDto dto.TokenDto //genera un token dto vacio

	user := userCliente.GetUserByUserName(loginDto)

	if user.Id == 0 {
		return tokenDto, e.NewBadRequestApiError("user not found")
	}

	if user.Password != loginDto.Password {
		return tokenDto, e.NewBadRequestApiError("Bad password")
	}

	tokenDto.Token = fmt.Sprintf("%d", user.Id)

	return tokenDto, nil //crear token, si no hay errores y devuelve nil porque no hay error

}
