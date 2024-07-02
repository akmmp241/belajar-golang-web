package belajar_golang_web

import (
	"fmt"
	"net/http"
	"testing"
	"time"
)

type LogMiddleware struct {
	Handler http.Handler
}

type ErrorHandler struct {
	Handler http.Handler
}

func (middleware *LogMiddleware) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Before Execute Handler")
	time.Sleep(5 * time.Second)
	middleware.Handler.ServeHTTP(w, r)
	time.Sleep(5 * time.Second)
	fmt.Println("After Execute Handler")
}

func TestMiddleware(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello Middleware")
	})

	logMiddleware := LogMiddleware{
		Handler: mux,
	}

	server := http.Server{
		Addr:    "localhost:8000",
		Handler: &logMiddleware,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}

func (handler *ErrorHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	defer func() {
		err := recover()
		fmt.Println("RECOVER", err)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprint(w, "Error: ", err)
		}
	}()
	handler.Handler.ServeHTTP(w, r)
}

func TestErrorHandler(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		panic("Ups, Error")
	})

	errorHandler := &ErrorHandler{
		Handler: &LogMiddleware{
			Handler: mux,
		},
	}

	server := http.Server{
		Addr:    "localhost:8000",
		Handler: errorHandler,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
