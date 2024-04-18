package repository

import (
	"testskripsi/module/entities"
	"testskripsi/module/feature/pengaduan"

	"gorm.io/gorm"
)

type PengaduanRepository struct {
	db *gorm.DB
}

func NewPengaduanRepository(db *gorm.DB) pengaduan.RepositoryPengaduanInterface {
	return &PengaduanRepository{
		db: db,
	}
}

func (r *PengaduanRepository) GetPelangganByID(id uint64) ([]*entities.UserModels, error) {
	var pelanggan []*entities.UserModels

	if err := r.db.Where("id = ?", id).Find(&pelanggan).Error; err != nil {
		return nil, err
	}

	return pelanggan, nil
}

func (r *PengaduanRepository) CreatePengaduan(newData *entities.PengaduanModel) (*entities.PengaduanModel, error) {

	if err := r.db.Create(&newData).Error; err != nil {
		return nil, err
	}

	return newData, nil
}
func (r *PengaduanRepository) GetAllPengaduan() ([]*entities.PengaduanModel, error) {
	var pengaduan []*entities.PengaduanModel

	if err := r.db.Find(&pengaduan).Error; err != nil {
		return nil, err
	}

	return pengaduan, nil

}
func (r *PengaduanRepository) EditStatusPengaduan(id uint64, status string) error {
	var pengaduan = entities.PengaduanModel{}

	if err := r.db.Model(&pengaduan).Where("id = ?", id).Update("status", status).Error; err != nil {
		return err
	}

	return nil
}

func (r *PengaduanRepository) GetPengaduanByID(id uint64) (*entities.PengaduanModel, error) {
	var pengaduan = entities.PengaduanModel{}
	if err := r.db.Where("id = ?", id).First(&pengaduan).Error; err != nil {
		return nil, err
	}

	return &pengaduan, nil
}

func (r *PengaduanRepository) GetPengaduanByEmailPelanggan(email string) ([]*entities.PengaduanModel, error) {
	var pengaduan []*entities.PengaduanModel

	if err := r.db.Where("email = ?", email).Order("created_at desc").Find(&pengaduan).Error; err != nil {
		return nil, err
	}

	return pengaduan, nil
}

func (r *PengaduanRepository) GetStatusPengaduan(id uint64) (string, error) {
	var pengaduan = entities.PengaduanModel{}
	if err := r.db.Where("id = ?", id).First(&pengaduan).Error; err != nil {
		return "", err
	}

	return pengaduan.Status, nil
}
