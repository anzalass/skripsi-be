package repository

import (
	"testskripsi/module/entities"
	"testskripsi/module/feature/invoice"

	"gorm.io/gorm"
)

type InvoiceRepository struct {
	db *gorm.DB
}

func NewInvoiceRepository(db *gorm.DB) invoice.RepositoryInvoice {
	return &InvoiceRepository{
		db: db,
	}
}

func (r *InvoiceRepository) GetAllData() ([]*entities.TagihanModels, error) {
	var data []*entities.TagihanModels
	err := r.db.Where("deleted_at IS NULL").Find(&data).Error
	if err != nil {
		return data, err
	}
	return data, nil
}

func (r *InvoiceRepository) CreateInvoice(newData *entities.TagihanModels) (*entities.TagihanModels, error) {
	if err := r.db.Create(&newData).Error; err != nil {
		return newData, err
	}
	return newData, nil
}

func (r *InvoiceRepository) GetTagihanByIdPelanggan(id uint64) ([]*entities.TagihanModels, error) {
	var data []*entities.TagihanModels

	if err := r.db.Where("id_pelanggan = ? AND deleted_at IS NULL AND status = ?", id, "belum lunas").Find(&data).Error; err != nil {
		return nil, err
	}

	return data, nil
}

func (r *InvoiceRepository) CreateTransaksi(newData *entities.TransaksiModels) (*entities.TransaksiModels, error) {
	if err := r.db.Create(&newData).Error; err != nil {
		return newData, err
	}
	return newData, nil
}

func (r *InvoiceRepository) GetTagihanByPeriodePemakaian(periode string) ([]*entities.TagihanModels, error) {
	var data []*entities.TagihanModels
	if err := r.db.Where("periode_pemakaian = ?", periode).Find(&data).Error; err != nil {
		return nil, err
	}

	return data, nil
}

func (r *InvoiceRepository) GetPeriodePemakaianByIdPembayaranTransaksi(idpembayaran string) (string, error) {
	var data entities.TransaksiModels
	if err := r.db.Where("id_pembayaran", idpembayaran).First(&data).Error; err != nil {
		return "", err
	}

	return data.PeriodePemakaian, nil
}

func (r *InvoiceRepository) UpdateStatusTagihan(periode string) (bool, error) {
	var data entities.TagihanModels
	if err := r.db.Model(&data).Where("periode_pemakaian = ?", periode).Update("status", "lunas").Error; err != nil {
		return false, err
	}
	return true, nil
}
func (r *InvoiceRepository) UpdateStatusTransaksi(id string) (bool, error) {
	var data entities.TransaksiModels
	if err := r.db.Model(&data).Where("id_pembayaran = ?", id).Update("status", "success").Error; err != nil {
		return false, err
	}
	return true, nil
}

func (r *InvoiceRepository) GetTransaksiByEmail(email string, page uint64) ([]*entities.TransaksiModels, error) {
	var data []*entities.TransaksiModels

	offset := (page - 1) * 8

	if err := r.db.Where("email = ? AND status = ? AND deleted_at IS NULL", email, "success").Limit(8).Offset(int(offset)).Find(&data).Error; err != nil {
		return nil, err
	}

	return data, nil
}

func (r *InvoiceRepository) GetTransaksiByIdPembayaran(idpembayaran string) (*entities.TransaksiModels, error) {
	var data *entities.TransaksiModels
	if err := r.db.Where("id_pembayaran = ? AND status = ? AND deleted_at IS null ", idpembayaran, "success").First(&data).Error; err != nil {
		return nil, err
	}

	return data, nil
}

func (r *InvoiceRepository) GetAllPembayaran() ([]*entities.TransaksiModels, error) {
	var data []*entities.TransaksiModels
	if err := r.db.Where("deleted_at IS NULL").Find(&data).Error; err != nil {
		return nil, err
	}

	return data, nil
}

// func (r *InvoiceRepository) GetTransaksiBulanan() ([]*map[any]interface{}, error) {
// 	var data []*entities.TransaksiModels
// 	if err := r.db.Where("tanggal_bayar BETWEEN ? AND ?", "2024-03-01", "2024-03-31").Select("tanggal_bayar", "total_amount").Find(&data).Error; err != nil {
// 		return nil, err
// 	}

// 	return data, nil
// }
