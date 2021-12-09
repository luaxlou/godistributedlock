package godistributedlock

import (
	"errors"
	"time"
)

type DistributedLock struct {
	engine DistributedLockEngine
}

func New(engine DistributedLockEngine) *DistributedLock {
	if engine == nil {
		panic("Engine must not nil")
	}

	return &DistributedLock{engine: engine}
}

func (d *DistributedLock) RunIfGetLock(lockKey string, expires time.Duration, exec func() error) (getLocked bool, err error) {

	getLocked, err = d.engine.GetLock(lockKey, expires)

	if err != nil {
		return false, err
	}

	if getLocked {

		err = exec()

		d.engine.ReleaseLock(lockKey)

		return
	}

	return false, nil
}

func (d *DistributedLock) RunWaitForLock(lockKey string, expires time.Duration, exec func() error) (getLocked bool, err error) {

	//重试10次

	for i := 0; i < 10; i++ {
		getLocked, err = d.RunIfGetLock(lockKey, expires, exec)
		if err != nil {
			return
		}

		if getLocked {
			return
		}

		time.Sleep(time.Second)

	}

	return false, errors.New("Get lock timeout:" + lockKey)

}
