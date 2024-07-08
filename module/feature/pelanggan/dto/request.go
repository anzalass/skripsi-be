package dto

type CreatePelangganRequest struct {
	ID             string `json:"id" validate:"required,noSpace"`
	Name           string `json:"name" validate:"required"`
	Status         string `json:"status"`
	Alamat         string `json:"alamat" validate:"required"`
	PaketLangganan string `json:"paket_langganan" validate:"required"`
	HargaLangganan uint64 `json:"harga_langganan" validate:"required"`
}

type InsertIDAkun struct {
	Email  string `json:"email" validate:"required"`
	IdUser string `json:"id_user" validate:"required"`
}
