//@File     global.go
//@Time     2022/05/12
//@Author   #Suyghur,

package ext

import (
	treemap "github.com/liyue201/gostl/ds/map"
	"github.com/liyue201/gostl/ds/set"
)

var (
	Game2PlayerMap     *treemap.Map
	CsMap              *treemap.Map
	Game2PlayerStatMap *treemap.Map
	CsStatSet          *set.Set
)
