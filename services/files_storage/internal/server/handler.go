package server

import (
	"errors"
	"fmt"
	"net/http"
	"os"
	"path"
	"strconv"
	"time"

	"github.com/go-park-mail-ru/2023_1_MRGA.git/services/files_storage/internal/repository"
	"github.com/go-park-mail-ru/2023_1_MRGA.git/services/files_storage/internal/service"
	"github.com/go-park-mail-ru/2023_1_MRGA.git/utils/writer"
	"github.com/gorilla/mux"
)

func getRouter() *mux.Router {
	router := mux.NewRouter()

	// router.Use(SetCorsMiddleware)

	router.HandleFunc("/api/files/upload", uploadFile).Methods("POST")
	router.HandleFunc("/api/files/{id}", getFile).Methods("GET")
	return router
}

func uploadFile(w http.ResponseWriter, r *http.Request) {
	file, handler, err := r.FormFile("file")
	if err != nil {
		writer.ErrorRespond(w, r, err, http.StatusBadRequest)
		return
	}
	defer file.Close()

	userID := r.FormValue("userID")
	if userID == "" {
		writer.ErrorRespond(w, r, errors.New("Не передан id пользователя"), http.StatusBadRequest)
		return
	}

	userIDUInt64, err := strconv.ParseUint(userID, 10, 32)
	if err != nil {
		writer.ErrorRespond(w, r, err, http.StatusBadRequest)
		return
	}

	filePath, err := service.UploadFile(file, handler, uint(userIDUInt64))
	if err != nil {
		writer.ErrorRespond(w, r, err, http.StatusInternalServerError)
		return
	}

	fileID, err := repository.UploadFile(filePath, uint(userIDUInt64))
	if err != nil {
		writer.ErrorRespond(w, r, err, http.StatusInternalServerError)
		return
	}

	writer.Respond(w, r, map[string]interface{}{
		"photoID": fileID,
	})
}

func getFile(w http.ResponseWriter, r *http.Request) {
	// Получаем ID файла из параметра запроса
	vars := mux.Vars(r)
	id := vars["id"]

	idUInt64, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Получаем содержимое файла из БД
	var file repository.File
	if file, err = repository.GetFile(uint(idUInt64)); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	gotFile, err := os.Open(file.Path)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer gotFile.Close()

	filename := path.Base(gotFile.Name())
	// Устанавливаем заголовки для ответа
	w.Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=%s", filename))

	// Автоматически ставит заголовок image/jpg или image/png
	http.ServeContent(w, r, filename, time.Now(), gotFile)
}