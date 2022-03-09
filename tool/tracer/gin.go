package tracer

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/micro/go-micro/v2/metadata"
	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
	"github.com/shixinshuiyou/mayo/tool/log"
)

const contextTracerKey = "Tracer-context"

// Jaeger 通过 middleware 将 tracer 和 ctx 注入到 gin.Context 中
func Jaeger(srvName string) gin.HandlerFunc {
	return func(c *gin.Context) {
		var sp opentracing.Span
		var md = make(map[string]string)
		log.Logger.Debugf("after extract req header :%v", c.Request.Header)
		// 直接从 c.Request.Header 中提取 span,如果没有就新建一个
		spanCtx, _ := opentracing.GlobalTracer().Extract(opentracing.HTTPHeaders, opentracing.HTTPHeadersCarrier(c.Request.Header))
		sp = opentracing.StartSpan(c.Request.URL.Path, opentracing.ChildOf(spanCtx))
		defer sp.Finish()

		sp.Tracer().Inject(sp.Context(), opentracing.TextMap, opentracing.TextMapCarrier(md))
		log.SpanLogger(sp).Debugf("after inject req header :%v", md)

		ctx := context.TODO()
		ctx = opentracing.ContextWithSpan(ctx, sp)
		ctx = metadata.NewContext(ctx, md)
		c.Set(contextTracerKey, ctx)

		c.Next()

		ext.HTTPStatusCode.Set(sp, uint16(c.Writer.Status()))
		ext.HTTPUrl.Set(sp, c.Request.URL.EscapedPath())
		ext.HTTPMethod.Set(sp, c.Request.Method)
		// 把tracer 写入到context中
	}
}

// ContextWithSpan 返回context
func ContextWithSpan(c *gin.Context) (ctx context.Context, ok bool) {
	if v, exist := c.Get(contextTracerKey); exist {
		ctx, ok = v.(context.Context)
		return
	}
	return context.TODO(), false

}
