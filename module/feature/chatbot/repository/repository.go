package repository

import (
	"context"
	"testskripsi/module/entities"
	"testskripsi/module/feature/chatbot"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"gopkg.in/mgo.v2/bson"
	"gorm.io/gorm"
)

type ChatbotRepository struct {
	collection *mongo.Collection
	mysql      *gorm.DB
}

func NewChatbotRepository(db *mongo.Client, mysql *gorm.DB) chatbot.ChatRepositoryInterface {
	collection := db.Database("GrasiNet").Collection("chats")
	return &ChatbotRepository{
		collection: collection,
		mysql:      mysql,
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
func (r *ChatbotRepository) GetAllChat() ([]entities.Chat, error) {
	ctx := context.Background()
	var datas []entities.Chat

	// Mendapatkan waktu awal dan akhir hari ini
	now := time.Now()
	startOfDay := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())
	endOfDay := time.Date(now.Year(), now.Month(), now.Day(), 23, 59, 59, 999000000, now.Location()) // Menambahkan 999ms untuk mendekati pukul 24:00

	filter := bson.M{
		"role": "question",
		"createdat": bson.M{
			"$gte": startOfDay,
			"$lt":  endOfDay,
		},
	}

	// Mengatur opsi pencarian untuk mengurutkan berdasarkan 'created_at' dalam urutan naik (ascending)

	cursor, err := r.collection.Find(ctx, filter)
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
		"iduser": iduser,
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

func (r *ChatbotRepository) CreateDataset(newData *entities.DatasetAi) (*entities.DatasetAi, error) {
	if err := r.mysql.Create(&newData).Error; err != nil {
		return nil, err
	}
	return newData, nil
}
func (r *ChatbotRepository) UpdateDatasetById(id int, newData *entities.DatasetAi) (*entities.DatasetAi, error) {
	dataset := &entities.DatasetAi{}
	if err := r.mysql.Model(dataset).Where("id = ?", id).Updates(&newData).Error; err != nil {
		return nil, err
	}
	return newData, nil
}
func (r *ChatbotRepository) DeleteDatasetById(id int) error {
	dataset := &entities.DatasetAi{}
	if err := r.mysql.Where("id=?", id).Delete(dataset).Error; err != nil {
		return err
	}
	return nil
}
func (r *ChatbotRepository) GetAllDataset() ([]*entities.DatasetAi, error) {
	var dataset []*entities.DatasetAi
	if err := r.mysql.Find(&dataset).Error; err != nil {
		return nil, err
	}
	return dataset, nil
}

func (r *ChatbotRepository) GetDatasetById(id int) (*entities.DatasetAi, error) {
	dataset := &entities.DatasetAi{}
	if err := r.mysql.Where("id=?", id).First(dataset).Error; err != nil {
		return nil, err
	}
	return dataset, nil
}
