package chatbotai

import (
	"context"
	"fmt"
	"testskripsi/module/entities"
	"time"

	"github.com/sashabaranov/go-openai"
	"go.mongodb.org/mongo-driver/mongo"
)

var (
	client *mongo.Client
)

var (
	debug = false
)

var chatCollection *mongo.Collection

func CollectionChat() {
	chatCollection = client.Database("chatbot").Collection("chats")
}

func GetCompletionFromMessage(client *openai.Client, ctx context.Context, messages []openai.ChatCompletionMessage) (openai.ChatCompletionResponse, error) {

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

func CreateQuestion(iduser uint64, name string, text string) (string, error) {
	ctx := context.Background()
	newChat := entities.ChatRequest{
		IdUser:    iduser,
		Role:      "question",
		Name:      name,
		Text:      text,
		CreatedAt: time.Now(),
	}

	chatCollection.InsertOne(ctx, newChat)
	return newChat.Text, nil
}

func CreateAnswer(iduser uint64, name string, text string) (string, error) {

	var apiKey string = "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"
	client := openai.NewClient(apiKey)
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

	if text != "" {
		message = append(message, openai.ChatCompletionMessage{
			Role:    openai.ChatMessageRoleUser,
			Content: text,
		})
	}
	resp, err := GetCompletionFromMessage(client, ctx, message)
	if err != nil {
		return "", err
	}
	if debug {
		fmt.Printf(
			"ID: %s. Created: %d. Model: %s. Choices: %v.\n",
			resp.ID, resp.Created, resp.Model, resp.Choices,
		)
	}

	answer := openai.ChatCompletionMessage{
		Role:    resp.Choices[0].Message.Role,
		Content: resp.Choices[0].Message.Content,
	}

	Assistant := entities.ChatRequest{
		IdUser:    iduser,
		Text:      answer.Content,
		CreatedAt: time.Now(),
		Role:      "answer",
		Name:      name,
	}

	chatCollection.InsertOne(ctx, Assistant)
	return answer.Content, nil
}
