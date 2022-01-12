package snowflake

import (
	"context"

	proto "github.com/shixinshuiyou/mayo/proto/id"
	"github.com/shixinshuiyou/mayo/tool/log"
)

type SnowID struct{}

func (s *SnowID) GetSnowflakeID(ctx context.Context, in *proto.SnowIDReq, out *proto.SnowIDResp) error {
	out.ID = getSnowflakeID()
	log.Logger.Debugf("snowflake ID is %d", out.ID)
	return nil
}

func (s *SnowID) ParseSnowflakeID(ctx context.Context, in *proto.PraseIDReq, out *proto.PraseIDResp) error {
	sid := ParseID(uint64(in.ID))
	out.MachineID = int64(sid.MachineID)
	out.Sequence = int64(sid.Sequence)
	out.Timestamp = int64(sid.Timestamp)
	return nil
}

func getSnowflakeID() int64 {
	// TODO 获取当前机器IP名称、或者docker名称
	// 读取当前机器的在数据库表的ID 作为sequence
	return int64(ID())
}
