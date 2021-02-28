package domain

// Folder is tree group of requests
type Folder struct {
	Name     string    `json:"name"`
	Requests []Request `json:"requests"`
	Folders  []Folder  `json:"folders"`
}
