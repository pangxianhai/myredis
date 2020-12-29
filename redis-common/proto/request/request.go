package request

import (
    "redis-common/bytes"
)

func ToByte(p *Request) []byte {
    return bytes.FromPb(p)
}

func FromByte(b []byte) *Request {
    request := new(Request)
    bytes.ToPb(b, request)
    return request
}

func New(cmd string, data []byte) *Request {
    request := new(Request)
    request.Cmd = cmd
    request.Data = data
    return request
}
