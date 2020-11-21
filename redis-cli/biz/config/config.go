package config

import (
    "flag"
    "strconv"
)

type Config struct {
    ServerAddress string
    Port          int
}

var config *Config

func Init() {
    config = new(Config)
    flag.StringVar(&config.ServerAddress, "h", "127.0.0.1", "服务器地址")
    flag.IntVar(&config.Port, "p", 6379, "端口号")
    flag.Parse()
}

func ServerAddress() string {
    return config.ServerAddress + ":" + strconv.Itoa(config.Port)
}
