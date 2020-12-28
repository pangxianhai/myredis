package main

import (
    "redis-server/biz/db"
    "redis-server/common/config"
    "redis-server/common/logger"
    "redis-server/service/net"
)

func main() {
    config.Load("./redis.conf")
    logger.Init()
    logger.PrintBanner()
    db.Init()
    err := net.Start()
    if err != nil {
        logger.Error("redis start failed", err)
    }
}
