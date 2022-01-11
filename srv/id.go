package srv

import (
	"context"

	proto "github.com/shixinshuiyou/mayo/proto/id"
	"github.com/shixinshuiyou/mayo/tool/log"
)

type SnowID struct {
	TraceContext context.Context
}

func NewIDSrv(ctx context.Context) (sid *SnowID) {
	sid = new(SnowID)
	sid.TraceContext = ctx
	return
}

func (sid *SnowID) GetSnowflakeID() (int64, error) {
	resp, err := idService.GetSnowflakeID(sid.TraceContext, &proto.SnowIDReq{})
	if err != nil {
		log.Logger.Errorf("get snowflakeID err:%s", err)
	}
	return resp.ID, nil
}
