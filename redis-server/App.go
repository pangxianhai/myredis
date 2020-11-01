package main

import (
    "fmt"
    "reids-server/data/config"
)

func main() {
    config.InitConfig("./redis.conf")
    conf := config.GetConfigInstance()
    fmt.Println(conf.GetServerPort())

}
