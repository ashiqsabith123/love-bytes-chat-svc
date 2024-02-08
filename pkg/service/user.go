package service

import (
	"context"
	"fmt"

	usecase "github.com/ashiqsabith123/chat-svc/pkg/usecases/interfaces"
	"github.com/ashiqsabith123/love-bytes-proto/chat/pb"
	"google.golang.org/protobuf/types/known/anypb"
)

type UserService struct {
	UserUsecase usecase.UserUsecase
	pb.UnimplementedChatServiceServer
}

func NewUserService(usecase usecase.UserUsecase) UserService {
	return UserService{UserUsecase: usecase}
}

func (U *UserService) CreateChatRoom(ctx context.Context, req *pb.ChatRoomRequest) (*pb.ChatResponse, error) {
	fmt.Println("hreeee")
	err := U.UserUsecase.CreateChatRoom(ctx, req)

	if err != nil {
		return &pb.ChatResponse{
			Code:    500,
			Message: "Server error",
			Error: &anypb.Any{
				Value: []byte(err.Error()),
			},
		}, nil
	}

	return &pb.ChatResponse{
		Code:    201,
		Message: "Chat room created succesfully",
	}, nil

}
