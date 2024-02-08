package interfaces

import (
	"context"

	"github.com/ashiqsabith123/chat-svc/pkg/domain"
)

type UserRepo interface {
	CreateChatRoom(ctx context.Context, chat_room domain.ChatRooom) error
}
