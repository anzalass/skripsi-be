package repository

import (
	"testskripsi/module/entities"
	"testskripsi/module/feature/auth"

	"gorm.io/gorm"
)

type AuthRepository struct {
	db *gorm.DB
}

func NewAuthRepository(db *gorm.DB) auth.RepositoryAuthInterface {
	return &AuthRepository{
		db: db,
	}
}

func (r *AuthRepository) Register(newData *entities.AkunModel) (*entities.AkunModel, error) {
	if err := r.db.Create(newData).Error; err != nil {
		return nil, err
	}

	return newData, nil
}

func (r *AuthRepository) LoginAdmin(email string) (*entities.AkunModel, error) {
	var user entities.AkunModel
	if err := r.db.Where("email = ? && role = ?", email, "admin").First(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *AuthRepository) CekEmail(email string) (bool, error) {
	if err := r.db.Where("email = ?", email).Error; err != nil {
		return false, err
	}

	return true, nil
}
