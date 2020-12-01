package main

import (
    "log"
    "redis-server/data/config"
    "redis-server/service/net"
)

func main() {
    config.Load("./redis.conf")
    err := net.Start()
    if err != nil {
        log.Fatalln("启动服务失败", err)
    }

}
