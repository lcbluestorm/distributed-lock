package lock

import (
	"fmt"
	"testing"
)

func TestEtcd(t *testing.T) {
	fmt.Println("etcd mutex lock test start")
	config := []string{":18000", ":18001", ":18002"}
	var options map[string]string
	RegisterLock("etcd", config, options)
	// lock := GetDefaultRedisLock()
	lock := GetEtcdMutexLock("test", 10000)
	fmt.Println(lock)
	err := lock.Lock()
	if err != nil {
		fmt.Println("get lock failed")
	}

	fmt.Println("etcd mutex locktest end")

}
