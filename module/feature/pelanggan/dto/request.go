package dto

type CreatePelangganRequest struct {
	ID             string `json:"id"`
	Name           string `json:"name"`
	Status         string `json:"status"`
	Alamat         string `json:"alamat"`
	PaketLangganan string `json:"paket_langganan"`
	HargaLangganan uint64 `json:"harga_langganan"`
}

type InsertIDAkun struct {
	Email  string `json:"email"`
	IdUser string `json:"id_user"`
}
