//@Author   : KaiShin
//@Time     : 2021/11/1

package handler

import (
	"call_center/call/rpc/internal/core"
	"call_center/call/rpc/pb"
)

func PlayerLogout(server *core.Server, id string) {
	server.KickPlayer(id, int32(pb.ErrorReason_PLAYER_CALL_LOGOUT))
}
