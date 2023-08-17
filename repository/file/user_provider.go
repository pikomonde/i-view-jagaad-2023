package File

import (
	"encoding/csv"
	"encoding/json"
	"errors"
	"fmt"
	"i-view-jagaad-2023/model"
	"i-view-jagaad-2023/repository"
	"os"
	"strings"

	log "github.com/sirupsen/logrus"
)

type File struct {
	Filename string
}

func NewRepository(
	filename string,
) repository.FileRepository {
	return &File{
		Filename: filename,
	}
}

func (r *File) SaveUsers(users []model.User) error {

	// setting up file
	file, err := os.Create(r.Filename)
	if err != nil {
		log.Errorf("Error create, err: %s", err.Error())
		return errors.New(model.ErrorInternalService)
	}
	defer file.Close()

	// setting up writer
	writer := csv.NewWriter(file)
	defer writer.Flush()

	// write header
	err = writer.Write([]string{
		"ID",
		"Index",
		"GUID",
		"IsActive",
		"Balance",
		"Tags",
		"Friends",
	})
	if err != nil {
		log.Errorf("Error write, err: %s", err.Error())
		return errors.New(model.ErrorInternalService)
	}

	// write data
	for _, user := range users {
		isActive := "false"
		if user.IsActive {
			isActive = "true"
		}

		tags := strings.Join(user.Tags, ", ")

		friends, err := json.Marshal(user.Friends)
		if err != nil {
			log.Errorf("Error marshal friends for user GUID %s, with %s", user.GUID, err)
			return errors.New(model.ErrorInternalService)
		}

		err = writer.Write([]string{
			user.ID,
			fmt.Sprintf("%d", user.Index),
			user.GUID,
			isActive,
			user.Balance,
			tags,
			string(friends),
		})
		if err != nil {
			log.Errorf("Error Write: %s", err.Error())
			return errors.New(model.ErrorInternalService)
		}
	}

	return nil
}

func (r *File) GetUsers() ([]model.User, error) {

	return nil, nil
}
