package usecases

import (
	"context"

	"github.com/ashiqsabith123/chat-svc/pkg/domain"
	repo "github.com/ashiqsabith123/chat-svc/pkg/repository/interfaces"
	usecase "github.com/ashiqsabith123/chat-svc/pkg/usecases/interfaces"
	"github.com/ashiqsabith123/love-bytes-proto/chat/pb"
)

type UserUsecase struct {
	UserRepo repo.UserRepo
}

func NewUserUsecase(repo repo.UserRepo) usecase.UserUsecase {
	return &UserUsecase{UserRepo: repo}
}

func (U *UserUsecase) CreateChatRoom(ctx context.Context, req *pb.ChatRoomRequest) error {

	var chat_room domain.ChatRooom

	chat_room.User1ID = uint(req.UserID1)
	chat_room.User2ID = uint(req.UserID2)

	err := U.UserRepo.CreateChatRoom(ctx, chat_room)

	if err != nil {
		return err
	}

	return nil
}
