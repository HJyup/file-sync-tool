package project

import (
	"database/sql"
	"file-sync-tool-api/types"
)

type Project struct {
	FileID  string `json:"file_id"`
	Path    string `json:"path"`
	Version string `json:"version"`
}

type Store struct {
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{db}
}

func (store *Store) DeleteFileFromProject(payload *types.DeleteFileFromProjectPayload) error {
	_, err := store.db.Exec("DELETE FROM Project WHERE fileId = ?", payload.FileID)
	if err != nil {
		return err
	}

	return nil
}

func (store *Store) EditFileInProject(payload *types.Project) error {
	_, err := store.db.Exec("UPDATE Project SET path = ?, version = ? WHERE fileId = ?",
		payload.Path, payload.Version, payload.FileID)
	if err != nil {
		return err
	}

	return nil
}

func (store *Store) AddFileToProject(payload types.AddFileToProjectPayload) error {
	_, err := store.db.Exec("INSERT INTO Project (fileId, path) VALUES (?, ?)",
		payload.FileID, payload.Path)
	if err != nil {
		return err
	}

	return nil
}

func (store *Store) GetProject() ([]types.Project, error) {
	rows, err := store.db.Query("SELECT * FROM Project")
	if err != nil {
		return nil, err
	}

	project := make([]types.Project, 0)
	for rows.Next() {
		projectRow, err := scanRowIntoProduct(rows)
		if err != nil {
			return nil, err
		}

		project = append(project, *projectRow)
	}

	return project, nil
}

func scanRowIntoProduct(rows *sql.Rows) (*types.Project, error) {
	project := new(types.Project)
	err := rows.Scan(
		&project.FileID,
		&project.Path,
		&project.Version,
	)
	if err != nil {
		return nil, err
	}

	return project, nil
}
