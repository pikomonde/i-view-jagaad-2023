package userprovider

import (
	"encoding/json"
	"errors"
	"i-view-jagaad-2023/model"
	"i-view-jagaad-2023/repository"
	"net/http"

	log "github.com/sirupsen/logrus"
)

type UserProvider struct {
	Cli *http.Client
	URL string
}

func NewRepository(
	cli *http.Client,
	url string,
) repository.UserProviderRepository {
	return &UserProvider{
		Cli: cli,
		URL: url,
	}
}

func (r *UserProvider) FetchUsersFromProvider() ([]model.User, error) {

	resp, err := r.Cli.Get(r.URL)
	if err != nil {
		log.Errorf("Error making http request, err : %s", err.Error())
		return nil, errors.New(model.ErrorProvider)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Errorf("Error from provider with code: %d", resp.StatusCode)
		return nil, errors.New(model.ErrorProvider)
	}

	var users []model.User
	err = json.NewDecoder(resp.Body).Decode(&users)
	if err != nil {
		log.Errorf("Error unmarshall response body, err : %s", err.Error())
		return nil, errors.New(model.ErrorProvider)
	}

	return users, nil
}
