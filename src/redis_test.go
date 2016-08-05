package lock

import "testing"
import "github.com/garyburd/redigo/redis"
import "fmt"
// import "strconv"


func TestRedis(t *testing.T)  {
    conn, err := redis.Dial("tcp", ":6379")
    if err != nil{
            panic(err)
        }
    var value int64 = 33
    n, _ := conn.Do("set", "test11", value, "nx", "px", 3000)
    if n == "OK"{
        // fmt.Printf("yes")
    }else {
        fmt.Print("no")
    }
    n1, _ := conn.Do("get", "test11")
ret, _ := redis.Int64(n1, nil)
    fmt.Printf("%T", ret)
    fmt.Printf("%s", ret)
    
}