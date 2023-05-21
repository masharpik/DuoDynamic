package delivery

import (
	"github.com/gorilla/mux"

	"github.com/go-park-mail-ru/2023_1_MRGA.git/utils/env_getter"

	"github.com/go-park-mail-ru/2023_1_MRGA.git/internal/pkg/photo"
)

type Handler struct {
	useCase    photo.UseCase
	serverHost string
}

func NewHandler(useCase photo.UseCase) *Handler {
	return &Handler{
		useCase:    useCase,
		serverHost: env_getter.GetHostFromEnv("FILE_STORAGE_SERVER_HOST"),
	}
}

func RegisterHTTPEndpoints(router *mux.Router, uc photo.UseCase) {
	h := NewHandler(uc)
	router.HandleFunc("/meetme/photos/upload", h.AddPhoto).Methods("POST")
	router.HandleFunc("/meetme/files/upload", h.AddFiles).Methods("POST")
	router.HandleFunc("/meetme/photo/{photo}", h.GetPhoto).Methods("GET")
	router.HandleFunc("/meetme/file/{pathToFile:.*}", h.GetFile).Methods("GET")
	router.HandleFunc("/meetme/audio/transcribe/{pathToFile:.*}", h.GetTranscription).Methods("GET")
	router.HandleFunc("/meetme/photo/{photo}", h.DeletePhoto).Methods("DELETE")
	router.HandleFunc("/meetme/photo/{photo}", h.ChangePhoto).Methods("PUT", "OPTIONS")
}
