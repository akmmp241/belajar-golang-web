package belajar_golang_web

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func SayHello(w http.ResponseWriter, req *http.Request) {
	name := req.URL.Query().Get("name")
	if name == "" {
		fmt.Fprint(w, "Hello")
	} else {
		fmt.Fprint(w, "Hello "+name)
	}
}

func TestQueryParam(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8000/hello?name=akmal", nil)
	recorder := httptest.NewRecorder()

	SayHello(recorder, request)

	result := recorder.Result()
	body, err := io.ReadAll(result.Body)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(body))
}

func MultipleQueryParam(w http.ResponseWriter, req *http.Request) {
	firstName := req.URL.Query().Get("first_name")
	lastName := req.URL.Query().Get("last_name")

	fmt.Fprint(w, "Hello "+firstName+" "+lastName)
}

func TestMultipleQueryParam(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8000/hello?first_name=akmal&last_name=mp", nil)
	recorder := httptest.NewRecorder()

	MultipleQueryParam(recorder, request)

	result := recorder.Result()
	body, err := io.ReadAll(result.Body)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(body))
}

func MultipleParamValues(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	names := query["name"]
	fmt.Fprint(w, strings.Join(names, ","))
}

func TestMultipleParamValues(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8000/hello?name=akmal&name=mp", nil)
	recorder := httptest.NewRecorder()

	MultipleParamValues(recorder, request)

	result := recorder.Result()
	body, err := io.ReadAll(result.Body)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(body))
}
