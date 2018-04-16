#!/bin/bash
GOOS=linux go build -gcflags="-m -m" -o main main.go number.go string.go args.go email.go uuid.go country.go currency.go enum.go
