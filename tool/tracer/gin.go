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
		log.Logger.Debugf("after extract req header :%v", ctx.Request.Header)
		// 直接从 c.Request.Header 中提取 span,如果没有就新建一个
		spanCtx, _ := opentracing.GlobalTracer().Extract(opentracing.HTTPHeaders, opentracing.HTTPHeadersCarrier(ctx.Request.Header))
		span = opentracing.StartSpan(ctx.Request.URL.Path, opentracing.ChildOf(spanCtx))

		defer span.Finish()

		// traceID 然后存到 g.ctx 中 供后续使用
		ctx.Set(traceId, getSpanTraceID(span))

		ctx.Next()

		ext.HTTPStatusCode.Set(span, uint16(ctx.Writer.Status()))
		ext.HTTPUrl.Set(span, ctx.Request.URL.EscapedPath())
		ext.HTTPMethod.Set(span, ctx.Request.Method)

		span.Tracer().Inject(span.Context(), opentracing.HTTPHeaders, opentracing.HTTPHeadersCarrier(ctx.Request.Header))
		log.SpanLogger(span).Debugf("after inject req header :%v", ctx.Request.Header)
	}
}
