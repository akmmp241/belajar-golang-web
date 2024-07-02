package belajar_golang_web

import (
	"net/http"
	"testing"
)

func DownloadFile(writer http.ResponseWriter, request *http.Request) {
	fileName := request.URL.Query().Get("file")
	if fileName == "" {
		http.Error(writer, "file name is empty", http.StatusBadRequest)
		return
	}

	writer.Header().Add("Content-Disposition", "inline; filename=\""+fileName+"\"")
	http.ServeFile(writer, request, "./resources/"+fileName)
}

func TestDownloadFile(t *testing.T) {
	server := http.Server{
		Addr:    "localhost:8000",
		Handler: http.HandlerFunc(DownloadFile),
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
