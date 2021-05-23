package message

import (
	"github.com/cherrison/cherrykv/proto/pkg/raft_cmdpb"
	"github.com/cherrison/cherrykv/proto/pkg/raft_serverpb"
)

type RaftRouter interface {
	Send(regionID uint64, msg Msg) error
	SendRaftMessage(msg *raft_serverpb.RaftMessage) error
	SendRaftCommand(req *raft_cmdpb.RaftCmdRequest, cb *Callback) error
}
