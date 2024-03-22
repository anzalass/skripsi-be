package service

import (
	"errors"
	"testskripsi/module/entities"
	"testskripsi/module/feature/auth"
	"testskripsi/utils"
)

type AuthService struct {
	repo auth.RepositoryAuthInterface
	jwt  utils.JWTInterface
}

func NewAuthService(repo auth.RepositoryAuthInterface, jwt utils.JWTInterface) auth.ServiceAuthInterface {
	return &AuthService{
		repo: repo,
		jwt:  jwt,
	}
}

func (s *AuthService) Register(newData *entities.AkunModel) (*entities.AkunModel, error) {
	value := &entities.AkunModel{
		Email:    newData.Email,
		Name:     newData.Name,
		Role:     "member",
		Password: newData.Password,
	}

	_, err := s.repo.CekEmail(newData.Email)
	if err != nil {
		return nil, errors.New("email tidak ada")
	}

	res2, err := s.repo.Register(value)
	if err != nil {
		return nil, errors.New("service : gagal membuat akun")
	}

	return res2, nil
}

func (s *AuthService) LoginAdmin(email, password string) (*entities.AkunModel, error) {
	res, err := s.repo.LoginAdmin(email)
	if err != nil {
		return nil, errors.New("Email not found")
	}

	if password != res.Password {
		return nil, errors.New("Password is incorrect")
	}

	return res, nil
}
