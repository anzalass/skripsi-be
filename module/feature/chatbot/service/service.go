package service

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
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
	nowa, _ := s.repouser.GetNoWhatsApp(iduser)

	lengthchatperday, _ := s.repo.GetChatPerDay(iduser)

	if len(lengthchatperday) > 9 {
		return nil, err
	}

	value := &entities.Chat{
		IdAkun:     idakun,
		IdUser:     iduser,
		Role:       "question",
		Text:       newData.Text,
		NoWhatsapp: nowa,
		Email:      newData.Email,
		Name:       newData.Name,
		Views:      0,
		CreatedAt:  time.Now(),
	}
	res, err := s.repo.CreateQuestion(*value)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (s *ChatService) CreateAnswer(newData entities.Chat) (string, error) {
	//finetune := Fine()
	client := openai.NewClient(config.InitConfig().OpenaiKey)
	//dataset, _ := s.repo.GetAllDataset()
	ctx := context.Background()
	message := []openai.ChatCompletionMessage{
		{
			Role:    openai.ChatMessageRoleUser,
			Content: "Kamu adalah chatbot yang berperan sebagai customer service PT Media Grasi Internet,dilarang menjawab selain pertanyaan tentang Perusahaan dan tentang keluhan pelanggan",
		},
	}

	// for _, traine := range dataset {
	// 	message = append(message, openai.ChatCompletionMessage{
	// 		Role:    openai.ChatMessageRoleSystem,
	// 		Content: traine.System,
	// 	})
	// 	message = append(message, openai.ChatCompletionMessage{
	// 		Role:    openai.ChatMessageRoleUser,
	// 		Content: traine.User,
	// 	})
	// 	message = append(message, openai.ChatCompletionMessage{
	// 		Role:    openai.ChatMessageRoleAssistant,
	// 		Content: traine.Assistant,
	// 	})
	// }

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
func (s *ChatService) GetAllChat() ([]entities.Chat, error) {
	res, err := s.repo.GetAllChat()

	if err != nil {
		return nil, err
	}

	return res, nil
}

func (s *ChatService) CreateDataset(newData *entities.DatasetAi) (*entities.DatasetAi, error) {
	value := &entities.DatasetAi{

		Role:    newData.Role,
		Content: newData.Content,
		Tipe:    newData.Tipe,
	}

	res, err := s.repo.CreateDataset(value)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (s *ChatService) UpdateDatasetById(id int, newData *entities.DatasetAi) (*entities.DatasetAi, error) {
	value := &entities.DatasetAi{
		Content: newData.Content,
	}

	res, err := s.repo.UpdateDatasetById(id, value)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (s *ChatService) DeleteDatasetById(id int) error {
	err := s.repo.DeleteDatasetById(id)
	if err != nil {
		return err
	}
	return nil
}
func (s *ChatService) GetDatasetById(id int) (*entities.DatasetAi, *entities.DatasetAi, error) {
	res, err := s.repo.GetDatasetById(id)
	if err != nil {
		return nil, nil, err
	}

	res2, err := s.repo.GetDatasetById(id - 1)
	if err != nil {
		return nil, nil, err
	}

	return res, res2, nil
}

func (s *ChatService) GetAllDataset() ([]*entities.DatasetAi, error) {
	res, err := s.repo.GetAllDataset()
	if err != nil {
		return nil, err
	}

	return res, nil
}

type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type DataRequest struct {
	Messages    []Message   `json:"messages"`
	Model       string      `json:"model"`
	Temperature int         `json:"temperature"`
	MaxTokens   int         `json:"max_tokens"`
	TopP        int         `json:"top_p"`
	Stream      bool        `json:"stream"`
	Stop        interface{} `json:"stop"`
}

type Response struct {
	Choices []Choice `json:"choices"`
}

type Choice struct {
	Message Message `json:"message"`
}

func (s *ChatService) GroqAi(newData entities.Chat) (any, error) {
	// tuning := Fine2()
	tuning2, err := s.repo.GetAllDataset()
	if err != nil {
		return nil, err
	}
	data := DataRequest{
		Messages: []Message{
			{
				Role:    "system",
				Content: "Kamu adalah chatbot yang berperan sebagai customer service PT Media Grasi Internet,dilarang menjawab selain pertanyaan tentang Perusahaan dan tentang keluhan pelanggan, dan jawablah menggunakan bahasa indonesia ",
			},
		},
		Model:       "llama3-8b-8192",
		Temperature: 1,
		MaxTokens:   500,
		TopP:        1,
		Stream:      false,
		Stop:        nil,
	}

	for _, traine := range tuning2 {
		data.Messages = append(data.Messages, Message{
			Role:    traine.Role,
			Content: traine.Content,
		})
	}

	data.Messages = append(data.Messages, Message{
		Role:    "user",
		Content: newData.Text,
	})

	jsonData, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	url := "https://api.groq.com/openai/v1/chat/completions"
	req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s ", config.InitConfig().GroqApiKey))

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result Response
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	Assistant := &entities.Chat{
		IdUser:    newData.IdUser,
		Text:      result.Choices[0].Message.Content,
		Email:     newData.Email,
		Views:     0,
		CreatedAt: time.Now(),
		Role:      "answer",
	}

	if err := s.repo.CreateAnswer(*Assistant); err != nil {
		return "", err
	}
	return result, nil

	// return result.Choices[0].Message.Content , nil
	// var response Response
	// err1 := json.Unmarshal([]byte(result), &response)
}
