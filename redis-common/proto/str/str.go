package str

import (
    "redis-common/bytes"
    "strconv"
)

func SetReqToByte(p *SetReq) []byte {
    return bytes.FromPb(p)
}

func SetReqFromByte(b []byte) *SetReq {
    p := new(SetReq)
    bytes.ToPb(b, p)
    return p
}

func SetResToByte(p *SetRes) []byte {
    return bytes.FromPb(p)
}

func SetResFromByte(b []byte) *SetRes {
    p := new(SetRes)
    bytes.ToPb(b, p)
    return p
}

func NewSetRes(data string) *SetRes {
    res := new(SetRes)
    res.Resp = data
    return res
}

func NewSetResInt(data int) *SetRes {
    res := new(SetRes)
    res.Resp = strconv.Itoa(data)
    return res
}

func GetReqToByte(p *GetReq) []byte {
    return bytes.FromPb(p)
}

func GetReqFromByte(b []byte) *GetReq {
    p := new(GetReq)
    bytes.ToPb(b, p)
    return p
}

func GetResToByte(p *GetRes) []byte {
    return bytes.FromPb(p)
}

func GetResFromByte(b []byte) *GetRes {
    p := new(GetRes)
    bytes.ToPb(b, p)
    return p
}

func NewGetRes(v string) *GetRes {
    r := new(GetRes)
    r.Value = v
    return r
}
