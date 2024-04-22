package service

import (
	"context"
	"fmt"
	"testskripsi/config"
	"testskripsi/module/entities"
	"testskripsi/module/feature/chatbot"
	"time"

	"github.com/sashabaranov/go-openai"
)

type ChatService struct {
	repo  chatbot.ChatRepositoryInterface
	debug bool
}

func NewChatService(repo chatbot.ChatRepositoryInterface) chatbot.ChatServiceInterface {
	return &ChatService{
		repo:  repo,
		debug: false,
	}

}

func (s *ChatService) GetAnswerFromAi(client *openai.Client, messages []openai.ChatCompletionMessage, ctx context.Context) (openai.ChatCompletionResponse, error) {

	model := openai.GPT3Dot5Turbo
	resp, err := client.CreateChatCompletion(
		ctx,
		openai.ChatCompletionRequest{
			Model:    model,
			Messages: messages,
		},
	)

	return resp, err

}

func (s *ChatService) CreateQuestion(newData entities.Chat) (string, error) {
	value := &entities.Chat{
		IdUser:    newData.IdUser,
		Role:      "question",
		Text:      newData.Text,
		CreatedAt: time.Now(),
	}
	if err := s.repo.CreateQuestion(*value); err != nil {
		return "", err
	}
	return newData.Text, nil
}

func (s *ChatService) CreateAnswer(newData entities.Chat) (string, error) {

	client := openai.NewClient(config.InitConfig().OpenaiKey)
	ctx := context.Background()
	message := []openai.ChatCompletionMessage{
		{
			Role:    openai.ChatMessageRoleSystem,
			Content: "Kamu adalah chatbot yang berperan sebagai customer service PT Media Grasi Internet,dilarang menjawab selain pertanyaan tentang Perusahaan dan tentang keluhan pelanggan",
		},
		{
			Role:    openai.ChatMessageRoleUser,
			Content: "dimana alamat PT Media Grasi Internet ?",
		},
		{
			Role:    openai.ChatMessageRoleAssistant,
			Content: "ada di perumahan puri rajeg",
		},
		{
			Role:    openai.ChatMessageRoleUser,
			Content: "Perusahaan ini bergerak dibidang apa?",
		},
		{
			Role:    openai.ChatMessageRoleAssistant,
			Content: "Bergerak dibidang pelayanan internet seperti internet rumah, internet sekolah dan internet instansi",
		},
		{
			Role:    openai.ChatMessageRoleUser,
			Content: "Paket apa saja yang tersedia untuk rumahan",
		},
		{
			Role:    openai.ChatMessageRoleAssistant,
			Content: "ada paket 3mbps dengan harga Rp.133.300, paket 5mbps dengan harga Rp.166.500, paket 10mbps dengan harga Rp.200.000 perbulan",
		},
		{
			Role:    openai.ChatMessageRoleSystem,
			Content: "Kamu adalah chatbot yang berperan sebagai customer service PT Media Grasi Internet, Apabila terjadi pengaduan tentang internet bermasalah",
		},
		{
			Role:    openai.ChatMessageRoleUser,
			Content: "Haii, internet saya bermasalah, apa yang harus saya lakukan?",
		},
		{
			Role:    openai.ChatMessageRoleAssistant,
			Content: "untuk langkah pertama silahkan restart modem dengan cara mencabut colokan dan colok kembali sebanyak tiga kali, apabila tidak bisa masukan nama, alamat dan no telp anda untuk kami teruskan ke admin, untuk ditangani",
		},
		{
			Role:    openai.ChatMessageRoleUser,
			Content: "Haii, saya sudah melakukan pembayaran kenapa masih dimatikan?",
		},
		{
			Role:    openai.ChatMessageRoleAssistant,
			Content: "untuk langkah pertama silahkan restart modem dengan cara mencabut colokan dan colok kembali sebanyak tiga kali, apabila tidak bisa masukan nama, alamat dan no telp anda untuk kami teruskan ke admin, untuk ditangani",
		},
		{
			Role:    openai.ChatMessageRoleSystem,
			Content: "Kamu adalah chatbot yang berperan sebagai customer service PT Media Grasi Internet, Apabila pelanggan ingin melakukan pembayaran",
		},
		{
			Role:    openai.ChatMessageRoleUser,
			Content: "Haii, saya ingin melakukan pembayaran bulanan bagaimana caranya",
		},
		{
			Role:    openai.ChatMessageRoleAssistant,
			Content: "silahkan kunjungi website www.layanangrasinet.com, kemudian login menggunakan google lalu pilih menu pembayaran, ketik id pelanggan, lalu jika tagihan sudah muncul silahkan klik bayar",
		},
	}

	if newData.Text != "" {
		message = append(message, openai.ChatCompletionMessage{
			Role:    openai.ChatMessageRoleUser,
			Content: newData.Text,
		})
	}
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
		CreatedAt: time.Now(),
		Role:      "answer",
	}

	if err := s.repo.CreateAnswer(*Assistant); err != nil {
		return "", err
	}
	return answer.Content, nil
}
