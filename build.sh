#!/bin/bash

hash oapi-codegen 2>/dev/null || go install github.com/deepmap/oapi-codegen/cmd/oapi-codegen@latest

go generate ./pets

go mod tidy

go test ./...

go install ./...

