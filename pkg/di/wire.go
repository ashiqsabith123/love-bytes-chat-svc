//go:build wireinject
// +build wireinject

package di

import (
	"github.com/ashiqsabith123/chat-svc/pkg/config"
	"github.com/ashiqsabith123/chat-svc/pkg/db"
	"github.com/ashiqsabith123/chat-svc/pkg/repository"
	"github.com/ashiqsabith123/chat-svc/pkg/service"
	usecase "github.com/ashiqsabith123/chat-svc/pkg/usecases"
	"github.com/google/wire"
)

func IntializeService(config config.Config) service.UserService {

	wire.Build(
		db.ConnectToDatabase,
		repository.NewUserRepo,
		usecase.NewUserUsecase,
		service.NewUserService,
	)

	return service.UserService{}

}
