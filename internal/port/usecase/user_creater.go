package usecase

import (
	"context"

	"github.com/Ilja-R/library-auth-service/internal/domain"
)

type UserCreater interface {
	CreateUser(ctx context.Context, user domain.User) (err error)
}
