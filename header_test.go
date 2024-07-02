package belajar_golang_web

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func RequestHeader(w http.ResponseWriter, r *http.Request) {
	ContentType := r.Header.Get("content-type")
	w.Header().Add("X-Powered-By", "AkmalMP")
	fmt.Fprint(w, ContentType)
}

func TestRequestHeader(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8000/", nil)
	request.Header.Add("Content-Type", "application/json")
	recorder := httptest.NewRecorder()

	RequestHeader(recorder, request)

	result := recorder.Result()
	body, _ := io.ReadAll(result.Body)
	header := result.Header.Get("X-Powered-By")

	fmt.Println(string(body))
	fmt.Println(header)
}
