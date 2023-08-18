package file

import (
	"encoding/csv"
	"encoding/json"
	"errors"
	"fmt"
	"i-view-jagaad-2023/model"
	"i-view-jagaad-2023/repository"
	"os"
	"strconv"
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
			log.Errorf("Error marshal friends for user GUID %s, with %s", user.GUID, err.Error())
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

	// opening up file
	file, err := os.Open(r.Filename)
	if err != nil {
		log.Errorf("Error open, err: %s", err.Error())
		return nil, errors.New(model.ErrorInternalService)
	}
	defer file.Close()

	// setting up writer
	reader := csv.NewReader(file)

	// read all from csv
	rows, err := reader.ReadAll()
	if err != nil {
		log.Errorf("Error read, err: %s", err.Error())
		return nil, errors.New(model.ErrorInternalService)
	}

	// read data, discard header
	users := make([]model.User, 0)
	for _, col := range rows[1:] {
		if len(col) != 7 {
			log.Errorf("Error number column not match, malformed data")
			return nil, errors.New(model.ErrorInternalService)
		}

		index, err := strconv.Atoi(col[1])
		if err != nil {
			log.Errorf("Error convert index to integer, err: %s", err.Error())
			return nil, errors.New(model.ErrorInternalService)
		}

		isActive := false
		if col[3] == "true" {
			isActive = true
		}

		tags := strings.Split(col[5], ", ")

		var friends []model.Friend
		err = json.Unmarshal([]byte(col[6]), &friends)
		if err != nil {
			log.Errorf("Error unmarshal friends for user GUID %s, with %s", col[2], err.Error())
			return nil, errors.New(model.ErrorInternalService)
		}

		users = append(users, model.User{
			ID:       col[0],
			Index:    uint32(index),
			GUID:     col[2],
			IsActive: isActive,
			Balance:  col[4],
			Tags:     tags,
			Friends:  friends,
		})
	}

	return users, nil
}
