package types

type FileStore interface {
	CreateFile(payload *CreateFilePayload) (*File, error)
}

type File struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	Content string `json:"content"`
}

type CreateFilePayload struct {
	Name    string `json:"name"`
	Content string `json:"content"`
}
