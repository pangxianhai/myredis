package cmd

import (
    "redis-common/proto/response"
    "redis-common/proto/str"
    "redis-server/biz/db"
    "redis-server/data/sds"
    "strings"
    "time"
)

func init() {
    set := SetInterpreter{}
    Register(&set)
}

type SetInterpreter struct {
}

func (set *SetInterpreter) Key() string {
    return "set"
}

// set 命令 set key value [EX|PX KEEPTTL] [NX|XX]
// 参数中不含set关键字
// EX-表示秒 PX-表示毫秒 KEEPTTL过期时间 单位为 EX|PX
// NX-当key不存在时才设置值 XX-当key 存在时才设置值
func (set *SetInterpreter) Interpreter(data []byte) *response.Response {
    param := str.SetReqFromByte(data)
    key := sds.NewWithStr(param.Key)
    value := sds.NewWithStr(param.Value)
    res := setStr(key, value, param.Expx, param.Nxxx, int64(param.Timeout))
    return response.New(str.SetResToByte(res))
}

func setStr(key, value *sds.Sds, expx, nxxx string, timeout int64) *str.SetRes {
    if sds.IsEmpty(key) {
        return str.NewSetResInt(0)
    }
    v := db.Get(key)
    nxxx = strings.ToLower(nxxx)
    if nxxx == "nx" {
        if v != nil {
            return str.NewSetResInt(0)
        }
    } else if nxxx == "ex" {
        if v == nil {
            return str.NewSetResInt(0)
        }
    }
    expx = strings.ToLower(expx)
    if expx == "ex" {
        timeout = (time.Now().Unix() + timeout) * 1000
    } else if expx == "px" {
        timeout += time.Now().UnixNano() / 1e6
    }
    db.Put(key, value)
    if len(expx) == 0 {
        db.SetExpire(key, timeout)
    }
    return str.NewSetResInt(1)
}
