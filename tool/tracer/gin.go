package tracer

import (
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
				ext.RPCServerOption(spCtx),
			)
			defer curSpan.Finish()
		}
		// 然后存到 g.ctx 中 供后续使用
		ext.SpanKindRPCClient.Set(curSpan)
		ext.HTTPUrl.Set(curSpan, c.Request.RequestURI)
		ext.HTTPMethod.Set(curSpan, c.Request.Method)
		jaegerTracer.Inject(curSpan.Context(), opentracing.HTTPHeaders, opentracing.HTTPHeadersCarrier(c.Request.Header))

		c.Next()

	}
}
