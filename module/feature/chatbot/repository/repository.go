package repository

import (
	"context"
	"testskripsi/module/entities"
	"testskripsi/module/feature/chatbot"

	"go.mongodb.org/mongo-driver/mongo"
)

type ChatbotRepository struct {
	collection *mongo.Collection
}

func NewChatbotRepository(db *mongo.Client) chatbot.ChatRepositoryInterface {
	collection := db.Database("GrasiNet").Collection("chats")
	return &ChatbotRepository{
		collection: collection,
	}
}

func (r *ChatbotRepository) CreateQuestion(newData entities.Chat) error {
	ctx := context.Background()
	if _, err := r.collection.InsertOne(ctx, newData); err != nil {
		return nil
	}
	return nil
}
func (r *ChatbotRepository) CreateAnswer(newData entities.Chat) error {
	ctx := context.Background()
	if _, err := r.collection.InsertOne(ctx, newData); err != nil {
		return nil
	}
	return nil
}
