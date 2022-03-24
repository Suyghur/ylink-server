package role

import (
	"call_center/call/rpc/pb"
	db "call_center/db/rpc/pb"
	"log"
	"time"
)

type Service struct {
	/*
		客服
	*/
	baseRole
	HangUpList []*pb.IdInfo
}

func (sel *Service) Init(stopChan chan int32) {
	sel.connChan = stopChan
	sel.LoginTimeStamp = time.Now().Unix()
	sel.LastTalkTimeStamp = sel.LoginTimeStamp
	sel.logOutChan = make(chan interface{})
}

func (sel *Service) InitHandUpList(pbList []*db.DbChatRecord) {
	for _, data := range pbList {
		sel.HangUpList = append(sel.HangUpList, &pb.IdInfo{Id: data.PlayerId, GameId: data.GameId})
	}
	log.Println("<Service.InitHanUpList> len:", len(sel.HangUpList), " id:", sel.Id)
}
