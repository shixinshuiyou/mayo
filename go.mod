module github.com/shixinshuiyou/mayo

go 1.14

require (
	github.com/HdrHistogram/hdrhistogram-go v1.1.2 // indirect
	github.com/gin-gonic/gin v1.7.7
	github.com/go-gomail/gomail v0.0.0-20160411212932-81ebce5c23df
	github.com/micro/go-micro/v2 v2.9.1
	github.com/micro/go-plugins/registry/etcdv3/v2 v2.0.0-20201212203005-87d43a62e4ad
	github.com/micro/go-plugins/wrapper/trace/opentracing/v2 v2.9.1
	github.com/micro/micro/v2 v2.9.3
	github.com/opentracing/opentracing-go v1.2.0
	github.com/prometheus/client_golang v1.11.0 // indirect
	github.com/sirupsen/logrus v1.8.1
	github.com/uber/jaeger-client-go v2.30.0+incompatible
	github.com/uber/jaeger-lib v2.4.1+incompatible // indirect
	gopkg.in/alexcesaro/quotedprintable.v3 v3.0.0-20150716171945-2caba252f4dc // indirect
	gopkg.in/gomail.v2 v2.0.0-20160411212932-81ebce5c23df // indirect
)

replace (
	github.com/cloudflare/cloudflare-go => github.com/cloudflare/cloudflare-go v0.10.2
	google.golang.org/grpc => google.golang.org/grpc v1.26.0
)
