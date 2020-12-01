package cmd

import (
    "redis-common/result"
    "strings"
)

type Interpreter interface {
    Interpreter(args []string) *result.Result
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

func Call(args string) *result.Result {
    argv := strings.Split(args, " ")
    key := strings.ToLower(strings.TrimSpace(argv[0]))
    interpreter, ok := factory.interpreters[key]
    if !ok {
        return result.NewOfCode(result.NOT_FOUND, "不支持该命令")
    }
    return interpreter.Interpreter(argv[1:])
}
