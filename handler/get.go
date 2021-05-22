package handler

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"server_go/storage"
	"strconv"
)

func GetStudents(w http.ResponseWriter, r *http.Request) {

	res := map[string]interface{}{}

	storage.Storage.Range(func(key, value interface{}) bool {
		res[fmt.Sprint(key)] = value

		return true
	})

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	data, err := json.Marshal(res)
	if err != nil {
		log.Fatal(err)
	}
	_, _ = w.Write(data)
}

func GetStudent(w http.ResponseWriter, r *http.Request) {

	id, _ := strconv.Atoi(r.URL.Path[len("/student/"):])

	val, ok := storage.Storage.Load(id)
	if !ok {
		w.WriteHeader(http.StatusNoContent)
		_, _ = w.Write(nil)
	}

	w.WriteHeader(http.StatusOK)
	data, _ := json.Marshal(val)
	_, _ = w.Write(data)
}
