package application

import (
	"github.com/alexeykirinyuk/putman/domain"
	"github.com/google/uuid"
)

type RequestDto struct {
	ID   uuid.UUID
	Name string
}

type FolderDto struct {
	Name     string
	Folders  []FolderDto
	Requests []RequestDto
}

type CollectionDto struct {
	ID       uuid.UUID
	Name     string
	Folders  []FolderDto
	Requests []RequestDto
}

func (s *CollectionService) GetTree() ([]CollectionDto, error) {
	collections, err := s.repo.GetAll()
	if err != nil {
		return []CollectionDto{}, err
	}

	colDtos := []CollectionDto{}

	for _, col := range collections {
		colDtos = append(colDtos, mapCollection(col))
	}

	return colDtos, nil
}

func mapCollection(col domain.Collection) CollectionDto {
	colDto := CollectionDto{
		ID:   col.ID,
		Name: col.Name,
	}

	for _, req := range col.Requests {
		colDto.Requests = append(colDto.Requests, mapRequest(req))
	}

	for _, fol := range col.Folders {
		colDto.Folders = append(colDto.Folders, mapFolder(fol))
	}

	return colDto
}

func mapFolder(fol domain.Folder) FolderDto {
	folDto := FolderDto{
		Name: fol.Name,
	}
	for _, req := range fol.Requests {
		folDto.Requests = append(folDto.Requests, RequestDto{
			ID:   req.ID,
			Name: req.Name,
		})
	}
	for _, childFol := range fol.Folders {
		folDto.Folders = append(folDto.Folders, mapFolder(childFol))
	}

	return folDto
}

func mapRequest(req domain.Request) RequestDto {
	return RequestDto{
		ID:   req.ID,
		Name: req.Name,
	}
}
