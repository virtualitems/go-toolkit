package main

import (
	"io"
	"net/http"
	"os"
)

type HealthHandler struct{}

func (HealthHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNoContent)
}

type Route struct {
	pattern string
	handler http.Handler
}

type Routes struct {
	entries []Route
}

type EchoHandler struct{}

func (EchoHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)

	_, copyError := io.Copy(w, r.Body)
	if copyError != nil {
		os.Exit(1)
	}
}

type ColorsIndexHandler struct{}

func (ColorsIndexHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	responseBytes := []byte("index")

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	_, writeError := w.Write(responseBytes)
	if writeError != nil {
		os.Exit(1)
	}
}

type ColorsCreateHandler struct{}

func (ColorsCreateHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	responseBytes := []byte("create")

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	_, writeError := w.Write(responseBytes)
	if writeError != nil {
		os.Exit(1)
	}
}

type ColorShowHandler struct{}

func (ColorShowHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	responseBytes := []byte("show")

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	_, writeError := w.Write(responseBytes)
	if writeError != nil {
		os.Exit(1)
	}
}

type ColorsStoreHandler struct{}

func (ColorsStoreHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNoContent)
}

type ColorUpdateHandler struct{}

func (ColorUpdateHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNoContent)
}

type ColorDestroyHandler struct{}

func (ColorDestroyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNoContent)
}

func main() {
	mux := http.NewServeMux()

	routes := Routes{
		entries: []Route{
			{pattern: "GET /", handler: HealthHandler{}},
			{pattern: "POST /", handler: EchoHandler{}},

			{pattern: "GET /colors", handler: ColorsIndexHandler{}},
			{pattern: "GET /colors/create", handler: ColorsCreateHandler{}},
			{pattern: "GET /colors/{slug}/edit", handler: ColorShowHandler{}},
			{pattern: "GET /colors/{slug}", handler: ColorShowHandler{}},

			{pattern: "POST /colors", handler: ColorsStoreHandler{}},
			{pattern: "PATCH /colors/{slug}", handler: ColorUpdateHandler{}},
			{pattern: "PUT /colors/{slug}", handler: ColorUpdateHandler{}},
			{pattern: "DELETE /colors/{slug}", handler: ColorDestroyHandler{}},
		},
	}

	for _, entry := range routes.entries {
		mux.Handle(entry.pattern, entry.handler)
	}

	serverError := http.ListenAndServe(":8000", mux)
	if serverError != nil {
		os.Exit(1)
	}
}
