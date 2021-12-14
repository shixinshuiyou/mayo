module github.com/shixinshuiyou/mayo

go 1.14

require (
	github.com/gin-gonic/gin v1.7.7
	github.com/go-gomail/gomail v0.0.0-20160411212932-81ebce5c23df
	github.com/go-redis/redis v6.15.9+incompatible
	github.com/micro/go-micro/v2 v2.9.1
	github.com/micro/go-plugins/registry/etcdv3/v2 v2.0.0-20201212203005-87d43a62e4ad
	github.com/micro/go-plugins/wrapper/trace/opentracing/v2 v2.9.1
	github.com/micro/micro/v2 v2.9.3
	github.com/shixinshuiyou/framework v0.0.0-20211203022543-61a68e71dad9
	gopkg.in/alexcesaro/quotedprintable.v3 v3.0.0-20150716171945-2caba252f4dc // indirect
	gopkg.in/gomail.v2 v2.0.0-20160411212932-81ebce5c23df // indirect
)

replace (
	github.com/cloudflare/cloudflare-go => github.com/cloudflare/cloudflare-go v0.10.2
	google.golang.org/grpc => google.golang.org/grpc v1.26.0
)
