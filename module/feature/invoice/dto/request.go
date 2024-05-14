package dto

type CreateInvoiceRequest struct {
	Bulan     string `json:"bulan"`
	Tahun     uint64 `json:"tahun"`
	Pemakaian string `json:"pemakaian"`
}

type TransaksiRequest struct {
	IDPelanggan      string `json:"id_pelanggan"`
	IDAkun           uint64 `json:"id_akun"`
	Email            string `json:"email"`
	Name             string `json:"name"`
	Alamat           string `json:"alamat"`
	PaketLangganan   string `json:"paket_langganan"`
	PeriodePemakaian string `json:"periode_pemakaian"`
	HargaLangganan   uint64 `json:"harga_langganan"`
	TotalAmount      uint64 `json:"total_amount"`
}
