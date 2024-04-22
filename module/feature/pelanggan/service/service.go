package service

import (
	"errors"
	"testskripsi/module/entities"
	"testskripsi/module/feature/pelanggan"
	"time"
)

type PelangganService struct {
	repo pelanggan.ServicePelanggan
}

func NewPelangganService(repo pelanggan.RepositoryPelanggan) pelanggan.ServicePelanggan {
	return &PelangganService{
		repo: repo,
	}
}

func (s *PelangganService) CreatePelanggan(newData *entities.UserModels) (*entities.UserModels, error) {
	value := &entities.UserModels{
		Name:           newData.Name,
		Alamat:         newData.Alamat,
		Status:         "aktif",
		PaketLangganan: newData.PaketLangganan,
		HargaLangganan: newData.HargaLangganan,
	}

	res, err := s.repo.CreatePelanggan(value)
	if err != nil {
		return nil, errors.New("service gagal membuat pelanggan")
	}
	return res, nil
}
func (s *PelangganService) UpdatePelanggan(id int, newData *entities.UserModels) (bool, error) {

	value := &entities.UserModels{
		Name:           newData.Name,
		Alamat:         newData.Alamat,
		Status:         newData.Status,
		PaketLangganan: newData.PaketLangganan,
		HargaLangganan: newData.HargaLangganan,
		UpdatedAt:      time.Now(),
	}

	res, err := s.repo.UpdatePelanggan(id, value)
	if err != nil {
		return false, errors.New("service gagal update pelanggan")
	}

	return res, nil
}
func (s *PelangganService) DeletePelanggan(id int) (bool, error) {
	res, err := s.repo.DeletePelanggan(id)
	if err != nil {
		return false, errors.New("service gagal hapus pelanggan")
	}

	return res, nil
}

func (s *PelangganService) GetAllPelanggan() ([]*entities.UserModels, error) {
	res, err := s.repo.GetAllPelanggan()
	if err != nil {
		return nil, errors.New("service gagal dapatkan semua pelanggan")
	}

	return res, nil
}
func (s *PelangganService) GetPelangganByID(id int) (*entities.UserModels, error) {
	res, err := s.repo.GetPelangganByID(id)
	if err != nil {
		return nil, errors.New("service gagal dapatkan semua pelanggan")
	}

	return res, nil
}

func (s *PelangganService) GetAllDetailPelanggan(id uint64) (*entities.UserModels, error) {
	res, err := s.repo.GetAllDetailPelanggan(id)
	if err != nil {
		return nil, err
	}

	return res, nil
}
