#!/bin/bash

CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -o users ./users/
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -o tasks ./tasks/