package main

import (
    "bytes"
    "encoding/binary"
    "fmt"
)

type CommandInterpreter interface {
    GetCommandName() string
    InterpreterCommand(argv []string)
}

type AbstractCommandInterpreter struct {
}

type SetInterpreter struct {
    AbstractCommandInterpreter
}

func (interpreter *SetInterpreter) GetCommandName() string {
    return "set"
}

func (interpreter *SetInterpreter) InterpreterCommand(argv []string) {
    fmt.Print("执行 set 命令 参数:", argv)
}

type RedisDbService interface {
    InterpreterCommand(name string, argv []string)
}

type RedisDbServiceImpl struct {
    RedisDbService
    // dict 数据
    factory CommandInterpreterFactory
}

func (redisDbService *RedisDbServiceImpl) InterpreterCommand(name string, argv []string) {
    interpreter := redisDbService.factory.InstanceCommandInterpreter(name)
    interpreter.InterpreterCommand(argv)
}

type CommandInterpreterFactory struct {
    interpreterDict map[string]CommandInterpreter
}

func (factory *CommandInterpreterFactory) InitCommandInterpreter() {
    interpreters := make([]CommandInterpreter, 0)
    interpreters = append(interpreters, &SetInterpreter{})

    factory.interpreterDict = make(map[string]CommandInterpreter)
    for _, interpreter := range interpreters {
        factory.interpreterDict[interpreter.GetCommandName()] = interpreter
    }
}

func (factory *CommandInterpreterFactory) InstanceCommandInterpreter(name string) CommandInterpreter {
    return factory.interpreterDict[name]
}

func IntToBytes(n int) []byte {
    data := int64(n)
    bytebuf := bytes.NewBuffer([]byte{})
    binary.Write(bytebuf, binary.BigEndian, data)
    return bytebuf.Bytes()
}

func BytesToInt(bys []byte) int {
    bytebuff := bytes.NewBuffer(bys)
    var data uint8
    binary.Read(bytebuff, binary.BigEndian, &data)
    //println(data)
    return int(data)
}

type SdshdrNode struct {
    buf      []byte
    encoding uint8
    len      []byte
    alloc    []byte
}

type ListNode struct {
    prev  *ListNode
    next  *ListNode
    value interface{}
}

type List struct {
    head *ListNode
    tail *ListNode
    len  *uint64
}

type DictEntry struct {
    key   interface{}
    value interface{}
    next  *DictEntry
}

type DictHt struct {
    size  uint64
    used  uint64
    table []DictEntry
}

func main() {
    b := IntToBytes(1)
    //fmt.Println("1=%X", b)
    //println(len(b))
    //
    b1 := b[4:8]
    fmt.Println("2=%X", b1)
    //println(len(b1))

    //b = []byte{0x01}
    //fmt.Println("3=%X", b)
    n := BytesToInt(b1)
    println(n)

}
