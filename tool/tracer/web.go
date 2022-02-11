package tracer

import (
	"net/http"

	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
	"github.com/shixinshuiyou/mayo/tool/log"
	"github.com/shixinshuiyou/mayo/tool/resp"
)

// TracerWrapper tracer wrapper
func TracerWrapper(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		spanCtx, _ := opentracing.GlobalTracer().Extract(opentracing.HTTPHeaders, opentracing.HTTPHeadersCarrier(r.Header))
		sp := opentracing.GlobalTracer().StartSpan(r.URL.Path, opentracing.ChildOf(spanCtx))
		defer sp.Finish()

		if err := opentracing.GlobalTracer().Inject(
			sp.Context(),
			opentracing.HTTPHeaders,
			opentracing.HTTPHeadersCarrier(r.Header)); err != nil {
			log.SpanLogger(sp).Errorf("inject tracer span error : %s", err.Error())
		}

		// 重写http ResponseWriter 方法从而获取执行结果的状态吗
		rsw := resp.ResponseStatusWriter{
			RW:         w,
			StatusCode: http.StatusOK,
		}
		h.ServeHTTP(&rsw, r)
		log.SpanLogger(sp).Debugf("执行结果是:%d", rsw.StatusCode)

		ext.HTTPMethod.Set(sp, r.Method)
		ext.HTTPUrl.Set(sp, r.URL.EscapedPath())
		ext.HTTPStatusCode.Set(sp, uint16(rsw.StatusCode))
		if rsw.StatusCode >= http.StatusInternalServerError {
			ext.Error.Set(sp, true)
		}
	})
}
