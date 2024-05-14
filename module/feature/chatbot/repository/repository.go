package repository

import (
	"context"
	"testskripsi/module/entities"
	"testskripsi/module/feature/chatbot"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"gopkg.in/mgo.v2/bson"
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

func (r *ChatbotRepository) CreateQuestion(newData entities.Chat) (*entities.Chat, error) {
	ctx := context.Background()
	_, err := r.collection.InsertOne(ctx, newData)
	if err != nil {
		return nil, err
	}

	return &newData, nil
}
func (r *ChatbotRepository) CreateAnswer(newData entities.Chat) error {
	ctx := context.Background()
	if _, err := r.collection.InsertOne(ctx, newData); err != nil {
		return nil
	}
	return nil
}
func (r *ChatbotRepository) GetChatByEmail(email string) ([]entities.Chat, error) {
	ctx := context.Background()
	var datas []entities.Chat
	cursor, err := r.collection.Find(ctx, bson.M{"email": email})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		var data entities.Chat
		if err := cursor.Decode(&data); err != nil {
			return nil, err
		}
		datas = append(datas, data)
	}

	return datas, nil
}

func (r *ChatbotRepository) GetChatPerDay(iduser string) ([]entities.Chat, error) {
	ctx := context.Background()
	var datas []entities.Chat
	now := time.Now()
	startDate := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())
	endDate := startDate.Add(24 * time.Hour)
	filter := bson.M{
		"createdat": bson.M{
			"$gte": startDate,
			"$lt":  endDate,
		},
	}
	cursor, err := r.collection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	for cursor.Next(ctx) {
		var data entities.Chat
		if err := cursor.Decode(&data); err != nil {
			return nil, err
		}
		datas = append(datas, data)
	}
	return datas, nil
}
