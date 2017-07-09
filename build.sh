#!/usr/bin/env bash
env GOOS=windows GOARCH=amd64 go build -o bin/rtd.exe
env GOOS=linux GOARCH=amd64 go build -o bin/rtd.sh
env GOOS=darwin GOARCH=amd64 go build -o bin/rtd.mac