#!/usr/bin/env bash

go mod tidy
go mod graph
GOOS=darwin go build -ldflags '-w -s' -o minion-x64-darwin -i -v cmd/minion.go
GOOS=linux  go build -ldflags '-w -s' -o minion-x64-linux  -i -v cmd/minion.go