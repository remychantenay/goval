#!/bin/bash
GOOS=linux go build -o main main.go number.go string.go args.go email.go
