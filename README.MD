## Introduction
    a distributed lock implemented by redis written in go.

## Usages
### redis lock
     // mutex lock
     config := []string{":18000", ":18001", ":18002"}
     RegisterLock("redis", config)
     // redisLock := GetDefaultRedisMutexLock()
     redisLock := GetRedisMutexLock("test", 10000) //10s
     err := redisLock.Lock()
     if err != nil{
        fmt.Println("get lock failed")
     }
     do something
     redisLock.Release()
     

### etcd lock
    
    