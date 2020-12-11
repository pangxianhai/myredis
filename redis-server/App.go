package main

import (
    "redis-server/common/logger"
    "redis-server/data/config"
    "redis-server/service/net"
)

func main() {
    config.Load("./redis.conf")
    logger.Init()
    logger.PrintBanner()
    err := net.Start()
    if err != nil {
        logger.Error("redis start failed", err)
    }
}
