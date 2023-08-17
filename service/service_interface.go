package service

import (
	"context"
)

type UserService interface {
	FetchUsersFromProviders(ctx context.Context) error
}
