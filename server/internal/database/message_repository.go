package database

import (
	"context"
	domain "reezanvisramportfolio/domain/message"

	"go.mongodb.org/mongo-driver/mongo"
)

type MessageRepository interface {
	InsertMessage(ctx context.Context, message domain.Message) error
}

type messageRepository struct {
	messageCollection *mongo.Collection
}

func NewMessageRepository(messageCollection *mongo.Collection) MessageRepository {
	return &messageRepository{
		messageCollection: messageCollection,
	}
}

func (mr *messageRepository) InsertMessage(ctx context.Context, message domain.Message) error {
	_, err := mr.messageCollection.InsertOne(ctx, message)

	return err
}
