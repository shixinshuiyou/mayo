package main

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/afex/hystrix-go/hystrix"
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/client"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-plugins/registry/etcdv3/v2"
	"github.com/micro/go-plugins/wrapper/trace/opentracing/v2"
	"github.com/micro/micro/v2/plugin"
	"github.com/shixinshuiyou/framework/tracer"
)

func init() {
	plugin.Register(plugin.NewPlugin(
		plugin.WithName("auth"),
		plugin.WithHandler(func(h http.Handler) http.Handler {
			return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				fmt.Println(r.URL.Path)
			})
		}),
	))

}

func main() {
	srvName := "czh.client.api"
	jaegerTracer, closer := tracer.InitJaegerTracer(srvName)
	defer closer.Close()

	service := micro.NewService(
		micro.Name(srvName),
		micro.Registry(etcdv3.NewRegistry(func(op *registry.Options) {

		})),
		micro.WrapHandler(opentracing.NewHandlerWrapper(jaegerTracer)),
		micro.WrapClient(NewMyClientWrapper()),
		micro.RegisterTTL(time.Second*30),
		micro.RegisterInterval(time.Second*10),
		micro.Address("127.0.0.1:9001"),
	)
	service.Init()
	service.Run()
}

//重新实现熔断方法
type myWrapper struct {
	client.Client
}

func (c *myWrapper) Call(ctx context.Context, req client.Request, rsp interface{}, opts ...client.CallOption) error {
	name := req.Service() + "." + req.Endpoint()
	return hystrix.Do(name, func() error {
		return c.Client.Call(ctx, req, rsp, opts...)
	}, func(e error) error {
		//这里是处理服务降级的地方
		fmt.Println("进入服务降级")
		return nil
	})
}

// NewClientWrapper returns a hystrix client Wrapper.
func NewMyClientWrapper() client.Wrapper {
	return func(c client.Client) client.Client {
		return &myWrapper{c}
	}
}
