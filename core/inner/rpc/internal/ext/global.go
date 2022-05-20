//@File     global.go
//@Time     2022/05/12
//@Author   #Suyghur,

package ext

import (
	"github.com/liyue201/gostl/ds/list/simplelist"
	"github.com/liyue201/gostl/ds/map"
)

var (
	GameVipMap           *treemap.Map
	CsInfoMap            *treemap.Map
	Game2PlayerStatusMap *treemap.Map
	GameConnMap          *treemap.Map
	WaitingQueue         *simplelist.List
)
