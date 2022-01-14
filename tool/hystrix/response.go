package hystrix

import (
	"net/http"

	"github.com/shixinshuiyou/mayo/tool/log"
)

// 重写http.ResponseWriter 方法实现自定义返回
const (
	SuccessCode = 200
	FailCode    = 500
)

// 关于重写的一些坑：组合式继承-https://mp.weixin.qq.com/s/maZa_DBGkdAuTt5jaTbzgg
type HystrixResponse struct {
	http.ResponseWriter
	Status int
	End    bool
}

func (hyr *HystrixResponse) Header() http.Header {
	return hyr.ResponseWriter.Header()
}

func (hyr *HystrixResponse) WriteHeader(statusCode int) {
	if hyr.End {
		return
	}
	hyr.Status = statusCode
	hyr.ResponseWriter.WriteHeader(statusCode)
}

// 重写最终接口返回的数据
// TODO 可以自定义增加字段
func (hyr *HystrixResponse) Write(data []byte) (int, error) {
	log.Logger.Debugf("hystrixResponse wirte data:%s", string(data))
	if hyr.End {
		return 0, nil
	}
	hyr.End = true
	return hyr.ResponseWriter.Write(data)
}
