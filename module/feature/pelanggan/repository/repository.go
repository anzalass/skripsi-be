package repository

import (
	"testskripsi/module/entities"
	"testskripsi/module/feature/pelanggan"
	"time"

	"gorm.io/gorm"
)

type PelangganRepository struct {
	db *gorm.DB
}

func NewPelangganRepository(db *gorm.DB) pelanggan.RepositoryPelanggan {
	return &PelangganRepository{
		db: db,
	}
}

func (r *PelangganRepository) CreatePelanggan(newData *entities.UserModels) (*entities.UserModels, error) {
	if err := r.db.Create(&newData).Error; err != nil {
		return nil, err
	}

	return newData, nil
}
func (r *PelangganRepository) UpdatePelanggan(id int, newData *entities.UserModels) (bool, error) {
	pelanggan := entities.UserModels{}

	if err := r.db.Model(&pelanggan).Where("id = ? AND deleted_at IS NULL", id).Updates(&newData).Error; err != nil {
		return false, err
	}

	return true, nil
}

func (r *PelangganRepository) DeletePelanggan(id int) (bool, error) {
	pelanggan := entities.UserModels{}

	if err := r.db.Model(&pelanggan).Where("id = ? AND deleted_at IS NULL", id).Update("deleted_at", time.Now()).Error; err != nil {
		return false, err
	}

	return true, nil

}
func (r *PelangganRepository) GetPelangganByID(id int) (*entities.UserModels, error) {
	pelanggan := entities.UserModels{}

	if err := r.db.Where("id = ?", id).First(&pelanggan).Error; err != nil {
		return nil, err
	}

	return &pelanggan, nil

}

func (r *PelangganRepository) GetAllPelanggan() ([]*entities.UserModels, error) {
	var data []*entities.UserModels

	// offset := (page - 1) * perPage

	// if err := r.db.Offset(offset).Limit(perPage).Where("deleted_at = NULL").Find(&data).Error; err != nil {
	// 	return nil, err
	// }

	if err := r.db.Where("deleted_at IS NULL").Find(&data).Error; err != nil {
		return nil, err
	}

	return data, nil
}
func (r *PelangganRepository) GetAllPelangganForCreateInvoice() ([]*entities.UserModels, error) {
	var data []*entities.UserModels

	// offset := (page - 1) * perPage

	// if err := r.db.Offset(offset).Limit(perPage).Where("deleted_at = NULL").Find(&data).Error; err != nil {
	// 	return nil, err
	// }

	if err := r.db.Where("status = ? AND deleted_at IS NULL", "aktif").Find(&data).Error; err != nil {
		return nil, err
	}

	return data, nil
}

func (r *PelangganRepository) GetAllDetailPelanggan(id uint64) (*entities.UserModels, error) {
	data := entities.UserModels{}
	if err := r.db.Preload("Tagihan").Preload("Transaksi").Where("id = ?", id).Find(&data).Error; err != nil {
		return nil, err
	}

	return &data, nil
}

func (r *PelangganRepository) GetIdAkunByEmail(email string) (uint64, error) {
	data := entities.AkunModel{}
	if err := r.db.Where("email = ?", email).First(&data).Error; err != nil {
		return 0, err
	}

	var id uint64
	if err := r.db.Model(&data).Select("id").Scan(&id).Error; err != nil {
		return 0, err
	}

	return id, nil
}
