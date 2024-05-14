package faq

import (
	"testskripsi/module/entities"

	"github.com/labstack/echo/v4"
)

type FaqRepositoryInterface interface {
	CreateFaq(newData *entities.FaqModel) (*entities.FaqModel, error)
	GetAllFaq() ([]*entities.FaqModel, error)
	GetFaqById(id uint64) (*entities.FaqModel, error)
	DeleteFaqById(id uint64) error
	UpdateFaqById(id uint64, newData *entities.FaqModel) (*entities.FaqModel, error)
	IncrementViewsFaq(id uint64) error
}
type FaqServiceInterface interface {
	CreateFaq(newData *entities.FaqModel) (*entities.FaqModel, error)
	GetAllFaq() ([]*entities.FaqModel, error)
	GetFaqById(id uint64) (*entities.FaqModel, error)
	DeleteFaqById(id uint64) error
	UpdateFaqById(id uint64, newData *entities.FaqModel) (*entities.FaqModel, error)
	IncrementViewsFaq(id uint64) error
}

type FaqHandlerInterface interface {
	CreateFaq() echo.HandlerFunc
	GetAllFaq() echo.HandlerFunc
	GetFaqById() echo.HandlerFunc
	UpdateFaqById() echo.HandlerFunc
	IncrementViewsFaq() echo.HandlerFunc
	DeleteFaqById() echo.HandlerFunc
}
