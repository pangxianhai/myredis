package request

import "github.com/gogo/protobuf/proto"

func ToByte(p *Request) []byte {
    if p == nil {
        return nil
    }
    b, err := proto.Marshal(p)
    if err != nil {
        panic(err)
    }
    return b
}

func FromByte(b []byte) *Request {
    if b == nil {
        return nil
    }
    request := new(Request)
    err := proto.Unmarshal(b, request)
    if err != nil {
        panic(err)
    }
    return request
}
