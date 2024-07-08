package entities

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type AkunModel struct {
	ID     uint64 `gorm:"column:id;type:BIGINT UNSIGNED;primaryKey" json:"id"`
	IDUser string `gorm:"column:id_user;type:VARCHAR(255);unique" json:"id_user"`
	// User      UserModels        `gorm:"foreignKey:IDUser;references:ID" json:"user"`
	Name      string            `gorm:"column:name;type:VARCHAR(255)" json:"name"`
	Email     string            `gorm:"column:email;type:VARCHAR(255);unique" json:"email"`
	Role      string            `gorm:"column:role;type:VARCHAR(20)" json:"role"`
	Password  string            `gorm:"column:password;type:VARCHAR(255)" json:"password"`
	Transaksi []TransaksiModels `gorm:"foreignKey:IDAkun" json:"transaksi"`
	Invoice   []InvoiceModels   `gorm:"foreignKey:IDAkun" json:"invoice"`
	Pengaduan []PengaduanModel  `gorm:"foreignKey:IDAkun" json:"pengaduan"`
}

type UserModels struct {
	ID             string            `gorm:"column:id;type:VARCHAR(255);primaryKey" json:"id"`
	Name           string            `gorm:"column:name;type:VARCHAR(255)" json:"name"`
	Status         string            `gorm:"column:status;type:VARCHAR(100)"  json:"status"`
	Alamat         string            `gorm:"column:alamat;type:TEXT" json:"alamat"`
	PaketLangganan string            `gorm:"column:paket_langganan;type:VARCHAR(20)" json:"paket_langganan"`
	HargaLangganan uint64            `gorm:"column:harga_langganan;type:BIGINT(30)" json:"harga_langganan"`
	NoWhatsapp     string            `gorm:"column:no_whatsapp;type:VARCHAR(15);" json:"no_whatsapp"`
	CreatedAt      time.Time         `gorm:"column:created_at;type:timestamp DEFAULT CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt      time.Time         `gorm:"column:updated_at;type:timestamp DEFAULT CURRENT_TIMESTAMP" json:"updated_at"`
	DeletedAt      *time.Time        `gorm:"column:deleted_at;type:TIMESTAMP NULL;index" json:"deleted_at"`
	Tagihan        []TagihanModels   `gorm:"foreignKey:IDPelanggan" json:"tagihan"`
	Transaksi      []TransaksiModels `gorm:"foreignKey:IDPelanggan" json:"transaksi"`
	Invoice        []InvoiceModels   `gorm:"foreignKey:IDPelanggan" json:"invoice"`
	SnapUrl        []SnapUrl         `gorm:"foreignKey:IDPelanggan" json:"snap_url"`
	Pengaduan      []PengaduanModel  `gorm:"foreignKey:IDPelanggan" json:"pengaduan"`
}
type TagihanModels struct {
	ID               uint64     `gorm:"column:id;type:BIGINT UNSIGNED;primaryKey" json:"id"`
	IDPelanggan      string     `gorm:"column:id_pelanggan;type:VARCHAR(255);" json:"id_pelanggan"`
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
	IDPelanggan      string     `gorm:"column:id_pelanggan;type:VARCHAR(255);" json:"id_pelanggan"`
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
	Expired          uint64     `gorm:"column:expired;type:BIGINT(50)" json:"expired"`
	CreatedAt        time.Time  `gorm:"column:created_at;type:timestamp DEFAULT CURRENT_TIMESTAMP" json:"created_at"`
	TanggalBayar     time.Time  `gorm:"column:tanggal_bayar;type:DATE" json:"tanggal_bayar"`
	UpdatedAt        time.Time  `gorm:"column:updated_at;type:timestamp DEFAULT CURRENT_TIMESTAMP" json:"updated_at"`
	DeletedAt        *time.Time `gorm:"column:deleted_at;type:TIMESTAMP NULL;index" json:"deleted_at"`
}
type InvoiceModels struct {
	ID               uint64     `gorm:"column:id;type:BIGINT UNSIGNED;primaryKey" json:"id"`
	IDPelanggan      string     `gorm:"column:id_pelanggan;type:VARCHAR(255);" json:"id_pelanggan"`
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
	IDPelanggan string     `gorm:"column:id_pelanggan;type:VARCHAR(255);" json:"id_pelanggan"`
	Amount      uint64     `gorm:"column:amount;INTEGER(255)" json:"amount"`
	Url         string     `gorm:"column:url;type:TEXT" json:"url"`
	Expired     uint64     `gorm:"column:expired:INTEGER(255)" json:"expired"`
	CreatedAt   time.Time  `gorm:"column:created_at;type:timestamp DEFAULT CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt   time.Time  `gorm:"column:updated_at;type:timestamp DEFAULT CURRENT_TIMESTAMP" json:"updated_at"`
	DeletedAt   *time.Time `gorm:"column:deleted_at;type:TIMESTAMP NULL;index" json:"deleted_at"`
}

type Chat struct {
	ID         primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	IdAkun     uint64             `json:"id_akun" form:"id_akun"`
	IdUser     string             `json:"id_user" form:"id_user"`
	NoWhatsapp string             `json:"no_whatsapp" form:"no_whatsapp"`
	Email      string             `json:"email" form:"email"`
	Role       string             `json:"role" form:"role"`
	Views      uint64             `json:"views" form:"views"`
	Name       string             `json:"name" form:"name"`
	Text       string             `json:"text" form:"text"`
	CreatedAt  time.Time          `json:"created_at" form:"created_at"`
}
type ChatRequest struct {
	ID        primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	IdUser    uint64             `json:"id_user" form:"id_user"`
	Email     string             `json:"email" form:"email"`
	Role      string             `json:"role" form:"role"`
	Views     uint64             `json:"views" form:"views"`
	Name      string             `json:"name" form:"name"`
	Text      string             `json:"text" form:"text"`
	CreatedAt time.Time          `json:"created_at" form:"created_at"`
}

type PengaduanModel struct {
	ID             uint64     `gorm:"column:id;type:BIGINT UNSIGNED;primaryKey" json:"id"`
	IDPelanggan    string     `gorm:"column:id_pelanggan;type:VARCHAR(255);" json:"id_pelanggan"`
	Email          string     `gorm:"column:email;type:VARCHAR(255)" json:"email"`
	IDAkun         uint64     `gorm:"column:id_akun;type:BIGINT UNSIGNED;" json:"id_akun"`
	Nama           string     `gorm:"column:nama;type:VARCHAR(255);" json:"nama"`
	Deskripsi      string     `gorm:"column:deskripsi;type:TEXT;" json:"deskripsi"`
	Alamat         string     `gorm:"column:alamat;type:TEXT;" json:"alamat"`
	Foto           string     `gorm:"column:foto;type:TEXT" json:"foto"`
	NoWhatsapp     string     `gorm:"column:no_whatsapp;type:VARCHAR(15);" json:"no_whatsapp"`
	Status         string     `gorm:"column:status;type:VARCHAR(50);" json:"status"`
	WaktuKunjungan string     `gorm:"column:waktu_kunjungan;type:VARCHAR(255); default:'Sekarang'" json:"waktu_kunjungan"`
	CreatedAt      time.Time  `gorm:"column:created_at;type:timestamp DEFAULT CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt      time.Time  `gorm:"column:updated_at;type:timestamp DEFAULT CURRENT_TIMESTAMP" json:"updated_at"`
	DeletedAt      *time.Time `gorm:"column:deleted_at;type:TIMESTAMP NULL;index" json:"deleted_at"`
}

type FaqModel struct {
	ID        uint64     `gorm:"column:id;type:BIGINT UNSIGNED;primaryKey" json:"id"`
	Question  string     `gorm:"column:question;type:TEXT;" json:"question" form:"question"`
	Answer    string     `gorm:"column:answer;type:TEXT;" json:"answer" form:"answer"`
	View      uint64     `gorm:"column:view;type:BIGINT UNSIGNED;" json:"view" form:"view"`
	CreatedAt time.Time  `gorm:"column:created_at;type:timestamp DEFAULT CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time  `gorm:"column:updated_at;type:timestamp DEFAULT CURRENT_TIMESTAMP" json:"updated_at"`
	DeletedAt *time.Time `gorm:"column:deleted_at;type:TIMESTAMP NULL;index" json:"deleted_at"`
}

type DatasetAi struct {
	ID      uint64 `gorm:"column:id;type:BIGINT UNSIGNED;primaryKey" json:"id" `
	Role    string `gorm:"column:role;type:TEXT" json:"role"`
	Content string `gorm:"column:content;type:TEXT" json:"content" validate:"required"`
	Tipe    string `gorm:"column:tipe;type:VARCHAR(50)" json:"tipe"`
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
func (PengaduanModel) TableName() string {
	return "pengaduans"
}

func (AkunModel) TableName() string {
	return "akuns"
}

func (FaqModel) TableName() string {
	return "faqs"
}

func (DatasetAi) TableName() string {
	return "datasets"
}
