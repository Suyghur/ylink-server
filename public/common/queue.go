package logic

import (
	"container/list"
	"sync"
)

type SyncQueue struct {
	lock sync.RWMutex
	data *list.List
}

func NewQueue() *SyncQueue {
	q := new(SyncQueue)
	q.data = list.New()
	return q
}

func (sel *SyncQueue) GetAll() []interface{} {
	sel.lock.RLock()
	var valList []interface{}
	for e := sel.data.Front(); e != nil; e = e.Next() {
		valList = append(valList, e.Value)
	}
	sel.lock.RUnlock()
	return valList
}

func (sel *SyncQueue) PushBack(v interface{}) int {
	sel.lock.Lock()
	sel.data.PushBack(v)
	sel.lock.Unlock()
	return sel.data.Len()
}

func (sel *SyncQueue) Pop() interface{} {
	sel.lock.Lock()
	e := sel.data.Back()
	resData := sel.data.Remove(e)
	sel.lock.Unlock()
	return resData
}

func (sel *SyncQueue) Back() interface{} {
	sel.lock.RLock()
	e := sel.data.Back()
	sel.lock.RUnlock()
	if e != nil {
		return e
	}
	return nil
}

func (sel *SyncQueue) Len() int {
	sel.lock.RLock()
	l := sel.data.Len()
	sel.lock.RUnlock()
	return l
}

func (sel *SyncQueue) RemoveE(e *list.Element) interface{} {
	sel.lock.Lock()
	res := sel.data.Remove(e)
	sel.lock.Unlock()
	return res
}

func (sel *SyncQueue) Remove(v interface{}) interface{} {
	sel.lock.Lock()
	var rmE *list.Element
	var rmV interface{}
	for e := sel.data.Front(); e != nil; e = e.Next() {
		if e.Value == v {
			rmE = e
			break
		}
	}
	if rmE != nil {
		rmV = sel.data.Remove(rmE)
	}
	sel.lock.Unlock()
	return rmV
}
