//@File     visitor.go
//@Time     2022/05/12
//@Author   #Suyghur,

package visitor

type Visitor func(value interface{}) bool

type KvVisitor func(key, value interface{}) bool
