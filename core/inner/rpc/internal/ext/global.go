//@File     global.go
//@Time     2022/05/12
//@Author   #Suyghur,

package ext

import (
	"github.com/liyue201/gostl/ds/list/simplelist"
	"github.com/liyue201/gostl/ds/map"
)

var (
	Game2PlayerMap     *treemap.Map
	CsMap              *treemap.Map
	Game2PlayerStatMap *treemap.Map
	WaitingQueue       *simplelist.List
)
