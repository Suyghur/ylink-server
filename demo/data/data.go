//@Author   : KaiShin
//@Time     : 2022/3/16

package data

import "call_center/call/rpc/call"

const (
	GameId = 1001
)

var (
	PlayerId  = "" // 游客id
	ServiceId = "" // 客服id
	Addr      = "localhost:3000"
)

var Client call.Call
