package lock

import "testing"
// import "github.com/garyburd/redigo/redis"
import "fmt"
// import "strconv"


func TestRedisLock(t *testing.T)  {
    fmt.Println("test start")
    config := []string{":18000", ":18001", ":18002"}
    RegisterLock("redis", config)
    // lock := GetDefaultRedisLock()
    lock := GetRedisLock("test", 10000)
    fmt.Println(lock)
    err := lock.Lock()
    if err != nil{
        fmt.Println("get lock failed")
    }
    
    fmt.Println("test end")

    
}
