package main

import (
	"fmt"
	"github.com/carmark/pseudo-terminal-go/terminal"
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
	term, err := terminal.NewWithStdInOut()
	if err != nil {
		log.Fatalln(err)
	}
	defer term.ReleaseFromStdInOut()
	for {
		//读取用户输入
		args, err := getInput(term, client)
		if err != nil {
			break
		}
		//执行用户输入命令
		key, reply := handleInput(args, client)
		//展示命令执行结果
		if reply != nil {
			cmd.HandleResult(key, reply, term)
		}
	}
}

func getInput(term *terminal.Terminal, client *net.Client) ([]string, error) {
	term.SetPrompt(client.ServerAddress + " > ")
	msg, err := term.ReadLine()
	if err != nil {
		return nil, err
	}
	msg = strings.TrimSpace(msg)
	return strings.Split(msg, " "), nil
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
		fmt.Println("(errors)", err)
		return key, nil
	}
	res := response.FromByte(reply)
	return key, res
}
