## Introduction
    a distributed lock implemented by redis written in go.

## Usages
### redis lock
     // mutex lock
     config := []string{":18000", ":18001", ":18002"}
     var options map[string]string
     RegisterLock("redis", config, options)
     // redisLock := GetDefaultRedisMutexLock()
     redisLock := GetRedisMutexLock("test", 10000) //10s
     err := redisLock.Lock()
     if err == nil{
        do something()
        redisLock.Release()
     }

### etcd lock
    // mutext lock
    config := []string{":18000"}
	var options map[string]string
	RegisterLock("etcd", config, options)
	// lock := GetDefaultRedisLock()
	lock := GetEtcdMutexLock("test", 10000)
	fmt.Println(lock)
	err := lock.Lock()
	if err == nil {
		do something()
        lock.Release()
	}    