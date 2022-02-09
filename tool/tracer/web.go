package tracer

import (
	"net/http"

	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
	"github.com/shixinshuiyou/mayo/tool/log"
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
			log.Logger.Errorf("inject tracer span error : %s", err.Error())
		}

		// 重写http ResponseWriter 方法从而获取执行结果的状态吗
		rsw := &ResponseStatusWriter{w, http.StatusOK}
		h.ServeHTTP(rsw, r)
		log.Logger.Debugf("执行结果是:%d", rsw.StatusCode)

		ext.HTTPMethod.Set(sp, r.Method)
		ext.HTTPUrl.Set(sp, r.URL.EscapedPath())
		ext.HTTPStatusCode.Set(sp, uint16(rsw.StatusCode))
		if rsw.StatusCode >= http.StatusInternalServerError {
			ext.Error.Set(sp, true)
		}
	})
}

// ResponseStatusWriter net/http status code tracker
type ResponseStatusWriter struct {
	w          http.ResponseWriter
	StatusCode int
}

func (rsw *ResponseStatusWriter) Header() http.Header {
	return rsw.w.Header()
}

func (rsw *ResponseStatusWriter) Write(data []byte) (int, error) {
	return rsw.w.Write(data)
}

func (rsw *ResponseStatusWriter) WriteHeader(statusCode int) {
	rsw.StatusCode = statusCode
	rsw.w.WriteHeader(statusCode)
}
