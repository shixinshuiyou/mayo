#!/bin/bash
set -ex

# 容器内编译
echo "build api_gateaway"
go build -o /go/src/mayo/bin/api_gateaway.bin /go/src/mayo-dev/main.go