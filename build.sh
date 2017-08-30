#!/usr/bin/env bash
env GOOS=windows GOARCH=amd64 go build -o rtd.exe
env GOOS=linux GOARCH=amd64 go build -o rtd.sh
env GOOS=darwin GOARCH=amd64 go build -o rtd.mac