package chatbot

import (
	"context"
	"testskripsi/module/entities"

	"github.com/labstack/echo/v4"
	"github.com/sashabaranov/go-openai"
)

type ChatRepositoryInterface interface {
	CreateQuestion(newData entities.Chat) (*entities.Chat, error)
	CreateAnswer(newData entities.Chat) error
	GetChatByEmail(email string) ([]entities.Chat, error)
	GetChatPerDay(iduser string) ([]entities.Chat, error)
	GetAllChat() ([]entities.Chat, error)
	CreateDataset(newData *entities.DatasetAi) (*entities.DatasetAi, error)
	GetAllDataset() ([]*entities.DatasetAi, error)
	DeleteDatasetById(id int) error
	UpdateDatasetById(id int, newData *entities.DatasetAi) (*entities.DatasetAi, error)
	GetDatasetById(id int) (*entities.DatasetAi, error)
}
type ChatServiceInterface interface {
	CreateQuestion(newData entities.Chat) (*entities.Chat, error)
	GroqAi(newData entities.Chat) (any, error)
	CreateAnswer(newData entities.Chat) (string, error)
	GetAnswerFromAi(client *openai.Client, chat []openai.ChatCompletionMessage, ctx context.Context) (openai.ChatCompletionResponse, error)
	GetChatByEmail(email string) ([]entities.Chat, error)
	GetAllChat() ([]entities.Chat, error)
	CreateDataset(newData *entities.DatasetAi) (*entities.DatasetAi, error)
	GetAllDataset() ([]*entities.DatasetAi, error)
	UpdateDatasetById(id int, newData *entities.DatasetAi) (*entities.DatasetAi, error)
	DeleteDatasetById(id int) error
	GetDatasetById(id int) (*entities.DatasetAi, *entities.DatasetAi, error)
}
type ChatHandlerInterface interface {
	CreateQuestion() echo.HandlerFunc
	CreateAnswer() echo.HandlerFunc
	GetChatByEmail() echo.HandlerFunc
	CreateDataset() echo.HandlerFunc
	GetAllDataset() echo.HandlerFunc
	DeleteDatasetById() echo.HandlerFunc
	UpdateDatasetById() echo.HandlerFunc
	GetDatasetById() echo.HandlerFunc
	GetAllChat() echo.HandlerFunc
	GroqAi() echo.HandlerFunc
	GroqAi2() echo.HandlerFunc
}
