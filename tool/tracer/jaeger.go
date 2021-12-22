package tracer

import (
	"io"
	"time"

	"github.com/opentracing/opentracing-go"
	"github.com/uber/jaeger-client-go"
	"github.com/uber/jaeger-client-go/config"
)

// InitJaegerTracer returns an instance of Jaeger Tracer
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
		// config.Logger(log.StdLogger),
		// config.Metrics(metrics.NullFactory),
		// 设置最大 Tag 长度，根据情况设置
		// config.MaxTagValueLength(65535),
	)
	return tracer, closer, err
}
