//@File     global.go
//@Time     2022/05/12
//@Author   #Suyghur,

package ext

import (
	"github.com/liyue201/gostl/ds/list/simplelist"
	"github.com/liyue201/gostl/ds/map"
	"ylink/comm/model"
)

var (
	// GameVipMap vip玩家信息(GameId,*PlayerInfoMap)
	GameVipMap *treemap.Map
	// CsInfoMap 客服信息(CsId,*CsInfo)
	CsInfoMap *treemap.Map
	// GameOnlinePlayerMap 在线玩家信息
	GameOnlinePlayerMap *treemap.Map
	// GameConnectedMap 已连接客服玩家
	GameConnectedMap *treemap.Map
	// WaitingQueue 玩家等待队列
	WaitingQueue *simplelist.List
)

func GetVipPlayer(gameId, playerId string) *model.PlayerInfo {
	if GameVipMap.Contains(gameId) {
		vipMap := GameVipMap.Get(gameId).(*treemap.Map)
		if vipMap.Contains(playerId) {
			return vipMap.Get(playerId).(*model.PlayerInfo)
		}
	}
	return nil
}

func GetCsInfo(csId string) *model.CsInfo {
	if CsInfoMap.Contains(csId) {
		return CsInfoMap.Get(csId).(*model.CsInfo)
	}
	return nil
}
