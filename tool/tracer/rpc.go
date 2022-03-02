package tracer

import (
	"context"
	"fmt"

	"github.com/micro/go-micro/v2/server"
	"github.com/opentracing/opentracing-go"
	"github.com/shixinshuiyou/mayo/tool/log"
)

// NewHandlerWrapper accepts an opentracing Tracer and returns a Handler Wrapper
func NewHandlerWrapper() server.HandlerWrapper {
	return func(h server.HandlerFunc) server.HandlerFunc {
		return func(ctx context.Context, req server.Request, rsp interface{}) (err error) {

			name := fmt.Sprintf("%s.%s", req.Service(), req.Endpoint())
			var span opentracing.Span

			span, _ = opentracing.StartSpanFromContext(ctx, name)
			defer span.Finish()

			if err = h(ctx, req, rsp); err != nil {
				span.SetTag("msg", err.Error())
				span.SetTag("error", true)
			}

			log.SpanLogger(span).Debugf("after inject req header :%v", req.Header())

			return err
		}
	}
}
