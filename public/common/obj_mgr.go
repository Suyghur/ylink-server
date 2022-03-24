package logic

import (
	"sync"
)

type ObjMgr struct {
	DictInfo sync.Map
}

func (sel *ObjMgr) Register(key, val interface{}) {
	//log.Printf("<ObjMgr.Register> key:%v, value:%v", key, reflect.TypeOf(val).String())
	sel.DictInfo.Store(key, val)
}

func (sel *ObjMgr) GetObj(objName interface{}) interface{} {
	obj, ok := sel.DictInfo.Load(objName)
	if ok == true {
		return obj
	} else {
		// log.Println("<ObjMgr> GetObj nil, objName:", objName)
		return nil
	}
}

func (sel *ObjMgr) DeleteObj(objName interface{}) {
	sel.DictInfo.Delete(objName)
	//log.Printf("<ObjMgr.DeleteObj> [%v] : %v \n", &sel.DictInfo, objName)
}

func (sel *ObjMgr) GetObjKeys() []interface{} {
	return GetSyncMapKeys(&sel.DictInfo)
}

func (sel *ObjMgr) GetObjValues() []interface{} {
	return GetSyncMapValues(&sel.DictInfo)
}
