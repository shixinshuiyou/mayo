package tracer

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
)

// Jaeger 通过 middleware 将 tracer 和 ctx 注入到 gin.Context 中
func Jaeger(srvName, jaegerAddress string) gin.HandlerFunc {
	return func(c *gin.Context) {
		var curSpan opentracing.Span
		jaegerTracer, closer, _ := InitJaegerTracer(srvName, jaegerAddress)
		defer closer.Close()
		// 直接从 c.Request.Header 中提取 span,如果没有就新建一个
		spCtx, err := opentracing.GlobalTracer().Extract(opentracing.HTTPHeaders, opentracing.HTTPHeadersCarrier(c.Request.Header))
		if err != nil {
			curSpan = jaegerTracer.StartSpan(c.Request.URL.Path)
			defer curSpan.Finish()
		} else {
			curSpan = opentracing.StartSpan(
				c.Request.URL.Path,
				opentracing.ChildOf(spCtx),
				opentracing.Tag{Key: string(ext.Component), Value: "HTTP"},
				ext.SpanKindRPCServer,
			)
			defer curSpan.Finish()
		}
		// 然后存到 g.ctx 中 供后续使用
		c.Set("tracer", jaegerTracer)
		c.Set("ctx", opentracing.ContextWithSpan(context.Background(), curSpan))
		c.Next()
	}
}
