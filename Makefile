.PHONY: run gateway user id
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
