package cmd

import (
    "redis-common/proto/request"
    "redis-common/proto/response"
    "strings"
)

type Interpreter interface {
    Interpreter(data []byte) *response.Response
    Key() string
}
type Factory struct {
    interpreters map[string]Interpreter
}

var factory *Factory

func init() {
    factory = &Factory{
        interpreters: make(map[string]Interpreter, 0),
    }
}

func Register(interpreter Interpreter) {
    key := strings.ToLower(strings.TrimSpace(interpreter.Key()))
    factory.interpreters[key] = interpreter
}

func Call(body *request.Request) *response.Response {
    cmd := strings.ToLower(body.Cmd)
    interpreter, ok := factory.interpreters[cmd]
    if !ok {
        return response.NewOfCode(response.NotFound, "不支持该命令")
    }
    return interpreter.Interpreter(body.Data)
}
