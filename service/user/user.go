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

func (s *User) GetUserByTags(inputTags []string) ([]model.User, error) {
	users, _ := s.FileRepo.GetUsers()

	selectedUsers := make([]model.User, 0)
	for _, user := range users {
		userTagMap := make(map[string]bool)
		for _, userTag := range user.Tags {
			userTagMap[userTag] = true
		}

		allTagExist := true
		if len(inputTags) == 0 {
			allTagExist = false
		}
		for _, inputTag := range inputTags {
			if _, exist := userTagMap[inputTag]; !exist {
				allTagExist = false
				break
			}
		}

		if allTagExist {
			selectedUsers = append(selectedUsers, user)
		}
	}

	return selectedUsers, nil
}
