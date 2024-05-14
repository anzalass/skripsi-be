package invoice

import (
	"testskripsi/module/entities"

	"github.com/labstack/echo/v4"
)

type RepositoryInvoice interface {
	GetAllData() ([]*entities.TagihanModels, error)
	CreateInvoice(newData *entities.TagihanModels) (*entities.TagihanModels, error)
	GetTagihanByIdPelanggan(id_pelanggan string) ([]*entities.TagihanModels, error)
	CreateTransaksi(nedData *entities.TransaksiModels) (*entities.TransaksiModels, error)
	GetPeriodePemakaianByIdPembayaranTransaksi(idpembayaran string) (string, error)
	GetTagihanByPeriodePemakaian(periode string) ([]*entities.TagihanModels, error)
	UpdateStatusTagihan(periode string) (bool, error)
	UpdateStatusTransaksi(id string) (bool, error)
	GetTransaksiByEmail(email string, page uint64) ([]*entities.TransaksiModels, error)
	GetTransaksiByIdPembayaran(idpembayaran string) (*entities.TransaksiModels, error)
	GetAllPembayaran() ([]*entities.TransaksiModels, error)
	GetTransaksiByIdPelangganBeforeExpired(id uint64, timenow uint64) ([]*entities.TransaksiModels, error)
	DeleteTransaksiPending(periodepemakaian string) (bool, error)
	// GetTransaksiBulanan() ([]map[string]interface{}, error)
}

type ServiceInvoice interface {
	CreateAllInvoice(bulan string, tahun uint64) (bool, error)
	GetAllData() ([]*entities.TagihanModels, error)
	GetTagihanByIdPelanggan(id string) (any, error)
	CreateTransaksi(newData *entities.TransaksiModels) (any, error)
	ConfirmedPayment(idpembayaran string) (bool, error)
	GetTransaksiByEmail(email string, page uint64) ([]*entities.TransaksiModels, error)
	GetTransaksiByIdPembayaran(idpembayaran string) (*entities.TransaksiModels, error)
	GetAllPembayaran() ([]*entities.TransaksiModels, error)
	// GetTransaksiBulanan() ([]map[string]interface{}, error)
}

type HandlerInvoice interface {
	CreateAllInvoice() echo.HandlerFunc
	GetAllInvoice() echo.HandlerFunc
	GetInvoiceById() echo.HandlerFunc
	CreateTransaksi() echo.HandlerFunc
	AfterPayment() echo.HandlerFunc
	GetTransaksiByEmail() echo.HandlerFunc
	GetTransaksiByIdPembayaran() echo.HandlerFunc
	GetAllPembayaran() echo.HandlerFunc

	// GetTransaksiBulanan() echo.HandlerFunc
}
