package user

import (
	"i-view-jagaad-2023/model"
	"i-view-jagaad-2023/repository"
	"i-view-jagaad-2023/service"
)

type User struct {
	ProviderRepos []repository.UserProviderRepository
	FileRepo      repository.FileRepository
}

func NewService(
	providerRepos []repository.UserProviderRepository,
	fileRepo repository.FileRepository,
) service.UserService {
	return &User{
		ProviderRepos: providerRepos,
		FileRepo:      fileRepo,
	}
}

func (s *User) FetchUsersFromProviders() error {
	allUsers := make([]model.User, 0)

	// Fetch all users
	for _, providerRepo := range s.ProviderRepos {
		users, _ := providerRepo.FetchUsersFromProvider()
		allUsers = append(allUsers, users...)
	}

	// Save all users to csv file
	err := s.FileRepo.SaveUsers(allUsers)

	return err
}

func (s *User) GetUserByTags(tags []string) ([]model.User, error) {
	return nil, nil
}
