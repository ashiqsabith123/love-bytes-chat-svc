package interfaces

import (
	"context"

	"github.com/ashiqsabith123/love-bytes-proto/chat/pb"
)

type UserUsecase interface {
	CreateChatRoom(ctx context.Context, req *pb.ChatRoomRequest) error
}
