package snowflake

import (
	"context"

	"github.com/shixinshuiyou/mayo/proto"
)

type SnowID struct{}

func (s *SnowID) GetSnowflakeID(ctx context.Context, in *proto.SnowIDReq, out *proto.SnowIDResp) error {
	out.ID = int64(ID())
	return nil
}

func (s *SnowID) ParseSnowflakeID(ctx context.Context, in *proto.PraseIDReq, out *proto.PraseIDResp) error {
	sid := ParseID(uint64(in.ID))
	out.MachineID = int64(sid.MachineID)
	out.Sequence = int64(sid.Sequence)
	out.Timestamp = int64(sid.Timestamp)
	return nil
}
