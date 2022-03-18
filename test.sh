#!/bin/bash

CVPKG=$(go list ./... | grep -v mocks | grep -v cmd | tr '\n' ',')

go test -short -count 1 -race -coverpkg $CVPKG -coverprofile coverage.txt ./...
go tool cover -func coverage.txt
go tool cover -html coverage.txt