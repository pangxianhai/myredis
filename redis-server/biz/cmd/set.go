package cmd

import (
    "log"
    "redis-common/result"
)

func init() {
    set := SetInterpreter{}
    Register(set)
}

type SetInterpreter struct {
}

func (set SetInterpreter) Key() string {
    return "set"
}

func (set SetInterpreter) Interpreter(args []string) *result.Result {
    log.Println("set 参数:", args)
    return result.NewOfData("执行 SET 命令")
}
