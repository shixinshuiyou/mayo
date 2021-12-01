package tracer

import (
	"io"
	"time"

	"github.com/opentracing/opentracing-go"
	"github.com/uber/jaeger-client-go"
	"github.com/uber/jaeger-client-go/config"
)

const JaegerHostPort = "127.0.0.1:6831"

// InitJaegerTracer returns an instance of Jaeger Tracer that samples 100% of traces and logs all spans to stdout.
func InitJaegerTracer(serviceName string) (opentracing.Tracer, io.Closer) {
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

	sender, _ := jaeger.NewUDPTransport(JaegerHostPort, 0)

	reporter := jaeger.NewRemoteReporter(sender)
	// Initialize tracer with a logger and a metrics factory
	tracer, closer, _ := cfg.NewTracer(
		config.Reporter(reporter),
		// config.Logger(log.StdLogger),
		// config.Metrics(metrics.NullFactory),
		// 设置最大 Tag 长度，根据情况设置
		// config.MaxTagValueLength(65535),
	)
	return tracer, closer
}
