package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"

	middleware "github.com/deepmap/oapi-codegen/pkg/chi-middleware"

	api "github.com/udhos/hello-go-openapi-apifirst/pets"
)

func main() {
	var addr string

	flag.StringVar(&addr, "addr", ":8080", "server http listen address")
	flag.Parse()

	swagger, err := api.GetSwagger()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error loading swagger spec\n: %s", err)
		os.Exit(1)
	}

	// Clear out the servers array in the swagger spec, that skips validating
	// that server names match. We don't know how this thing will be run.
	swagger.Servers = nil

	// Create an instance of our handler which satisfies the generated interface
	petStore := NewPetStore()

	// This is how you set up a basic chi router
	r := chi.NewRouter()

	// Use our validation middleware to check all requests against the
	// OpenAPI schema.
	r.Use(middleware.OapiRequestValidator(swagger))

	// We now register our petStore above as the handler for the interface
	api.HandlerFromMux(petStore, r)

	s := &http.Server{
		Handler: r,
		Addr:    addr,
	}

	log.Printf("listning on TCP %s", addr)
	log.Fatal(s.ListenAndServe())
}
