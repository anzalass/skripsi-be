package service

import (
	"errors"
	"testskripsi/module/entities"
	"testskripsi/module/feature/pelanggan"
	"testskripsi/module/feature/pengaduan"
	"testskripsi/utils"
)

type ServicePengaduan struct {
	repo     pengaduan.RepositoryPengaduanInterface
	repouser pelanggan.RepositoryPelanggan
}

func NewPengaduanService(repo pengaduan.RepositoryPengaduanInterface, repouser pelanggan.RepositoryPelanggan) pengaduan.ServicePengaduanInterface {
	return &ServicePengaduan{
		repo:     repo,
		repouser: repouser,
	}
}

func (s *ServicePengaduan) GetPelangganByID(id string) ([]*entities.UserModels, error) {
	res, err := s.repo.GetPelangganByID(id)
	if err != nil {
		return nil, errors.New(err.Error())
	}

	return res, nil
}
func (s *ServicePengaduan) CreatePengaduan(newData *entities.PengaduanModel, filename string, file interface{}, waktu string) (*entities.PengaduanModel, error) {
	upload, err := utils.ImageHandler(file, filename)
	if err != nil {
		return nil, errors.New(err.Error())
	}

	// originalTimeStr := "Wed Apr 17 2024 10:05:56 GMT+0700 (Indochina Time)"

	res2, err2 := s.repouser.GetIdAkunByEmail(newData.Email)
	if err2 != nil {
		return nil, errors.New(err2.Error())
	}

	value := &entities.PengaduanModel{
		IDPelanggan:    newData.IDPelanggan,
		IDAkun:         res2,
		Nama:           newData.Nama,
		Email:          newData.Email,
		Deskripsi:      newData.Deskripsi,
		Alamat:         newData.Alamat,
		Foto:           upload,
		NoWhatsapp:     newData.NoWhatsapp,
		WaktuKunjungan: waktu,
		Status:         "diproses",
	}

	res, err := s.repo.CreatePengaduan(value)
	if err != nil {
		return nil, errors.New(err.Error())
	}

	return res, nil
}
func (s *ServicePengaduan) GetAllPengaduan() ([]*entities.PengaduanModel, error) {
	res, err := s.repo.GetAllPengaduan()
	if err != nil {
		return nil, errors.New(err.Error())
	}

	return res, nil
}
func (s *ServicePengaduan) EditStatusPengaduan(id uint64, status string) error {
	err := s.repo.EditStatusPengaduan(id, status)
	if err != nil {
		return errors.New(err.Error())
	}

	return nil
}

func (s *ServicePengaduan) GetPengaduanByID(id uint64) (*entities.PengaduanModel, error) {
	res, err := s.repo.GetPengaduanByID(id)
	if err != nil {
		return nil, errors.New(err.Error())
	}

	return res, nil
}

func (s *ServicePengaduan) GetPengaduanByEmailPelanggan(email string) ([]*entities.PengaduanModel, error) {
	res, err := s.repo.GetPengaduanByEmailPelanggan(email)
	if err != nil {
		return nil, errors.New(err.Error())
	}

	return res, nil
}

func (s *ServicePengaduan) GetStatusPengaduan(id uint64) (string, error) {
	res, err := s.repo.GetStatusPengaduan(id)
	if err != nil {
		return "", errors.New(err.Error())
	}

	return res, nil
}
