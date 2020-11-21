package main

import (
    "fmt"
    "log"
    "os"
    "redis-cli/biz/config"
    "redis-cli/biz/net"
    "strings"
)

func main() {
    config.Init()
    err := net.Init()
    if err != nil {
        log.Fatalln("connect to server failed ", err)
    }

    for {
        var msg string
        fmt.Print(config.ServerAddress() + " > ")
        _, err := fmt.Scan(&msg)
        if err != nil {
            fmt.Println("your input error", err)
            continue
        }
        msg = strings.TrimSpace(msg)
        if strings.ToLower(msg) == "exit" {
            os.Exit(0)
        }
        reply, err := net.Send(msg)
        if err != nil {
            fmt.Println("failed", err)
            continue
        } else {
            fmt.Println(reply)
        }
    }
}
