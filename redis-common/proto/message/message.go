package message

import (
    "redis-common/bytes"
)

func New(id uint64, content []byte) *Message {
    return &Message{Id: id, Content: content}
}

func ToByte(msg *Message) []byte {
    return bytes.FromPb(msg)
}

// ToPacket 转换为网络通信包 8位长度+数据
func ToPacket(msg *Message) (buf []byte) {
    msgBuf := ToByte(msg)
    buf = bytes.FromUint64(uint64(len(msgBuf)))
    buf = append(buf, msgBuf...)
    return
}

func FromByte(b []byte) *Message {
    p := new(Message)
    bytes.ToPb(b, p)
    return p
}
