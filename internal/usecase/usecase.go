package usecase

import (
	"github.com/Ilja-R/library-auth-service/internal/adapter/driven/broker"
	"github.com/Ilja-R/library-auth-service/internal/adapter/driven/dbstore"
	"github.com/Ilja-R/library-auth-service/internal/config"
	"github.com/Ilja-R/library-auth-service/internal/port/usecase"
	authenticate "github.com/Ilja-R/library-auth-service/internal/usecase/authenticator"
	usercreater "github.com/Ilja-R/library-auth-service/internal/usecase/user_creater"
)

type UseCases struct {
	UserCreater   usecase.UserCreater
	Authenticator usecase.Authenticate
}

func New(cfg config.Config, store *dbstore.DBStore, publisher *broker.MessagePublisher) *UseCases {
	return &UseCases{
		UserCreater:   usercreater.New(&cfg, store.UserStorage, publisher),
		Authenticator: authenticate.New(&cfg, store.UserStorage),
	}
}
