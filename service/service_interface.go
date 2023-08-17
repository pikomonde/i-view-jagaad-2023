package service

import "i-view-jagaad-2023/model"

type UserService interface {
	FetchUsersFromProviders() error
	GetUserByTags(tags []string) ([]model.User, error)
}
