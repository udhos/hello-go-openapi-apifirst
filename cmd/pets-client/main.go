package main

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/deepmap/oapi-codegen/pkg/securityprovider"

	"github.com/udhos/hello-go-openapi-apifirst/env"
	"github.com/udhos/hello-go-openapi-apifirst/pets"
)

func main() {

	// Example BasicAuth
	// See: https://swagger.io/docs/specification/authentication/basic-authentication/
	basicAuthProvider, basicAuthProviderErr := securityprovider.NewSecurityProviderBasicAuth("MY_USER", "MY_PASS")
	if basicAuthProviderErr != nil {
		panic(basicAuthProviderErr)
	}

	// Exhaustive list of some defaults you can use to initialize a Client.
	// If you need to override the underlying httpClient, you can use the option
	//
	// WithHTTPClient(httpClient *http.Client)
	//
	client, errClient := pets.NewClient(env.String("SERVER_URL", "http://localhost:8080"),
		pets.WithRequestEditorFn(basicAuthProvider.Intercept))
	if errClient != nil {
		panic(errClient)
	}

	cmd := "list"
	if len(os.Args) > 1 {
		cmd = os.Args[1]
	}

	fmt.Printf("cmd: %s\n", cmd)

	me := filepath.Base(os.Args[0])

	switch cmd {
	case "list":
		params := &pets.FindPetsParams{}
		resp, errFind := client.FindPets(context.TODO(), params)
		showResponse(resp, errFind)
	case "get":
		if len(os.Args) < 3 {
			fmt.Fprintf(os.Stderr, "get: missing id")
			usage(me)
			os.Exit(1)
		}
		idStr := os.Args[2]
		id, errConv := strconv.ParseInt(idStr, 10, 64)
		if errConv != nil {
			fmt.Fprintf(os.Stderr, "get: bad id: %v", errConv)
			usage(me)
			os.Exit(1)
		}
		resp, errFind := client.FindPetByID(context.TODO(), id)
		showResponse(resp, errFind)
	case "add":
		if len(os.Args) < 3 {
			fmt.Fprintf(os.Stderr, "add: missing name")
			usage(me)
			os.Exit(1)
		}
		name := os.Args[2]
		body := fmt.Sprintf(`{"name":"%s"}`, name)
		resp, errAdd := client.AddPetWithBody(context.TODO(), "application/json", strings.NewReader(body))
		showResponse(resp, errAdd)
	case "del":
		if len(os.Args) < 3 {
			fmt.Fprintf(os.Stderr, "del: missing id")
			usage(me)
			os.Exit(1)
		}
		idStr := os.Args[2]
		id, errConv := strconv.ParseInt(idStr, 10, 64)
		if errConv != nil {
			fmt.Fprintf(os.Stderr, "del: bad id: %v", errConv)
			usage(me)
			os.Exit(1)
		}
		resp, errDel := client.DeletePet(context.TODO(), id)
		showResponse(resp, errDel)
	default:
		usage(me)
		os.Exit(1)
	}
}

func usage(me string) {
	fmt.Printf("usage: %s list|get <id>|add <name>|del <id>", me)
}

func showResponse(resp *http.Response, err error) {
	if err != nil {
		fmt.Printf("error: %v\n", err)
		return
	}
	fmt.Printf("status: %d\n", resp.StatusCode)
	fmt.Println("body:")
	io.Copy(os.Stdout, resp.Body)
}
