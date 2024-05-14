package service

import (
	"testskripsi/module/entities"
	"testskripsi/module/feature/faq"
)

type FaqService struct {
	repo faq.FaqRepositoryInterface
}

func NewFaqService(repo faq.FaqRepositoryInterface) faq.FaqServiceInterface {
	return &FaqService{
		repo: repo,
	}
}

func (r *FaqService) CreateFaq(newData *entities.FaqModel) (*entities.FaqModel, error) {
	value := &entities.FaqModel{
		Question: newData.Question,
		Answer:   newData.Answer,
		View:     0,
	}
	res, err := r.repo.CreateFaq(value)
	if err != nil {
		return nil, err
	}
	return res, nil
}
func (r *FaqService) GetAllFaq() ([]*entities.FaqModel, error) {
	res, err := r.repo.GetAllFaq()
	if err != nil {
		return nil, err
	}
	return res, nil
}
func (r *FaqService) GetFaqById(id uint64) (*entities.FaqModel, error) {
	res, err := r.repo.GetFaqById(id)
	if err != nil {
		return nil, err
	}
	return res, nil
}
func (r *FaqService) DeleteFaqById(id uint64) error {
	err := r.repo.DeleteFaqById(id)
	if err != nil {
		return err
	}
	return nil

}
func (r *FaqService) UpdateFaqById(id uint64, newData *entities.FaqModel) (*entities.FaqModel, error) {
	value := &entities.FaqModel{
		Question: newData.Question,
		Answer:   newData.Answer,
		View:     0,
	}
	res, err := r.repo.UpdateFaqById(id, value)
	if err != nil {
		return nil, err
	}
	return res, nil
}
func (r *FaqService) IncrementViewsFaq(id uint64) error {
	err := r.repo.IncrementViewsFaq(id)
	if err != nil {
		return err
	}
	return nil
}
