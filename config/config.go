package config

const (
	SrvApiGateAway = "czh.micro.api"
	SrvActionName  = "czh.micro.api.user"
)

var (
	JaegerAddress string
	EtcdAddress   string
	// Once          sync.Once
)

// TODO  读取配置文件
func init() {
	JaegerAddress = "127.0.0.1:6831"
	EtcdAddress = "127.0.0.1:2380"
}
