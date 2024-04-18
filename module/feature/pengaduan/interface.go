package pengaduan

import (
	"testskripsi/module/entities"

	"github.com/labstack/echo/v4"
)

type RepositoryPengaduanInterface interface {
	GetPelangganByID(id uint64) ([]*entities.UserModels, error)
	CreatePengaduan(newData *entities.PengaduanModel) (*entities.PengaduanModel, error)
	GetAllPengaduan() ([]*entities.PengaduanModel, error)
	EditStatusPengaduan(id uint64, status string) error
	GetPengaduanByID(id uint64) (*entities.PengaduanModel, error)
	GetPengaduanByEmailPelanggan(email string) ([]*entities.PengaduanModel, error)
	GetStatusPengaduan(id uint64) (string, error)
}

type ServicePengaduanInterface interface {
	GetPelangganByID(id uint64) ([]*entities.UserModels, error)
	CreatePengaduan(newData *entities.PengaduanModel, filename string, file interface{}, waktu string) (*entities.PengaduanModel, error)
	GetAllPengaduan() ([]*entities.PengaduanModel, error)
	EditStatusPengaduan(id uint64, status string) error
	GetPengaduanByID(id uint64) (*entities.PengaduanModel, error)
	GetPengaduanByEmailPelanggan(email string) ([]*entities.PengaduanModel, error)
	GetStatusPengaduan(id uint64) (string, error)
}

type HandlerPengaduanInterface interface {
	GetPelangganByID() echo.HandlerFunc
	CreatePengaduan() echo.HandlerFunc
	GetAllPengaduan() echo.HandlerFunc
	EditStatusPengaduan() echo.HandlerFunc
	GetPengaduanByID() echo.HandlerFunc
	GetPengaduanByEmailPelanggan() echo.HandlerFunc
	GetStatusPengaduan() echo.HandlerFunc
}
