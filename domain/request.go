package domain

import "github.com/google/uuid"

// Request is just HTTP request
type Request struct {
	ID      uuid.UUID `json:"id"`
	Headers []Header  `json:"headers"`
	Method  string    `json:"method"`
	Name    string    `json:"name"`
	URL     string    `json:"url"`
	Body    string    `json:"body"`
}

// Header is request HTTP header
type Header struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}
