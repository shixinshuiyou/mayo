package tracer

import (
	"github.com/gin-gonic/gin"
	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
	"github.com/shixinshuiyou/mayo/tool/log"
)

// Jaeger 通过 middleware 将 tracer 和 ctx 注入到 gin.Context 中
func Jaeger(srvName string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var span opentracing.Span
		// 直接从 c.Request.Header 中提取 span,如果没有就新建一个
		spanCtx, _ := opentracing.GlobalTracer().Extract(opentracing.HTTPHeaders, opentracing.HTTPHeadersCarrier(ctx.Request.Header))
		span = opentracing.StartSpan(ctx.Request.URL.Path, ext.RPCServerOption(spanCtx))
		// span, _ = opentracing.StartSpanFromContext(ctx, ctx.Request.URL.Path, opentracing.Tag{Key: string(ext.Component), Value: "HTTP"})
		defer span.Finish()
		log.SpanLogger(span).Debug()

		// 然后存到 g.ctx 中 供后续使用
		ext.SpanKindRPCClient.Set(span)
		ext.HTTPUrl.Set(span, ctx.Request.RequestURI)
		ext.HTTPMethod.Set(span, ctx.Request.Method)

		span.Tracer().Inject(span.Context(), opentracing.HTTPHeaders, opentracing.HTTPHeadersCarrier(ctx.Request.Header))

		ctx.Next()
	}
}
