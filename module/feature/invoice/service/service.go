package service

import (
	"context"
	"errors"
	"fmt"
	_ "fmt"
	"strings"

	"testskripsi/module/entities"
	"testskripsi/module/feature/invoice"
	"testskripsi/module/feature/pelanggan"
	"testskripsi/utils/midtrans"
	"time"

	"github.com/sirupsen/logrus"
)

type InvoiceService struct {
	repo     invoice.RepositoryInvoice
	repouser pelanggan.RepositoryPelanggan
	midtrans midtrans.MidtransServiceInterface
}

func NewInvoiceService(repo invoice.RepositoryInvoice, repouser pelanggan.RepositoryPelanggan, midtrans midtrans.MidtransServiceInterface) invoice.ServiceInvoice {
	return &InvoiceService{
		repo:     repo,
		repouser: repouser,
		midtrans: midtrans,
	}
}

func (s *InvoiceService) CreateAllInvoice(bulan string, tahun uint64) (bool, error) {
	alldata, err := s.repouser.GetAllPelangganForCreateInvoice()
	if err != nil {
		logrus.Error("Error")
	}

	for _, data := range alldata {
		// Access each userModel here
		newInvoice := &entities.TagihanModels{
			IDPelanggan:      data.ID,
			Name:             data.Name,
			Status:           "belum lunas",
			PaketLangganan:   data.PaketLangganan,
			HargaLangganan:   data.HargaLangganan,
			Bulan:            bulan,
			Alamat:           data.Alamat,
			Tahun:            tahun,
			PeriodePemakaian: fmt.Sprintf("ID%s%s%d", data.ID, bulan, tahun),
		}

		// strid := strconv.FormatUint(data.ID, 16)
		// // fmt.Println("woi", strid)

		_, err := s.repo.CreateInvoice(newInvoice)
		if err != nil {
			logrus.Error("Failed to create invoice")
		}

	}
	return true, nil
}

func (s *InvoiceService) GetAllData() ([]*entities.TagihanModels, error) {
	res, err := s.repo.GetAllData()
	if err != nil {
		return nil, errors.New("error service get all invoice data")
	}

	return res, nil
}

func (s *InvoiceService) GetTagihanByIdPelanggan(id string) (any, error) {

	res, err := s.repo.GetTagihanByIdPelanggan(id)
	if err != nil {
		return nil, err
	}

	return res, nil

}

func (s *InvoiceService) CreateTransaksi(newData *entities.TransaksiModels) (any, error) {

	var ctx context.Context

	res, res2, err := s.midtrans.GenerateSnapURL(ctx, int64(newData.TotalAmount))
	if err != nil {
		return "", errors.New("gagal membuat snap url")

	}

	res4, err4 := s.repouser.GetIdAkunByEmail(newData.Email)
	if err4 != nil {
		return "", errors.New("gagal membuat snap url")
	}

	value := &entities.TransaksiModels{
		IDPelanggan:      newData.IDPelanggan,
		IDAkun:           res4,
		IDPembayaran:     res2,
		Email:            newData.Email,
		Name:             newData.Name,
		Status:           "pending",
		Alamat:           newData.Alamat,
		Url:              res,
		PeriodePemakaian: newData.PeriodePemakaian,
		TanggalBayar:     time.Now(),
		PaketLangganan:   newData.PaketLangganan,
		Expired:          uint64(time.Now().Add(1 * 23 * time.Hour).Unix()),
		HargaLangganan:   newData.HargaLangganan,
		TotalAmount:      newData.TotalAmount,
		MetodePembayaran: "midtrans",
	}

	_, err2 := s.repo.CreateTransaksi(value)
	if err2 != nil {
		return "", errors.New("gagal membuat transaksi")
	}

	return res, nil
}

func (s *InvoiceService) ConfirmedPayment(idpembayaran string) (bool, error) {
	res, err := s.repo.GetPeriodePemakaianByIdPembayaranTransaksi(idpembayaran)
	if err != nil {
		return false, err
	}

	_, err3 := s.repo.UpdateStatusTransaksi(idpembayaran)
	if err3 != nil {
		return false, err3
	}

	parts := strings.Fields(res)
	for _, part := range parts {
		_, err2 := s.repo.UpdateStatusTagihan(part)
		if err2 != nil {
			return false, err2
		}
		_, err := s.repo.DeleteTransaksiPending(res)
		if err != nil {
			return false, err
		}
	}

	return true, nil

}

func (s *InvoiceService) GetTransaksiByEmail(email string, page uint64) ([]*entities.TransaksiModels, error) {
	res, err := s.repo.GetTransaksiByEmail(email, page)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (s *InvoiceService) GetTransaksiByIdPembayaran(idpembayaran string) (*entities.TransaksiModels, error) {
	res, err := s.repo.GetTransaksiByIdPembayaran(idpembayaran)
	if err != nil {
		return nil, err
	}

	return res, nil
}
func (s *InvoiceService) GetAllPembayaran() ([]*entities.TransaksiModels, error) {
	res, err := s.repo.GetAllPembayaran()
	if err != nil {
		return nil, err
	}
	return res, nil
}

// func (s *InvoiceService) GetTransaksiBulanan() ([]map[string]interface{}, error) {
// 	res, err := s.repo.GetTransaksiBulanan()
// 	if err != nil {
// 		return nil, err
// 	}

// 	groupedData := make(map[string]int)
// 	for _, data := range res {
// 		tanggal_bayar := data["tanggal_bayar"].(string)
// 		total_amount := data["total_amount"].(uint64)

// 		if _, ok := groupedData[tanggal_bayar]; !ok {
// 			groupedData[tanggal_bayar] += total_amount

// 		} else {
// 			groupedData[tanggal_bayar] = total_amount
// 		}
// 	}
// 	return res, nil
// }
