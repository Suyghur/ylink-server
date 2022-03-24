package logic

var instance *SingleClassMgr

type SingleClassMgr struct {
	// DictInfo map[string] interface{}
	ObjMgr
}

func init() {
	instance = new(SingleClassMgr)
	// instance.DictInfo = make(map[interface{}]interface{})
}

func GetClassInstance() *SingleClassMgr {
	return instance
}

//func (se *SingleClassMgr) Register(i interface{}){
//	fmt.Println("<SingleClassMgr> Register", reflect.ValueOf(i), reflect.TypeOf(i).String())
//	className := reflect.TypeOf(i).String()
//	se.DictInfo[className] = i
//	fmt.Println(se.DictInfo)
//}
//
//func (se *SingleClassMgr) GetObj(objName string) interface{}{
//	obj, ok := se.DictInfo[objName]
//	if ok == true {
//		return obj
//	} else {
//		fmt.Println("<SingleClassMgr> GetObj nil, objName:", objName)
//		return nil
//	}
//}
