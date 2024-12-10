package types

type ProjectStore interface {
	GetProject() ([]Project, error)
	AddFileToProject(payload AddFileToProjectPayload) error
	DeleteFileFromProject(payload *DeleteFileFromProjectPayload) error
	EditFileInProject(payload *Project) error
}

type Project struct {
	FileID  string `json:"file_id"`
	Path    string `json:"path"`
	Version string `json:"version"`
}

type AddFileToProjectPayload struct {
	FileID string `json:"file_id"`
	Path   string `json:"path"`
}

type DeleteFileFromProjectPayload struct {
	FileID string `json:"file_id"`
}
