package types

import "time"

type HistoryStore interface {
	GetHistory() ([]*History, error)
	CreateHistory(payload *CreateHistoryPayload) (*History, error)
}

type History struct {
	ID             string    `json:"id"`
	Author         string    `json:"author"`
	Description    string    `json:"description"`
	UpdatedAt      time.Time `json:"updated_at"`
	UpdatedFileIDs []string  `json:"updated_files"`
}

type CreateHistoryPayload struct {
	Author         string   `json:"author"`
	Description    string   `json:"description"`
	UpdatedFileIDs []string `json:"updated_files"`
}
