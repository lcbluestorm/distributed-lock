package lock

import (
    "errors"
    "time"

    "github.com/garyburd/redigo/redis"
)


var redisConns []redis.Conn


func registRedisInstance(addrs []string)  {
    if len(redisConns) > 0{
        return 
    }
    size := len(addrs)
    if size % 2 == 0{
        panic("The redis address(instance) count must be odd")
    }
   
    for i:=0;i<size;i++{
        conn, err := redis.Dial("tcp", addrs[i])
        if err != nil{
            panic(err)
        }
        redisConns = append(redisConns, conn)
    }
    
}


type RedisLock struct {
    Name string
    Expired int64
    randomKey int64
}

func (redisLock RedisLock) Lock() error {
    totalCount := len(redisConns)
    successCount := 0
    beginTime := time.Now()
    for _, conn := range redisConns{
        span := time.Now().Sub(beginTime)
        expired := redisLock.Expired - span.Nanoseconds()/1000000
        ret, err := conn.Do("set", redisLock.Name, redisLock.randomKey, "nx", "px", expired)
        if err == nil{
            if ret == "OK"{
                successCount ++
            }
        }
    }
    if totalCount >> 2 >= successCount{
        err := redisLock.Release()
        if err != nil{
            return err
        }
        return errors.New("Get lock failed(less than half of locks)")
    }
    return nil
}

func (redisLock RedisLock) Release()  error {
    for _, conn := range redisConns{
        ret, err := conn.Do("get", redisLock.Name)
        if err != nil{
            return err
        }
        ret1, err := redis.Int64(ret, nil)
        if err != nil{
            return err
        }
        if ret1 == redisLock.randomKey{
            _, err := conn.Do("del", redisLock.Name)
            if err != nil{
                return err
            }
        }
    }
    return nil
}
