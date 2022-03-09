package srv

import (
	"github.com/micro/go-micro/v2/client"
	"github.com/shixinshuiyou/mayo/config"
	proto "github.com/shixinshuiyou/mayo/proto/id"
)

/**
 * @Author: czh
 * @Date: 2022/3/9 15:00
 * @Desc: 初始化所有proto文件协议client
 */

var iDService proto.IDService

func InitProtoService(client client.Client) {
	iDService = proto.NewIDService(config.SrvSnowflakeID, client)
}
