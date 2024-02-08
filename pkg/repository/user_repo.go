package repository

import (
	"context"

	"github.com/ashiqsabith123/chat-svc/pkg/db"
	"github.com/ashiqsabith123/chat-svc/pkg/domain"
	"github.com/ashiqsabith123/chat-svc/pkg/repository/interfaces"
)

type UserRepo struct {
	Mongo *db.Collections
}

func NewUserRepo(db *db.Collections) interfaces.UserRepo {
	return &UserRepo{Mongo: db}
}

func (U *UserRepo) CreateChatRoom(ctx context.Context, chat_room domain.ChatRooom) error {

	_, err := U.Mongo.Chat_rooms.InsertOne(ctx, chat_room)

	if err != nil {
		return err
	}

	return nil

}
