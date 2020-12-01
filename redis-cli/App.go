package main

import (
    "bufio"
    "fmt"
    "log"
    "os"
    "redis-cli/biz/config"
    "redis-cli/biz/net"
    "redis-common/result"
    "strings"
)

func main() {
    config.Init()
    client, err := net.New()
    if err != nil {
        log.Fatalln("connect to server failed ", err)
    }
    for {
        //读取用户输入
        msg := getInput(client)
        //执行用户输入命令
        reply := handleInput(msg, client)
        //展示命令执行结果
        displayReply(reply)
    }
}

func getInput(client *net.Client) string {
    fmt.Print(client.ServerAddress + " > ")
    reader := bufio.NewReader(os.Stdin)
    msg, err := reader.ReadString('\n')
    if err != nil {
        fmt.Println("your input error", err)
        return ""
    }
    msg = strings.TrimSpace(msg)
    return msg
}

func handleInput(msg string, client *net.Client) string {
    if len(msg) == 0 {
        return ""
    }
    if strings.ToLower(msg) == "exit" {
        os.Exit(0)
    }
    reply, err := client.Send(msg)
    if err != nil {
        fmt.Println("failed", err)
        return ""
    }
    return reply
}

func displayReply(reply string) {
    reply = strings.TrimSpace(reply)
    if len(reply) == 0 {
        return
    }
    res, err := result.FromJson(reply)
    if err != nil {
        fmt.Println("failed", err)
        return
    }
    if res.Code == result.SUCCESS {
        if res.Data == nil {
            fmt.Println(res.Msg)
        } else {
            fmt.Println(res.Data)
        }
    } else {
        fmt.Println(res.Msg)
    }
}
