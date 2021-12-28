# api 
go run main.go api --namespace=czh.micro.api --enable_cors=false --address 0.0.0.0:9084 --handler=web
# user
go run app/user/main.go