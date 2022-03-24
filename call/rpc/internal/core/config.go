package core

type Config struct {
	HeartBeatInterval     int32 // 心跳间隔
	WaitConnServiceLimit  int64 // 等待队列时长
	LastTalkIntervalLimit int64 // 发言超时间隔
}
