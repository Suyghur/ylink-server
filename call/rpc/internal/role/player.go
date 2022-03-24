package role

import "time"

type Player struct {
	/*
		玩家
	*/
	baseRole
	GameId    int32
	IsVisitor bool
	SessionId string
}

func (sel *Player) Init(gameId int32, isVisitor bool, stopChan chan int32) {
	sel.GameId = gameId
	sel.IsVisitor = isVisitor
	sel.connChan = stopChan
	sel.LoginTimeStamp = time.Now().Unix()
	sel.LastTalkTimeStamp = sel.LoginTimeStamp
	sel.SessionId = sel.Id
	sel.logOutChan = make(chan interface{})
}
