package domain

import "github.com/google/uuid"

// Collection is structured group of http requests
type Collection struct {
	ID       uuid.UUID `json:"id"`
	Name     string    `json:"name"`
	Folders  []Folder  `json:"folders"`
	Requests []Request `json:"requests"`
}
