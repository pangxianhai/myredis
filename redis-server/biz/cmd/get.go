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
    value := db.Get(sds.NewWithStr(getReq.Key))
    v := value.(sds.Sds)
    getRes := str.NewGetRes(v.String())
    return response.New(str.GetResToByte(getRes))
}
