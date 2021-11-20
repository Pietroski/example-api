#!/usr/bin/env bash

go mod download

# shellcheck disable=SC2034
# shellcheck disable=SC2034
GO111MODULE=on \
CGO_ENABLED=0 \
GOOS=linux \
GOARCH=amd64 &&
go build -o build/lambda/auth/signup/main ./cmd/lambda/auth/signup/main.go
