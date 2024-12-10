package project

import (
	"errors"
	"file-sync-tool-api/types"
	"file-sync-tool-api/utils"
	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
	"net/http"
)

type Handler struct {
	projectStore types.ProjectStore
}

func NewHandler(projectStore *Store) *Handler {
	return &Handler{projectStore: projectStore}
}

func (handler *Handler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/project", handler.handleGetProject).Methods(http.MethodGet)
	router.HandleFunc("/project/add", handler.handleCreateProjectFile).Methods(http.MethodPost)
}

func (handler *Handler) handleGetProject(w http.ResponseWriter, _ *http.Request) {
	project, err := handler.projectStore.GetProject()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = utils.WriteJSON(w, http.StatusOK, project)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (handler *Handler) handleCreateProjectFile(w http.ResponseWriter, r *http.Request) {
	var payload types.AddFileToProjectPayload
	if err := utils.ParseJSON(r, &payload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	if err := utils.Validate.Struct(payload); err != nil {
		var errValidation validator.ValidationErrors
		errors.As(err, &errValidation)
		utils.WriteError(w, http.StatusBadRequest, errValidation)
	}

	err := handler.projectStore.AddFileToProject(types.AddFileToProjectPayload{
		FileID: payload.FileID,
		Path:   payload.Path,
	})
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	err = utils.WriteJSON(w, http.StatusCreated, nil)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}
}
