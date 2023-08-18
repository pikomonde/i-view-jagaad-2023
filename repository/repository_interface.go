//go:generate rm -fr mocks
//go:generate mockery --all --dir . --keeptree --with-expecter --output ./mocks
package repository

import (
	"i-view-jagaad-2023/model"
)

type UserProviderRepository interface {
	FetchUsersFromProvider() ([]model.User, error)
}

type FileRepository interface {
	SaveUsers(users []model.User) error
	GetUsers() ([]model.User, error)
}
