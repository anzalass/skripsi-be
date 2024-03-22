package auth

import (
	"testskripsi/module/entities"

	"github.com/labstack/echo/v4"
)

type RepositoryAuthInterface interface {
	Register(newData *entities.AkunModel) (*entities.AkunModel, error)
	LoginAdmin(email string) (*entities.AkunModel, error)
	CekEmail(email string) (bool, error)
}
type ServiceAuthInterface interface {
	Register(newData *entities.AkunModel) (*entities.AkunModel, error)
	LoginAdmin(email string, password string) (*entities.AkunModel, error)
}
type HandlerAuthInterface interface {
	Register() echo.HandlerFunc
	LoginAdmin() echo.HandlerFunc
	GetUserFomCookies() echo.HandlerFunc
}
