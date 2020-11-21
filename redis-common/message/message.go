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
func ToByte(msg *Message) (buf []byte, err error) {
    buf = make([]byte, 0)
    b, err := bytes.FromUint64(msg.ID)
    if err != nil {
        return
    }
    buf = append(buf, b...)
    buf = append(buf, []byte(msg.Content)...)
    return
}

// FromByte  二进制数据转换为 message
func FromByte(buf []byte) (msg *Message, err error) {
    msg = new(Message)
    if len(buf) < 8 {
        err = errors.New("数据异常不能转换")
        return
    }
    msg.ID, err = bytes.ToUint64(buf[0:8])
    //更新剩余数据
    buf = buf[8:]
    if err != nil {
        return
    }
    msg.Content = string(buf)
    return
}
