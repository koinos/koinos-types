#!/bin/bash

export GOPATH="`go env GOPATH`:`pwd`/../../../../generated/golang"
go run test.go
