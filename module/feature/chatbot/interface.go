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
}
type ChatServiceInterface interface {
	CreateQuestion(newData entities.Chat) (*entities.Chat, error)
	CreateAnswer(newData entities.Chat) (string, error)
	GetAnswerFromAi(client *openai.Client, chat []openai.ChatCompletionMessage, ctx context.Context) (openai.ChatCompletionResponse, error)
	GetChatByEmail(email string) ([]entities.Chat, error)
}
type ChatHandlerInterface interface {
	CreateQuestion() echo.HandlerFunc
	CreateAnswer() echo.HandlerFunc
	GetChatByEmail() echo.HandlerFunc
}
