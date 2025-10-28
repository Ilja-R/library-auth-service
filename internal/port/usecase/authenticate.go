package usecase

import (
	"context"

	"github.com/Ilja-R/library-auth-service/internal/domain"
)

type Authenticate interface {
	Authenticate(ctx context.Context, user domain.User) (int, domain.Role, error)
}
