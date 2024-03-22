package entities

import "time"

type AkunModel struct {
	ID        uint64            `gorm:"column:id;type:BIGINT UNSIGNED;primaryKey" json:"id"`
	Name      string            `gorm:"column:name;type:VARCHAR(255)" json:"name"`
	Email     string            `gorm:"column:email;type:VARCHAR(255);unique" json:"email"`
	Role      string            `gorm:"column:role;type:VARCHAR(20)" json:"role"`
	Password  string            `gorm:"column:password;type:VARCHAR(255)" json:"password"`
	Transaksi []TransaksiModels `gorm:"foreignKey:IDAkun" json:"transaksi"`
	Invoice   []InvoiceModels   `gorm:"foreignKey:IDAkun" json:"invoiice"`
}

func (AkunModel) TableName() string {
	return "akuns"
}

type UserModels struct {
	ID             uint64            `gorm:"column:id;type:BIGINT UNSIGNED;primaryKey" json:"id"`
	Name           string            `gorm:"column:name;type:VARCHAR(255)" json:"name"`
	Status         string            `gorm:"column:status;type:VARCHAR(100)"  json:"status"`
	Alamat         string            `gorm:"column:alamat;type:TEXT" json:"alamat"`
	PaketLangganan string            `gorm:"column:paket_langganan;type:VARCHAR(20)" json:"paket_langganan"`
	HargaLangganan uint64            `gorm:"column:harga_langganan;type:BIGINT(30)" json:"harga_langganan"`
	CreatedAt      time.Time         `gorm:"column:created_at;type:timestamp DEFAULT CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt      time.Time         `gorm:"column:updated_at;type:timestamp DEFAULT CURRENT_TIMESTAMP" json:"updated_at"`
	DeletedAt      *time.Time        `gorm:"column:deleted_at;type:TIMESTAMP NULL;index" json:"deleted_at"`
	Tagihan        []TagihanModels   `gorm:"foreignKey:IDPelanggan" json:"tagihan"`
	Transaksi      []TransaksiModels `gorm:"foreignKey:IDPelanggan" json:"transaksi"`
	Invoice        []InvoiceModels   `gorm:"foreignKey:IDPelanggan" json:"invoice"`
	SnapUrl        []SnapUrl         `gorm:"foreignKey:IDPelanggan" json:"snap_url"`
}
type TagihanModels struct {
	ID               uint64     `gorm:"column:id;type:BIGINT UNSIGNED;primaryKey" json:"id"`
	IDPelanggan      uint64     `gorm:"column:id_pelanggan;type:BIGINT UNSIGNED;" json:"id_pelanggan"`
	Name             string     `gorm:"column:name;type:VARCHAR(255)" json:"name"`
	Status           string     `gorm:"column:status;type:VARCHAR(100)"  json:"status"`
	Alamat           string     `gorm:"column:alamat;type:TEXT" json:"alamat"`
	PaketLangganan   string     `gorm:"column:paket_langganan;type:VARCHAR(20)" json:"paket_langganan"`
	HargaLangganan   uint64     `gorm:"column:harga_langganan;type:BIGINT(30)" json:"harga_langganan"`
	Bulan            string     `gorm:"column:bulan;type:VARCHAR(50)" json:"bulan"`
	Tahun            uint64     `gorm:"column:tahun;type:VARCHAR(50)" json:"tahun"`
	PeriodePemakaian string     `gorm:"column:periode_pemakaian;type:VARCHAR(255);unique" json:"periode_pemakaian"`
	CreatedAt        time.Time  `gorm:"column:created_at;type:timestamp DEFAULT CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt        time.Time  `gorm:"column:updated_at;type:timestamp DEFAULT CURRENT_TIMESTAMP" json:"updated_at"`
	DeletedAt        *time.Time `gorm:"column:deleted_at;type:TIMESTAMP NULL;index" json:"deleted_at"`
}
type TransaksiModels struct {
	ID               uint64     `gorm:"column:id;type:BIGINT UNSIGNED;primaryKey" json:"id"`
	IDPelanggan      uint64     `gorm:"column:id_pelanggan;type:BIGINT UNSIGNED;" json:"id_pelanggan"`
	IDAkun           uint64     `gorm:"column:id_akun;type:BIGINT UNSIGNED;" json:"id_akun"`
	IDPembayaran     string     `gorm:"column:id_pembayaran;type:VARCHAR(255);unique;" json:"id_pembayaran"`
	Email            string     `gorm:"column:email;type:VARCHAR(255)" json:"email"`
	Name             string     `gorm:"column:name;type:VARCHAR(255)" json:"name"`
	Status           string     `gorm:"column:status;type:VARCHAR(100)"  json:"status"`
	Alamat           string     `gorm:"column:alamat;type:TEXT" json:"alamat"`
	Url              string     `gorm:"column:url;type:TEXT" json:"url"`
	PeriodePemakaian string     `gorm:"column:periode_pemakaian;type:VARCHAR(255)" json:"periode_pemakaian"`
	PaketLangganan   string     `gorm:"column:paket_langganan;type:VARCHAR(20)" json:"paket_langganan"`
	HargaLangganan   uint64     `gorm:"column:harga_langganan;type:BIGINT(30)" json:"harga_langganan"`
	TotalAmount      uint64     `gorm:"column:total_amount;type:BIGINT(30)" json:"total_amount"`
	MetodePembayaran string     `gorm:"column:metode_pembayaran;type:VARCHAR(100)" json:"metode_pembayaran"`
	CreatedAt        time.Time  `gorm:"column:created_at;type:timestamp DEFAULT CURRENT_TIMESTAMP" json:"created_at"`
	TanggalBayar     time.Time  `gorm:"column:tanggal_bayar;type:DATE" json:"tanggal_bayar"`
	UpdatedAt        time.Time  `gorm:"column:updated_at;type:timestamp DEFAULT CURRENT_TIMESTAMP" json:"updated_at"`
	DeletedAt        *time.Time `gorm:"column:deleted_at;type:TIMESTAMP NULL;index" json:"deleted_at"`
}
type InvoiceModels struct {
	ID               uint64     `gorm:"column:id;type:BIGINT UNSIGNED;primaryKey" json:"id"`
	IDPelanggan      uint64     `gorm:"column:id_pelanggan;type:BIGINT UNSIGNED;" json:"id_pelanggan"`
	IDTransaksi      uint64     `gorm:"column:id_transaksi;type:BIGINT UNSIGNED;unique" json:"id_transaksi"`
	IDAkun           uint64     `gorm:"column:id_akun;type:BIGINT UNSIGNED;" json:"id_akun"`
	Email            string     `gorm:"column:email;type:VARCHAR(255)" json:"email"`
	Name             string     `gorm:"column:name;type:VARCHAR(255)" json:"name"`
	Status           string     `gorm:"column:status;type:VARCHAR(100)"  json:"status"`
	Alamat           string     `gorm:"column:alamat;type:TEXT" json:"alamat"`
	PaketLangganan   string     `gorm:"column:paket_langganan;type:VARCHAR(20)" json:"paket_langganan"`
	HargaLangganan   uint64     `gorm:"column:harga_langganan;type:BIGINT(30)" json:"harga_langganan"`
	Bulan            string     `gorm:"column:bulan;type:VARCHAR(50)" json:"bulan"`
	Tahun            uint64     `gorm:"column:tahun;type:VARCHAR(50)" json:"t ahun"`
	PeriodePemakaian string     `gorm:"column:periode_pemakaian;type:VARCHAR(255);unique" json:"periode_pemakaian"`
	TanggalBayar     time.Time  `gorm:"column:tanggal_bayar;type:DATE" json:"tanggal_bayar"`
	CreatedAt        time.Time  `gorm:"column:created_at;type:timestamp DEFAULT CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt        time.Time  `gorm:"column:updated_at;type:timestamp DEFAULT CURRENT_TIMESTAMP" json:"updated_at"`
	DeletedAt        *time.Time `gorm:"column:deleted_at;type:TIMESTAMP NULL;index" json:"deleted_at"`
}

type SnapUrl struct {
	ID          uint64     `gorm:"column:id;type:BIGINT UNSIGNED;primaryKey" json:"id"`
	IDPelanggan uint64     `gorm:"column:id_pelanggan;type:BIGINT UNSIGNED;" json:"id_pelanggan"`
	Amount      uint64     `gorm:"column:amount;INTEGER(255)" json:"amount"`
	Url         string     `gorm:"column:url;type:TEXT" json:"url"`
	CreatedAt   time.Time  `gorm:"column:created_at;type:timestamp DEFAULT CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt   time.Time  `gorm:"column:updated_at;type:timestamp DEFAULT CURRENT_TIMESTAMP" json:"updated_at"`
	DeletedAt   *time.Time `gorm:"column:deleted_at;type:TIMESTAMP NULL;index" json:"deleted_at"`
}

func (UserModels) TableName() string {
	return "users"
}
func (TagihanModels) TableName() string {
	return "tagihans"
}
func (TransaksiModels) TableName() string {
	return "transaksis"
}
func (InvoiceModels) TableName() string {
	return "invoices"
}
func (SnapUrl) TableName() string {
	return "snapurls"
}
