package auth

import (
	"efishery-golang/entity"
	"efishery-golang/repository"
	"efishery-golang/utils/stringify"
	"github.com/golang-jwt/jwt"
	"os"
)

type ServiceInterface interface {
	Login(
		body LoginRequest,
	) (tokenString string, err error)
	Register(
		body RegisterRequest,
	) (res RegisterResponse, err error)
}

type Service struct {
	usersRepo repository.User
}

func NewService(usersRepo repository.User) Service {
	return Service{
		usersRepo: usersRepo,
	}
}

func (s Service) Login(body LoginRequest) (tokenString string, err error) {

	user, err := s.usersRepo.FindByPhonePassword(body.Phone, body.Password)
	if user.ID == 0 {
		return "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"data": map[string]string{
			"name":  user.Name,
			"phone": body.Phone,
			"role":  user.Role,
		},
	})

	// Sign the token with a secret key
	tokenString, err = token.SignedString([]byte(os.Getenv("SECRET")))
	return tokenString, err
}

func (s Service) Register(body RegisterRequest) (res RegisterResponse, err error) {
	user, err := s.usersRepo.FindByPhone(body.Phone)
	if user.ID != 0 {
		return RegisterResponse{}, err
	}

	password := stringify.GenRandomString(4)
	newUser, err := s.usersRepo.StoreUser(entity.Users{
		Phone:    body.Phone,
		Name:     body.Name,
		Role:     body.Role,
		Password: password,
	})

	data := RegisterResponse{
		ID:       int(newUser.ID),
		Phone:    newUser.Phone,
		Name:     newUser.Name,
		Role:     newUser.Role,
		Password: password,
	}

	return data, err
}
