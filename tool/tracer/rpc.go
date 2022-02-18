package tracer

import (
	"context"
	"fmt"

	"github.com/micro/go-micro/v2/server"
	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
	"github.com/shixinshuiyou/mayo/tool/log"
)

// NewHandlerWrapper accepts an opentracing Tracer and returns a Handler Wrapper
func NewHandlerWrapper() server.HandlerWrapper {
	return func(h server.HandlerFunc) server.HandlerFunc {
		return func(ctx context.Context, req server.Request, rsp interface{}) (err error) {
			err = h(ctx, req, rsp)
			// 获取context携带的值
			md := ContextGetTraceID(ctx)
			name := fmt.Sprintf("%s.%s", req.Service(), req.Endpoint())
			var span opentracing.Span

			log.Logger.Debugf("context value:%v", md)

			spanCtx, _ := opentracing.GlobalTracer().Extract(opentracing.HTTPHeaders, opentracing.TextMapCarrier(md))

			span = opentracing.StartSpan(name, ext.RPCServerOption(spanCtx))
			defer span.Finish()

			if err != nil {
				span.SetTag("msg", err.Error())
				span.SetTag("error", true)
			}

			ext.SpanKindRPCClient.Set(span)
			ext.HTTPMethod.Set(span, req.Endpoint())
			opentracing.GlobalTracer().Inject(span.Context(), opentracing.HTTPHeaders, opentracing.TextMapCarrier(md))
			log.SpanLogger(span).Debugf("after inject req header :%v", req.Header())

			return err
		}
	}
}
