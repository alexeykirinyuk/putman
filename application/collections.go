package application

import (
	"github.com/alexeykirinyuk/putman/domain"
	"github.com/google/uuid"
)

type CollectionService struct {
	repo IStorage
}

func NewCollectionService(repo IStorage) *CollectionService {
	return &CollectionService{repo: repo}
}

func (s *CollectionService) Create(name string) (uuid.UUID, error) {
	id := uuid.New()
	col := domain.Collection{
		ID:   id,
		Name: name,
	}

	err := s.repo.Create(col)
	if err != nil {
		return uuid.Nil, err
	}

	return id, nil
}
