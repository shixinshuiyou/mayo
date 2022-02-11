package hystrix

import (
	"net/http"

	"github.com/afex/hystrix-go/hystrix"
	"github.com/shixinshuiyou/mayo/tool/resp"
)

func HystrixHandler(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		name := r.Method + "-" + r.RequestURI

		// 绑定对应command 和对应的熔断、超时控制
		//TODO 后续可以读取配置 实现不同服务接口定制化
		config := hystrix.CommandConfig{
			Timeout: 1000, // 超时时间 ： 超时控制
			// MaxConcurrentRequests: 100,  // 最大请求数 ： 熔断控制
		}
		hystrix.ConfigureCommand(name, config)

		hystrix.Do(name, func() error {
			hw := resp.ResponseStatusWriter{
				RW:         w,
				StatusCode: http.StatusOK,
			}
			h.ServeHTTP(&hw, r)
			return nil
		}, nil)
	})
}
