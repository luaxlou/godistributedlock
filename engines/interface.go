package engines

import "time"

type DistributedLockEngine interface {
	GetLock(lockKey string, expires time.Duration) (bool, error)
	ReleaseLock(lockKey string) error
}
