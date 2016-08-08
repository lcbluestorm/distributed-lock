package lock

import (
	"math/rand"
	"time"
)

type Lock interface {
	Lock() error
	Release() error
}

func RegisterLock(lockType string, addrs []string, options map[string]string) {
	if lockType == "redis" {
		registRedisInstance(addrs, options)
	} else if lockType == "etcd" {
		registEtcdInstance(addrs, options)
	}
}

func GetDefaultRedisMutexLock() *RedisMutexLock {
	if len(redisConns) == 0 {
		panic("You must register redis instance first")
	}
	rand.Seed(time.Now().Unix())
	ramdomNum := rand.Int63()
	return &RedisMutexLock{Name: "defaultRedisMutexLock", Expired: 3000, randomKey: ramdomNum}
}

func GetRedisMutexLock(name string, expired int64) *RedisMutexLock {
	if len(redisConns) == 0 {
		panic("You must register redis instance first")
	}
	rand.Seed(time.Now().Unix())
	ramdomNum := rand.Int63()
	return &RedisMutexLock{Name: name, Expired: expired, randomKey: ramdomNum}
}

// func GetDefaultRedisRWLock()  {

// }

func GetDefaultEtcdMutexLock() *EtcdMutexLock {
	if ectdClient == nil {
		panic("You must register etct instance first")
	}
	return &EtcdMutexLock{Name: "defaultEtcdMutexLock", Expired: 3000}
}

func GetEtcdMutexLock(name string, expired int64) *EtcdMutexLock {
	if ectdClient == nil {
		panic("You must register etct instance first")
	}
	return &EtcdMutexLock{Name: name, Expired: expired}
}
