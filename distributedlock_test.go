package godistributedlock

import (
	"github.com/luaxlou/godistributedlock/engines/redisengine"
	"log"
	"testing"
	"time"
)

func getTestLock() *DistributedLock {
	engine, err := redisengine.New("127.0.0.1:6379", "", 0)

	if err != nil {
		panic(err.Error())

	}

	return New(engine)
}

func TestDistributedLock_RunIfGetLock(t *testing.T) {

	lock := getTestLock()

	ok, err := lock.RunIfGetLock("test_key", time.Second*10, func() error {
		log.Println("Get Locked")

		return nil
	})
	log.Println(ok, err)

}

func TestDistributedLock_RunWaitForLock(t *testing.T) {
	lock := getTestLock()

	ok, err := lock.RunWaitForLock("test_key", time.Second*10, func() error {
		log.Println("Get Locked")

		return nil
	})
	log.Println(ok, err)
}
