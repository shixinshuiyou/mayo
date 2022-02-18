package tracer

import (
	"context"
	"io"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/opentracing/opentracing-go"
	mylog "github.com/shixinshuiyou/mayo/tool/log"
	"github.com/uber/jaeger-client-go"
	"github.com/uber/jaeger-client-go/config"
	"github.com/uber/jaeger-client-go/log"
)

type contextKey struct{}

var (
	contextTracerKey = contextKey{}
	uberTrace        = "Uber-Trace-Id"
	traceId          = "trace-id"
)

// InitJaegerTracer 返回 Jaeger Tracer
func InitJaegerTracer(serviceName, jaegerAddr string) (opentracing.Tracer, io.Closer, error) {
	cfg := &config.Configuration{
		ServiceName: serviceName,
		Sampler: &config.SamplerConfig{
			Type:  jaeger.SamplerTypeConst,
			Param: 1,
		},
		Reporter: &config.ReporterConfig{
			LogSpans:            true,
			BufferFlushInterval: 1 * time.Second,
			// LocalAgentHostPort:  JaegerHostPort,
		},
	}

	sender, err := jaeger.NewUDPTransport(jaegerAddr, 0)
	if err != nil {
		return nil, nil, err
	}

	reporter := jaeger.NewRemoteReporter(sender)
	// Initialize tracer with a logger and a metrics factory
	tracer, closer, err := cfg.NewTracer(
		config.Reporter(reporter),
		config.Logger(log.StdLogger),
		// config.Metrics(metrics.NullFactory),
		// 设置最大 Tag 长度，根据情况设置
		// config.MaxTagValueLength(65535),
	)
	return tracer, closer, err
}

// SetJaegerGlobalTracer 设置默认全局tracer
func SetJaegerGlobalTracer(serviceName, jaegerAddr string) (opentracing.Tracer, io.Closer, error) {
	tracer, closer, err := InitJaegerTracer(serviceName, jaegerAddr)
	if err != nil {
		return nil, nil, err
	}
	opentracing.SetGlobalTracer(tracer)
	return tracer, closer, err
}

// ContextWithTraceID  context 携带 traceID
func ContextWithTraceID(c *gin.Context) (ctx context.Context) {
	md := make(map[string]string)
	id, _ := c.Get(traceId)
	md[uberTrace] = id.(string)
	ctx = context.WithValue(c, contextTracerKey, md)
	return ctx
}

// 获取Context中携带的traceID
func ContextGetTraceID(c context.Context) map[string]string {
	val := c.Value(contextTracerKey)
	mylog.Logger.Debugf("context with values is:%v", val)
	md, ok := val.(map[string]string)
	if ok {
		return md
	}
	return make(map[string]string)
}

func getSpanTraceID(sp opentracing.Span) string {
	sc, _ := sp.Context().(jaeger.SpanContext)
	return sc.TraceID().String()
}
