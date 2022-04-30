package main

import (
	"bytes"
	"net/http"
	"testing"

	pets "github.com/udhos/hello-go-openapi-apifirst/pets"
)

type responseWriter struct {
	buf    *bytes.Buffer
	status *int
}

func (w responseWriter) Header() http.Header { return http.Header{} }
func (w responseWriter) Write(p []byte) (int, error) {
	return w.buf.Write(p)
}
func (w responseWriter) WriteHeader(statusCode int) { w.status = &statusCode }

func FindPetsEmptyStoreTest(t *testing.T) {
	store := NewPetStore()
	w := responseWriter{buf: &bytes.Buffer{}}
	r := http.Request{}
	params := pets.FindPetsParams{}
	store.FindPets(w, &r, params)

	if *w.status != 200 {
		t.Errorf("unexpected status: %d", *w.status)
	}

	result := w.buf.String()
	if result != "[]" {
		t.Errorf("unexpected non-empty response: %s", result)
	}
}
