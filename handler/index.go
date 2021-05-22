package handler

import (
	"fmt"
	"log"
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request)  {
	fmt.Println(r.Method, r.URL)
	_, err := w.Write([]byte("Welcome"))
	if err != nil {
		log.Fatal(err)
	}
}
