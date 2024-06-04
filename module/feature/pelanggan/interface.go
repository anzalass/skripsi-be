package pelanggan

import (
	"testskripsi/module/entities"

	"github.com/labstack/echo/v4"
)

type RepositoryPelanggan interface {
	CreatePelanggan(newData *entities.UserModels) (*entities.UserModels, error)
	UpdatePelanggan(id string, newData *entities.UserModels) (bool, error)
	DeletePelanggan(id string) (bool, error)
	GetAllPelanggan() ([]*entities.UserModels, error)
	GetAllPelangganForCreateInvoice() ([]*entities.UserModels, error)
	GetPelangganByID(id string) (*entities.UserModels, error)
	GetAllDetailPelanggan(id string) (*entities.UserModels, error)
	GetIdAkunByEmail(email string) (uint64, error)
	InsertIdUserByEmail(email string, idakun string) (*entities.AkunModel, error)
	CheckIdUserByEmail(email string) (string, error)
	SetNullIdUser(iduser string) error
	GetNoWhatsApp(iduser string) (string, error)
}
type ServicePelanggan interface {
	CreatePelanggan(newData *entities.UserModels) (*entities.UserModels, error)
	UpdatePelanggan(id string, newData *entities.UserModels) (bool, error)
	DeletePelanggan(id string) (bool, error)
	GetAllPelanggan() ([]*entities.UserModels, error)
	GetPelangganByID(id string) (*entities.UserModels, error)
	GetAllDetailPelanggan(id string) (*entities.UserModels, error)
	InsertIdUserByEmail(email string, iduser string) (*entities.AkunModel, error)
	CheckIdUserByEmail(email string) (string, error)
}
type HandlerPelanggan interface {
	CreatePelanggan() echo.HandlerFunc
	UpdatePelanggan() echo.HandlerFunc
	DeletePelanggan() echo.HandlerFunc
	GetAllPelanggan() echo.HandlerFunc
	GetPelangganByID() echo.HandlerFunc
	GetAllDetailPelanggan() echo.HandlerFunc
	InsertIdUserByEmail() echo.HandlerFunc
	CheckIdUserByEmail() echo.HandlerFunc
}
