package repository

import (
	"testskripsi/module/entities"
	"testskripsi/module/feature/faq"

	"gorm.io/gorm"
)

type FaqRepository struct {
	db *gorm.DB
}

func NewFaqRepository(db *gorm.DB) faq.FaqRepositoryInterface {
	return &FaqRepository{
		db: db,
	}
}

func (r *FaqRepository) CreateFaq(newData *entities.FaqModel) (*entities.FaqModel, error) {
	if err := r.db.Create(&newData).Error; err != nil {
		return nil, err
	}
	return newData, nil
}
func (r *FaqRepository) GetAllFaq() ([]*entities.FaqModel, error) {
	var faq []*entities.FaqModel
	if err := r.db.Find(&faq).Error; err != nil {
		return nil, err
	}
	return faq, nil
}
func (r *FaqRepository) GetFaqById(id uint64) (*entities.FaqModel, error) {
	faq := &entities.FaqModel{}
	if err := r.db.Where("id = ?", id).First(faq).Error; err != nil {
		return nil, err
	}
	return faq, nil
}
func (r *FaqRepository) DeleteFaqById(id uint64) error {
	faq := &entities.FaqModel{}
	if err := r.db.Where("id =?", id).Delete(faq).Error; err != nil {
		return err
	}
	return nil
}
func (r *FaqRepository) UpdateFaqById(id uint64, newData *entities.FaqModel) (*entities.FaqModel, error) {
	faq := &entities.FaqModel{}
	if err := r.db.Model(faq).Where("id= ?", id).Updates(&newData).Error; err != nil {
		return nil, err
	}

	return faq, nil
}
func (r *FaqRepository) IncrementViewsFaq(id uint64) error {
	faq := &entities.FaqModel{}
	if err := r.db.Where("id = ?", id).First(faq).Error; err != nil {
		return err
	}

	if err := r.db.Model(faq).Where("id = ?", id).Update("view", faq.View+1).Error; err != nil {
		return err
	}

	return nil
}
