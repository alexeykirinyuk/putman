package application

import (
	"bytes"
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/alexeykirinyuk/putman/domain"
	"github.com/google/uuid"
)

type RequstService struct {
	repo IStorage
}

func NewRequestService(repo IStorage) *RequstService {
	return &RequstService{repo: repo}
}

func (s *RequstService) Create(collectionName string, requestName string, method string, url string) (uuid.UUID, error) {
	collections, err := s.repo.GetAll()
	if err != nil {
		return uuid.Nil, err
	}

	var col domain.Collection
	for _, c := range collections {
		if c.Name == collectionName {
			col = c
		}
	}

	id := uuid.New()
	col.Requests = append(col.Requests, domain.Request{
		ID:      id,
		Headers: []domain.Header{},
		Method:  method,
		Name:    collectionName,
		URL:     url,
		Body:    "",
	})

	s.repo.Update(col)

	return id, nil
}

func (s *RequstService) Execute(id uuid.UUID) (resp string, err error) {
	cols, err := s.repo.GetAll()
	if err != nil {
		return
	}

	ok, req := find(cols, id)
	if !ok {
		err = errors.New("request not found")
		return
	}

	resp, err = do(req, resp)
	if err != nil {
		err = fmt.Errorf("error when make request: %s", err)
	}

	return
}

func do(req domain.Request, resp string) (string, error) {
	httpReq, err := http.NewRequest(strings.ToUpper(req.Method), req.URL, nil)
	if err != nil {
		err = fmt.Errorf("error when create request: %s", err)
		return "", err
	}

	httpResp, err := http.DefaultClient.Do(httpReq)
	if err != nil {
		return "", err
	}

	defer httpResp.Body.Close()

	buf := new(bytes.Buffer)
	buf.ReadFrom(httpResp.Body)
	resp = buf.String()

	return resp, nil
}

func find(cols []domain.Collection, reqID uuid.UUID) (bool, domain.Request) {
	for _, c := range cols {
		for _, r := range c.Requests {
			if r.ID == reqID {
				return true, r
			}
		}
	}

	return false, domain.Request{}
}
