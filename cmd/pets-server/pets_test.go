package main

import (
	"bytes"
	"net/http"
	"testing"

	pets "github.com/udhos/hello-go-openapi-apifirst/pets"
)

type responseWriter struct {
	buf *bytes.Buffer
}

func (w responseWriter) Header() http.Header { return http.Header{} }
func (w responseWriter) Write(p []byte) (int, error) {
	return w.buf.Write(p)
}
func (w responseWriter) WriteHeader(statusCode int) {}

func FindPetsEmptyStoreTest(t *testing.T) {
	store := NewPetStore()
	w := responseWriter{buf: &bytes.Buffer{}}
	r := http.Request{}
	params := pets.FindPetsParams{}
	store.FindPets(w, &r, params)
	result := w.buf.String()
	if result != "[]" {
		t.Errorf("not empty response: %s", result)
	}
}
