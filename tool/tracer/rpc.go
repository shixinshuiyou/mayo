package tracer

import (
	"context"
	"fmt"

	"github.com/micro/go-micro/v2/server"
	"github.com/opentracing/opentracing-go"
)

// NewHandlerWrapper accepts an opentracing Tracer and returns a Handler Wrapper
func NewHandlerWrapper(ot opentracing.Tracer) server.HandlerWrapper {
	return func(h server.HandlerFunc) server.HandlerFunc {
		return func(ctx context.Context, req server.Request, rsp interface{}) error {
			var span opentracing.Span
			if ot == nil {
				ot = opentracing.GlobalTracer()
			}
			name := fmt.Sprintf("%s.%s", req.Service(), req.Endpoint())
			span, _ = opentracing.StartSpanFromContext(ctx, name)
			defer span.Finish()
			return nil
		}
	}
}
