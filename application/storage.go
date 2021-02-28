package application

import (
	"github.com/alexeykirinyuk/putman/domain"
	"github.com/google/uuid"
)

type IStorage interface {
	GetAll() ([]domain.Collection, error)
	Get(id uuid.UUID) (domain.Collection, error)
	Create(col domain.Collection) error
	Update(col domain.Collection) error
}
