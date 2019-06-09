#!/bin/bash
go build -buildmode=plugin ./plugins/plug1 || (echo "Unable to build plugin." && exit 1);
go run cmd/app/main.go
