package _type

type ProjectStore interface {
	GetProject() (*Project, error)
	AddFileToProject(payload *AddFileToProjectPayload) (*File, error)
	DeleteFileFromProject(payload *DeleteFileFromProjectPayload) error
	EditFileInProject(payload *EditFileInProjectPayload) (*File, error)
}

type Project struct {
	FileID  string `json:"file_id"`
	Path    string `json:"path"`
	Version string `json:"version"`
}

type AddFileToProjectPayload struct {
	Path string `json:"path"`
}

type DeleteFileFromProjectPayload struct {
	FileID string `json:"file_id"`
}

type EditFileInProjectPayload struct {
	FileID  string `json:"file_id"`
	Name    string `json:"name"`
	Content string `json:"content"`
}
