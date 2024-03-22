package dto

type CreatePelangganRequest struct {
	Name           string `json:"name"`
	Status         string `json:"status"`
	Alamat         string `json:"alamat"`
	PaketLangganan string `json:"paket_langganan"`
	HargaLangganan uint64 `json:"harga_langganan"`
}
