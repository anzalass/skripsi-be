package dto

type EditStatusPengaduan struct {
	Status string `json:"status"`
}

type CreateSPengaduan struct {
	IDPelanggan    uint64 `json:"id_pelanggan" form:"id_pelanggan"`
	Nama           string `json:"nama" form:"nama"`
	Email          string `json:"email" form:"email"`
	Deskripsi      string `json:"deskripsi" form:"deskripsi"`
	Alamat         string `json:"alamat" form:"alamat"`
	NoWhatsapp     string `json:"no_whatsapp" form:"no_whatsapp"`
	WaktuKunjungan string `json:"waktu_kunjungan" form:"waktu_kunjungan"`
}
