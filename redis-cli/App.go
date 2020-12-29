package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"redis-cli/biz/cmd"
	"redis-cli/biz/config"
	"redis-cli/biz/net"
	"redis-common/proto/request"
	"redis-common/proto/response"
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
		args := getInput(client)
		//执行用户输入命令
		key, reply := handleInput(args, client)
		//展示命令执行结果
		cmd.HandleResult(key, reply, os.Stdout)
	}
}

func getInput(client *net.Client) []string {
	fmt.Print(client.ServerAddress + " > ")
	reader := bufio.NewReader(os.Stdin)
	msg, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("your input error", err)
		return nil
	}
	msg = strings.TrimSpace(msg)
	return strings.Split(msg, " ")
}

func handleInput(args []string, client *net.Client) (string, *response.Response) {
	if len(args) == 0 {
		return "", nil
	}
	key := args[0]
	args = args[1:]
	if strings.ToLower(key) == "exit" {
		os.Exit(0)
	}
	data, err := cmd.HandleInput(key, args)
	if err != nil {
		fmt.Println("(error)", err.Error())
		return key, nil
	}
	req := request.New(key, data)
	reply, err := client.Send(request.ToByte(req))
	if err != nil {
		fmt.Println("failed", err)
		return key, nil
	}
	res := response.FromByte(reply)
	return key, res
}
