package repository

import (
	"context"
	"i-view-jagaad-2023/model"
)

type UserProviderRepository interface {
	FetchUsersFromProvider(ctx context.Context) ([]model.User, error)
}

type FileRepository interface {
	SaveUsers(users []model.User) error
	GetUsers() ([]model.User, error)
}
