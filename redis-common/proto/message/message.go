package message

import "github.com/gogo/protobuf/proto"

func New(id uint64, content []byte) *Message {
    return &Message{Id: id, Content: content}
}

func ToByte(msg *Message) []byte {
    if msg == nil {
        return nil
    }
    b, err := proto.Marshal(msg)
    if err != nil {
        panic(err)
    }
    return b
}

func FromByte(b []byte) *Message {
    if b == nil {
        return nil
    }
    p := new(Message)
    err := proto.Unmarshal(b, p)
    if err != nil {
        panic(err)
    }
    return p
}
