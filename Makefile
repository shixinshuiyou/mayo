GOPATH:=$(shell go env GOPATH)

.PHONY: run gateway user id proto
build-images:
	docker build -f docker/base.dockerfile -t shixinshuiyou/mayo:${v} .
run:
	sh run.sh ${v}
gateway:
	go run -o main.go api --namespace=czh.micro.api --enable_cors=false --address 0.0.0.0:9084 --handler=web 
user:
	go run app/user/main.go
id:
	go run srv/id/main.go
proto:
	protoc --proto_path=${GOPATH}/src:. --micro_out=. --go_out=. proto/id/snowflake.proto
