package cmd

import (
    "redis-common/proto/response"
    "redis-common/proto/str"
    "redis-server/biz/db"
    "redis-server/data/sds"
)

func init() {
    get := GetInterpreter{}
    Register(get)
}

type GetInterpreter struct {
}

func (get GetInterpreter) Key() string {
    return "get"
}

func (get GetInterpreter) Interpreter(data []byte) *response.Response {
    getReq := str.GetReqFromByte(data)
    obj := db.Get(sds.NewWithStr(getReq.Key))
    if obj == nil {
        getRes := str.NewGetRes("")
        return response.New(str.GetResToByte(getRes))
    }
    if !obj.IsSds() {
        return response.NewOfCode(response.Error, "WRONGTYPE Operation against a key holding the wrong kind of value")
    } else {
        getRes := str.NewGetRes(obj.SdsVal().String())
        return response.New(str.GetResToByte(getRes))
    }
}
