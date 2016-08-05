package lock

import (
    "math/rand"
    "time"
)


type Lock interface {
    Lock() error
    Release() error
}


func RegisterLock(lockType string, config []string)  {
    if lockType == "redis" {
        registRedisInstance(config)
    }
}


func GetDefaultRedisLock() *RedisLock {
    if len(redisConns) == 0{
        panic("You must register redis lock first")
    }
    rand.Seed(time.Now().Unix())
    ramdomNum := rand.Int63()
    return &RedisLock{Name:"defaultRedisLock", Expired:3000, randomKey:ramdomNum}
}

func GetRedisLock(name string, expired int64) *RedisLock {
    if len(redisConns) == 0{
        panic("You must register redis lock first")
    }
    rand.Seed(time.Now().Unix())
    ramdomNum := rand.Int63()
    return &RedisLock{Name:name, Expired:expired, randomKey:ramdomNum}
}
