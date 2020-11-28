package main

import (
    "log"
    "redis-server/biz/net"
    "redis-server/data/config"
)

func main() {
    config.Load("./redis.conf")
    err := net.Start()
    if err != nil {
        log.Fatalln("启动服务失败", err)
    }

}
