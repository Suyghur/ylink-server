//@File     globalkey.go
//@Time     2022/05/13
//@Author   #Suyghur,

package globalkey

const (
	ConnectTypeError        = -1
	ConnectTypeNormalPlayer = 0
	ConnectTypeVipPlayer    = 1
	ConnectTypeCs           = 2
)

const (
	EventRemoveTimeoutJob     = "EventRemoveTimeoutJob"
	EventHandleRmqJob         = "EventHandleRmqJob"
	EventUnsubscribeRmqJob    = "EventUnsubscribeRmq"
	EventNotifyUserOfflineJob = "EventNotifyUserOfflineJob"
)

const (
	All             = "all"
	AllPlayer       = "all_player"
	AllVipPlayer    = "all_vip_player"
	AllNormalPlayer = "all_normal_player"
	AllCs           = "all_cs"
)
