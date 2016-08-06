package lock


import (
    "time"
    // "unicode"

    "github.com/coreos/etcd/client"
)


var ectdClient client.Client


func registEtcdInstance(addrs []string, options map[string]string)  {
    if ectdClient != nil{
        return
    }
    transportType, ok := options["transport"]
    transport := client.DefaultTransport
    switch transportType {
    case "type1":     
    }
    timeoutStr, ok := options["timeout"]
    timeout := time.Second*30
    if ok {
        parsedTimeout, err := time.ParseDuration(timeoutStr)
        if err == nil{
            timeout = parsedTimeout
        }else {
            panic("the timeout value is err(should be 1s, 1us, 1ns, 1ms, 1m, 1h")
        }
    }
    cfg := client.Config{
        Endpoints: addrs, // []string{"http://127.0.0.1:2379"},
        Transport: transport,
        HeaderTimeoutPerRequest: timeout}
    var err error = nil
    ectdClient, err = client.New(cfg)
    if err != nil{
        panic(err)
    }
}

type EtcdMutexLock struct {
    
}