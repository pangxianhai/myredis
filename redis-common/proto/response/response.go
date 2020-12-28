package response

import "github.com/gogo/protobuf/proto"

const (
    Success  int32 = 200
    Error    int32 = 500
    NotFound int32 = 404
)

func New(data []byte) *Response {
    return &Response{Code: Success, Msg: "OK", Data: data}
}

func NewOfCode(code int32, msg string) *Response {
    return &Response{Code: code, Msg: msg}
}

func ToByte(res *Response) []byte {
    if res == nil {
        return nil
    }
    b, err := proto.Marshal(res)
    if err != nil {
        panic(err)
    }
    return b
}

func FromByte(b []byte) *Response {
    if b == nil {
        return nil
    }
    res := new(Response)
    err := proto.Unmarshal(b, res)
    if err != nil {
        panic(err)
    }
    return res
}
