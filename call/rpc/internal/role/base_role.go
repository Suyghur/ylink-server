package role

import (
	"log"
	"time"
)

type baseRole struct {
	/*
		基础角色
	*/
	Id                string
	LoginTimeStamp    int64 // 登录时间戳
	LastTalkTimeStamp int64 // 最近发言时间戳

	connChan chan int32  // 连接channel，释放则断开连接
	Stream   interface{} // 连接stream流

	logOutChan chan interface{}
}

func (sel *baseRole) SetChan(connChan chan int32) {
	/*
		设置连接信号
	*/
	sel.connChan = connChan
}

func (sel *baseRole) RefreshTalkTimeStamp() {
	/*
		刷新发言时间
	*/
	nowStamp := time.Now().Unix()
	sel.LastTalkTimeStamp = nowStamp
}

func (sel *baseRole) StopChan(reason int32) {
	/*
		发送信号：断开连接
	*/
	if sel.connChan == nil {
		return
	}
	log.Printf("<baseRole.StopChan> id:%s, reason:%d \n", sel.Id, reason)
	sel.connChan <- reason
	close(sel.connChan)
	sel.connChan = nil
}

func (sel *baseRole) WaitLogOut() <-chan interface{} {
	return sel.logOutChan
}

func (sel *baseRole) Final() {
	sel.logOutChan <- true
	log.Println("<baseRole.Final> id:", sel.Id)
}
