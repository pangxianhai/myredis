package message

import (
    "errors"
    "redis-common/bytes"
)

type Message struct {
    ID      uint64 // 消息ID
    Content string // 消息内容
}

func New(id uint64, content string) *Message {
    return &Message{ID: id, Content: content}
}

// ToByte message 转换为二进制数据
func ToByte(msg *Message) (buf []byte) {
    buf = make([]byte, 0)
    b := bytes.FromUint64(msg.ID)
    buf = append(buf, b...)
    buf = append(buf, []byte(msg.Content)...)
    return
}

// ToPacket 转换为网络通信包 8位长度+数据
func ToPacket(msg *Message) (buf []byte) {
    msgBuf := ToByte(msg)
    buf = bytes.FromUint64(uint64(len(msgBuf)))
    buf = append(buf, msgBuf...)
    return
}

// FromByte  二进制数据转换为 message
func FromByte(buf []byte) (msg *Message, err error) {
    msg = new(Message)
    if len(buf) < 8 {
        err = errors.New("数据异常不能转换")
        return
    }
    msg.ID = bytes.ToUint64(buf[0:8])
    //更新剩余数据
    buf = buf[8:]
    msg.Content = string(buf)
    return
}
