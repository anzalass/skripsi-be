package service

import (
	"errors"
	"testskripsi/module/entities"
	"testskripsi/module/feature/pelanggan"
	"time"
)

type PelangganService struct {
	repo pelanggan.RepositoryPelanggan
}

func NewPelangganService(repo pelanggan.RepositoryPelanggan) pelanggan.ServicePelanggan {
	return &PelangganService{
		repo: repo,
	}
}

func (s *PelangganService) CreatePelanggan(newData *entities.UserModels) (*entities.UserModels, error) {
	value := &entities.UserModels{
		ID:             newData.ID,
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
func (s *PelangganService) UpdatePelanggan(id string, newData *entities.UserModels) (bool, error) {

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
func (s *PelangganService) DeletePelanggan(id string) (bool, error) {
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
func (s *PelangganService) GetPelangganByID(id string) (*entities.UserModels, error) {
	res, err := s.repo.GetPelangganByID(id)
	if err != nil {
		return nil, errors.New("service gagal dapatkan semua pelanggan")
	}

	return res, nil
}

func (s *PelangganService) GetAllDetailPelanggan(id string) (*entities.UserModels, error) {
	res, err := s.repo.GetAllDetailPelanggan(id)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (s *PelangganService) CheckIdUserByEmail(email string) (string, error) {
	res, err := s.repo.CheckIdUserByEmail(email)
	if err != nil {
		return "error", nil
	}

	return res, nil
}

func (s *PelangganService) InsertIdUserByEmail(email string, iduser string) (bool, error) {
	err := s.repo.SetNullIdUser(iduser)
	if err != nil {
		return true, err
	}

	res2, _ := s.repo.GetPelangganByID(iduser)
	if res2 != nil {
		_, err := s.repo.InsertIdUserByEmail(email, iduser)
		if err != nil {
			return true, err
		}
	} else {
		return false, err
	}

	return true, nil
}
