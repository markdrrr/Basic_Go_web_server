package handler

import (
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"io"
	"net/http"
	"server_go/model"
	"server_go/storage"
)

func Add(w http.ResponseWriter, r *http.Request)  {
	switch r.Method {
	case "GET":
		// redirect to index page
		http.Redirect(w, r, fmt.Sprintf("/"), http.StatusFound)
	case "POST":
		b, err := io.ReadAll(r.Body)
		if err != nil {
			fmt.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
		}

		var student model.Student
		if err := json.Unmarshal(b, &student); err != nil {
			fmt.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
		}
		fmt.Println("add new student", student)

		u, _ := uuid.NewUUID()
		student.ID = int(u.ID())
		// record student to storage
		storage.Storage.Store(student.ID, student)

		// return object student
		w.WriteHeader(http.StatusOK)
		data, _ := json.Marshal(student)
		_, _ = w.Write(data)

	}
}
