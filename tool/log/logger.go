package log

import (
	"context"

	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
	"github.com/uber/jaeger-client-go"
)

//log 中打印当前trace ID,span ID,以及父Span ID
func JaegerLogger(ctx context.Context) *logrus.Entry {
	sp := opentracing.SpanFromContext(ctx)
	if sp != nil {
		if sc, ok := sp.Context().(jaeger.SpanContext); ok {
			return Logger.WithField("traceId", sc.TraceID().String()).
				WithField("spanId", sc.SpanID().String()).
				WithField("parentId", sc.ParentID().String())
		}
	}
	return Logger
}

// 创建tracer 使用
func TracerLogger(trace opentracing.Tracer) *logrus.Entry {
	sp := trace.StartSpan("tracer_logger")
	if sc, ok := sp.Context().(jaeger.SpanContext); ok {
		return Logger.WithField("traceId", sc.TraceID().String()).
			WithField("spanId", sc.SpanID().String()).
			WithField("parentId", sc.ParentID().String())
	}
	return Logger
}

// 使用span
func SpanLogger(sp opentracing.Span) *logrus.Entry {
	if sc, ok := sp.Context().(jaeger.SpanContext); ok {
		return Logger.WithField("traceId", sc.TraceID().String()).
			WithField("spanId", sc.SpanID().String()).
			WithField("parentId", sc.ParentID().String())
	}
	return Logger
}
