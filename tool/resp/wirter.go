package resp

import (
	"net/http"
)

// 重写http.ResponseWriter 方法实现自定义返回
const (
	SuccessCode = 200
	FailCode    = 500
)

// 关于重写的一些坑：组合式继承-https://mp.weixin.qq.com/s/maZa_DBGkdAuTt5jaTbzgg
// TODO 可以自定义增加字段
// ResponseStatusWriter net/http status code tracker
type ResponseStatusWriter struct {
	RW         http.ResponseWriter
	StatusCode int
}

func (rsw *ResponseStatusWriter) Header() http.Header {
	return rsw.RW.Header()
}

func (rsw *ResponseStatusWriter) Write(data []byte) (int, error) {
	return rsw.RW.Write(data)
}

func (rsw *ResponseStatusWriter) WriteHeader(statusCode int) {
	rsw.StatusCode = statusCode
	rsw.RW.WriteHeader(statusCode)
}
