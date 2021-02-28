package persistence

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/alexeykirinyuk/putman/domain"
	"github.com/google/uuid"
)

// JsonStorage @impl application.ICollectionRepository
type JsonStorage struct {
	filePath string
}

func NewJsonStorage(filePath string) *JsonStorage {
	return &JsonStorage{filePath: filePath}
}

func (r *JsonStorage) GetAll() ([]domain.Collection, error) {
	if err := r.createFileIfNotExists(); err != nil {
		return []domain.Collection{}, err
	}

	byteArray, err := ioutil.ReadFile(r.filePath)
	if err != nil {
		return []domain.Collection{}, fmt.Errorf("error when read json file: %s", err)
	}

	var collections []domain.Collection

	err = json.Unmarshal(byteArray, &collections)
	if err != nil {
		return []domain.Collection{}, fmt.Errorf("error when unmarshal json: %s", err)
	}

	return collections, nil
}

func (r *JsonStorage) Get(id uuid.UUID) (domain.Collection, error) {
	if err := r.createFileIfNotExists(); err != nil {
		return domain.Collection{}, err
	}

	collections, err := r.GetAll()
	if err != nil {
		return domain.Collection{}, err
	}

	for _, collection := range collections {
		if collection.ID == id {
			return collection, nil
		}
	}

	return domain.Collection{}, errors.New("Collection not found")
}

func (r *JsonStorage) Create(col domain.Collection) error {
	if err := r.createFileIfNotExists(); err != nil {
		return err
	}

	collections, err := r.GetAll()
	if err != nil {
		return err
	}

	collections = append(collections, col)
	return r.updateAll(collections)
}

func (r *JsonStorage) Update(col domain.Collection) error {
	if err := r.createFileIfNotExists(); err != nil {
		return err
	}

	collections, err := r.GetAll()
	if err != nil {
		return err
	}

	ok, index := findIndex(collections, col.ID)
	if !ok {
		return errors.New("collection not found")
	}
	collections[index] = col

	return r.updateAll(collections)
}

func findIndex(collections []domain.Collection, id uuid.UUID) (bool, int) {
	for i, collection := range collections {
		if collection.ID == id {
			return true, i

		}
	}
	return false, 0
}

func (r *JsonStorage) createFileIfNotExists() error {
	_, err := os.Stat(r.filePath)
	if err == nil {
		return nil
	} else if !os.IsExist(err) {
		return r.updateAll([]domain.Collection{})
	} else {
		return fmt.Errorf("error when check file exists: %s", err)
	}
}

func (r *JsonStorage) updateAll(collections []domain.Collection) error {
	byteArray, err := json.Marshal(collections)
	if err != nil {
		return fmt.Errorf("error when marshall json: %s", err)
	}

	err = ioutil.WriteFile(r.filePath, byteArray, 0644)
	if err != nil {
		return fmt.Errorf("error when write file: %s", err)
	}

	return nil
}
