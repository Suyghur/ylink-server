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
	// WaitingList 玩家等待队列
	WaitingList *simplelist.List
)

func GetVipPlayer(gameId, playerId string) *model.PlayerInfo {
	if GameVipMap.Contains(gameId) {
		playerInfoMap := GameVipMap.Get(gameId).(*treemap.Map)
		if playerInfoMap.Contains(playerId) {
			return playerInfoMap.Get(playerId).(*model.PlayerInfo)
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

func GetConnectedPlayerInfo(gameId, playerId string) *model.PlayerInfo {
	if GameConnectedMap.Contains(gameId) {
		connectedMap := GameConnectedMap.Get(gameId).(*treemap.Map)
		if connectedMap.Contains(playerId) {
			return connectedMap.Get(playerId).(*model.PlayerInfo)
		}
	}
	return nil
}

func GetOnlinePlayerInfo(gameId, playerId string) *model.PlayerInfo {
	if GameOnlinePlayerMap.Contains(gameId) {
		onlinePlayerMap := GameOnlinePlayerMap.Get(gameId).(*treemap.Map)
		if onlinePlayerMap.Contains(playerId) {
			return onlinePlayerMap.Get(playerId).(*model.PlayerInfo)
		}
	}
	return nil
}

func GetWaitingPlayerInfo(gameId, playerId string) *model.PlayerInfo {
	for n := WaitingList.FrontNode(); n != nil; n = n.Next() {
		playerInfo := n.Value.(*model.PlayerInfo)
		if playerInfo.GameId == gameId && playerInfo.PlayerId == playerId {
			return playerInfo
		}
	}
	return nil
}
