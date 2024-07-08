package database

import (
	"testskripsi/module/entities"

	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) {
	err := db.AutoMigrate(entities.UserModels{}, entities.AkunModel{}, entities.TagihanModels{}, entities.TransaksiModels{}, entities.InvoiceModels{}, entities.SnapUrl{}, entities.PengaduanModel{}, entities.FaqModel{}, entities.DatasetAi{})
	if err != nil {
		return
	}
}
