package pelanggan

import (
	"testskripsi/module/entities"

	"github.com/labstack/echo/v4"
)

type RepositoryPelanggan interface {
	CreatePelanggan(newData *entities.UserModels) (*entities.UserModels, error)
	UpdatePelanggan(id int, newData *entities.UserModels) (bool, error)
	DeletePelanggan(id int) (bool, error)
	GetAllPelanggan() ([]*entities.UserModels, error)
	GetAllPelangganForCreateInvoice() ([]*entities.UserModels, error)
	GetPelangganByID(id int) (*entities.UserModels, error)
}
type ServicePelanggan interface {
	CreatePelanggan(newData *entities.UserModels) (*entities.UserModels, error)
	UpdatePelanggan(id int, newData *entities.UserModels) (bool, error)
	DeletePelanggan(id int) (bool, error)
	GetAllPelanggan() ([]*entities.UserModels, error)
	GetPelangganByID(id int) (*entities.UserModels, error)
}
type HandlerPelanggan interface {
	CreatePelanggan() echo.HandlerFunc
	UpdatePelanggan() echo.HandlerFunc
	DeletePelanggan() echo.HandlerFunc
	GetAllPelanggan() echo.HandlerFunc
	GetPelangganByID() echo.HandlerFunc
}
