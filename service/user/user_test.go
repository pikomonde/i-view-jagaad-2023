package user_test

import (
	"i-view-jagaad-2023/model"
	"i-view-jagaad-2023/repository"
	mocks "i-view-jagaad-2023/repository/mocks"
	serviceUser "i-view-jagaad-2023/service/user"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUserService_FetchFromProviders_NoUser_Success(t *testing.T) {
	mockUserProviderRepository := mocks.NewUserProviderRepository(t)
	mockUserProviderRepository.EXPECT().FetchUsersFromProvider().Return(nil, nil)

	mockFileRepository := mocks.NewFileRepository(t)
	mockFileRepository.EXPECT().SaveUsers([]model.User{}).Return(nil)

	s := serviceUser.NewService(
		[]repository.UserProviderRepository{mockUserProviderRepository},
		mockFileRepository,
	)

	err := s.FetchUsersFromProviders()
	assert.Nil(t, err)
}

func TestUserService_FetchFromProviders_UserExist_Success(t *testing.T) {
	users := []model.User{{
		ID:       "test_id",
		Index:    0,
		GUID:     "test_guid",
		IsActive: true,
		Balance:  "test_balance",
		Tags:     []string{"tag_01", "tag_02"},
	}}

	mockUserProviderRepository := mocks.NewUserProviderRepository(t)
	mockUserProviderRepository.EXPECT().FetchUsersFromProvider().Return(users, nil)

	mockFileRepository := mocks.NewFileRepository(t)
	mockFileRepository.EXPECT().SaveUsers(users).Return(nil)

	s := serviceUser.NewService(
		[]repository.UserProviderRepository{mockUserProviderRepository},
		mockFileRepository,
	)

	err := s.FetchUsersFromProviders()
	assert.Nil(t, err)
}

func TestUserService_GetUserByTags_UserExist_Success(t *testing.T) {
	users := []model.User{{
		ID:       "test_id",
		Index:    0,
		GUID:     "test_guid",
		IsActive: true,
		Balance:  "test_balance",
		Tags:     []string{"tag_01", "tag_02"},
	}, {
		ID:       "test_id",
		Index:    1,
		GUID:     "test_guid",
		IsActive: true,
		Balance:  "test_balance",
		Tags:     []string{"tag_02", "tag_03"},
	}, {
		ID:       "test_id",
		Index:    2,
		GUID:     "test_guid",
		IsActive: true,
		Balance:  "test_balance",
		Tags:     []string{"tag_02", "tag_03", "tag_04"},
	}}

	mockUserProviderRepository := mocks.NewUserProviderRepository(t)

	mockFileRepository := mocks.NewFileRepository(t)
	mockFileRepository.EXPECT().GetUsers().Return(users, nil)

	s := serviceUser.NewService(
		[]repository.UserProviderRepository{mockUserProviderRepository},
		mockFileRepository,
	)

	selectedUsers, err := s.GetUserByTags([]string{"tag_01"})
	assert.Nil(t, err)
	assert.Equal(t, []model.User{users[0]}, selectedUsers)

	selectedUsers, err = s.GetUserByTags([]string{"tag_02"})
	assert.Nil(t, err)
	assert.Equal(t, []model.User{users[0], users[1], users[2]}, selectedUsers)

	selectedUsers, err = s.GetUserByTags([]string{"tag_02", "tag_03"})
	assert.Nil(t, err)
	assert.Equal(t, []model.User{users[1], users[2]}, selectedUsers)

	selectedUsers, err = s.GetUserByTags([]string{})
	assert.Nil(t, err)
	assert.Equal(t, []model.User{}, selectedUsers)

}
