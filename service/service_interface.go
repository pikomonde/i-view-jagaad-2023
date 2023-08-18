//go:generate rm -fr mocks
//go:generate mockery --all --dir . --keeptree --with-expecter --output ./mocks
package service

import "i-view-jagaad-2023/model"

type UserService interface {
	FetchUsersFromProviders() error
	GetUserByTags(tags []string) ([]model.User, error)
}
