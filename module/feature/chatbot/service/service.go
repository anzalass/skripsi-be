package service

import (
	"context"
	"fmt"
	"testskripsi/config"
	"testskripsi/module/entities"
	"testskripsi/module/feature/chatbot"
	"testskripsi/module/feature/pelanggan"
	"time"

	"github.com/sashabaranov/go-openai"
)

type ChatService struct {
	repo     chatbot.ChatRepositoryInterface
	repouser pelanggan.RepositoryPelanggan
	debug    bool
}

func NewChatService(repo chatbot.ChatRepositoryInterface, repouser pelanggan.RepositoryPelanggan) chatbot.ChatServiceInterface {
	return &ChatService{
		repo:     repo,
		repouser: repouser,
		debug:    false,
	}

}

func (s *ChatService) GetAnswerFromAi(client *openai.Client, messages []openai.ChatCompletionMessage, ctx context.Context) (openai.ChatCompletionResponse, error) {

	model := openai.GPT3Dot5Turbo
	// resp, err := client.CreateFineTuningJob(
	// 	ctx,
	// 	openai.ChatCompletionRequest{
	// 		Model:    model,
	// 		Messages: messages,

	// 		// MaxTokens: 10,
	// 	},
	// )
	resp, err := client.CreateChatCompletion(
		ctx,
		openai.ChatCompletionRequest{
			Model:    model,
			Messages: messages,

			// MaxTokens: 10,
		},
	)

	return resp, err

}

func (s *ChatService) CreateQuestion(newData entities.Chat) (*entities.Chat, error) {

	idakun, err := s.repouser.GetIdAkunByEmail(newData.Email)
	if err != nil {
		return nil, err
	}
	iduser, _ := s.repouser.CheckIdUserByEmail(newData.Email)

	lengthchatperday, _ := s.repo.GetChatPerDay(iduser)

	if len(lengthchatperday) > 9 {
		return nil, err
	}

	value := &entities.Chat{
		IdAkun:    idakun,
		IdUser:    iduser,
		Role:      "question",
		Text:      newData.Text,
		Email:     newData.Email,
		Name:      newData.Name,
		Views:     0,
		CreatedAt: time.Now(),
	}
	res, err := s.repo.CreateQuestion(*value)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (s *ChatService) CreateAnswer(newData entities.Chat) (string, error) {
	finetune := Fine()
	client := openai.NewClient(config.InitConfig().OpenaiKey)
	ctx := context.Background()
	message := []openai.ChatCompletionMessage{
		{
			Role:    openai.ChatMessageRoleUser,
			Content: "Kamu adalah chatbot yang berperan sebagai customer service PT Media Grasi Internet,dilarang menjawab selain pertanyaan tentang Perusahaan dan tentang keluhan pelanggan",
		},
	}

	for _, traine := range finetune {
		message = append(message, openai.ChatCompletionMessage{
			Role:    openai.ChatMessageRoleSystem,
			Content: traine.System,
		})
		message = append(message, openai.ChatCompletionMessage{
			Role:    openai.ChatMessageRoleUser,
			Content: traine.User,
		})
		message = append(message, openai.ChatCompletionMessage{
			Role:    openai.ChatMessageRoleAssistant,
			Content: traine.Assistant,
		})
	}

	message = append(message, openai.ChatCompletionMessage{
		Role:    openai.ChatMessageRoleUser,
		Content: newData.Text,
	})

	resp, err := s.GetAnswerFromAi(client, message, ctx)
	if err != nil {
		return "", err
	}
	if s.debug {
		fmt.Printf(
			"ID: %s. Created: %d. Model: %s. Choices: %v.\n",
			resp.ID, resp.Created, resp.Model, resp.Choices,
		)
	}

	answer := openai.ChatCompletionMessage{
		Role:    resp.Choices[0].Message.Role,
		Content: resp.Choices[0].Message.Content,
	}

	Assistant := &entities.Chat{
		IdUser:    newData.IdUser,
		Text:      answer.Content,
		Email:     newData.Email,
		Views:     0,
		CreatedAt: time.Now(),
		Role:      "answer",
	}

	if err := s.repo.CreateAnswer(*Assistant); err != nil {
		return "", err
	}
	return answer.Content, nil
}

func (s *ChatService) GetChatByEmail(email string) ([]entities.Chat, error) {
	res, err := s.repo.GetChatByEmail(email)

	if err != nil {
		return nil, err
	}

	return res, nil
}
