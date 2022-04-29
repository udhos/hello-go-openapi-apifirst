//go:generate oapi-codegen -package pets -o pets.gen.go -generate types,client,chi-server,spec ../openapi/petstore-expanded.yaml

package pets
