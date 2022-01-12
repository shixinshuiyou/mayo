#!/bin/bash
set -ex

# 容器内编译
echo "build api_gateway"
go build -o /go/src/mayo/bin/api_gateway.bin /go/src/mayo/main.go

echo "build user"
go build -o /go/src/mayo/bin/user.bin /go/src/mayo/app/user/main.go

echo "build id"
go build -o /go/src/mayo/bin/id.bin /go/src/mayo/app/id/main.go