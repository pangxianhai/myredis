package cmd

import (
	"io"
	"redis-common/proto/response"
)

type Cmd interface {
	HandleInput(args []string) ([]byte, error)
	HandleResult(res *response.Response, writer io.Writer)
	Name() string
}

type Factory struct {
	cmdMap map[string]Cmd
}

var factory *Factory

func init() {
	factory = &Factory{
		cmdMap: make(map[string]Cmd, 0),
	}
}

func Register(cmd Cmd) {
	factory.cmdMap[cmd.Name()] = cmd
}

func HandleInput(key string, args []string) ([]byte, error) {
	cmd := factory.cmdMap[key]
	if cmd == nil {
		return nil, UnknownCmd(key)
	}
	return cmd.HandleInput(args)
}

func HandleResult(key string, res *response.Response, writer io.Writer) {
	cmd := factory.cmdMap[key]
	if cmd == nil {
		return
	}
	cmd.HandleResult(res, writer)
}
