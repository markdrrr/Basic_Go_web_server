package handler

import (
	"net/http"
	"server_go/storage"
	"strconv"
)

func Delete(w http.ResponseWriter, r *http.Request) {

	id, _ := strconv.Atoi(r.URL.Path[len("/student/"):])

	storage.Storage.Delete(id)

	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte("Deleted"))
}