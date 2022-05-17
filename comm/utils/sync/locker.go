//@File     locker.go
//@Time     2022/05/12
//@Author   #Suyghur,

package sync

import gosync "sync"

type Locker interface {
	Lock()
	Unlock()
	RLock()
	RUnlock()
}

var _ Locker = (*gosync.RWMutex)(nil)

type FakeLocker struct {
}

// Lock does nothing
func (l FakeLocker) Lock() {

}

// Unlock does nothing
func (l FakeLocker) Unlock() {

}

// RLock does nothing
func (l FakeLocker) RLock() {

}

// RUnlock does nothing
func (l FakeLocker) RUnlock() {

}
