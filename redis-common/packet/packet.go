package packet

import (
    "redis-common/bytes"
    "redis-common/message"
)

const MaxLen = 10240

// Packet 定义数据包格式
type Packet struct {
    Len    uint32
    Accept uint32
    Buf    []byte
}

func New(msg *message.Message) *Packet {
    packet := Packet{}
    packet.Buf, _ = message.ToByte(msg)
    packet.Len = uint32(len(packet.Buf))
    return &packet
}

// FromByte 将二进制数据转换为 PackageInfo 对象
// buf 格式为 4位MsgID + 4位len + 数据
// 如果发生包解析错误 直接返回 error 服务端收到该错误可以主动与客户端断开连接 客户端重连服务器
func FromByte(lastPacket *Packet, buf []byte) ([]Packet, error) {
    bufL := uint32(len(buf))
    if bufL == 0 {
        return nil, nil
    }
    if lastPacket != nil && lastPacket.Accept >= lastPacket.Len {
        lastPacket = nil
    }
    var newBuf []byte
    if lastPacket != nil && lastPacket.Buf != nil {
        newBuf = append(lastPacket.Buf, buf...)
        //清空元数据
        lastPacket.Buf = []byte{}
    } else {
        newBuf = buf
    }
    var newPacketList []Packet
    for len(newBuf) > 0 {
        var err error
        newBuf, err = parseLastPacket(lastPacket, newBuf)
        if err != nil {
            return nil, err
        }
        var newPacket *Packet
        newPacket, newBuf, err = parsePacket(newBuf)
        if err != nil {
            return nil, err
        }
        if newPacket != nil {
            newPacketList = append(newPacketList, *newPacket)
        }
    }
    return newPacketList, nil
}

// ToByte 将 PackageInfo 对象转换二进制
func ToByte(packageInfo *Packet) ([]byte, error) {
    buf := make([]byte, 0)
    b, err := bytes.FromUint32(packageInfo.Len)
    if err != nil {
        return nil, err
    }
    buf = append(buf, b...)
    buf = append(buf, packageInfo.Buf...)
    return buf, nil
}

func parseLastPacket(lastPacket *Packet, buf []byte) ([]byte, error) {
    if lastPacket == nil {
        //数据没有发生变化
        return buf, nil
    }
    if lastPacket.Len == 0 {
        if len(buf) < 4 {
            //数据不够 数据存于 lastPackage.Buf中
            lastPacket.Buf = buf
            return nil, nil
        }
        l, err := bytes.ToUint32(buf[0:4])
        if err != nil {
            return nil, err
        }
        lastPacket.Len = l
        //更新剩余数据
        buf = buf[4:]
    }
    bufL := uint32(len(buf))
    if bufL == 0 {
        return nil, nil
    }
    needLen := lastPacket.Len - lastPacket.Accept
    if needLen == 0 {
        return buf, nil
    }
    if bufL <= needLen {
        lastPacket.Buf = append(lastPacket.Buf, buf...)
        lastPacket.Accept += bufL
        //buf中数据已经用完
        return nil, nil
    } else {
        lastPacket.Buf = append(lastPacket.Buf, buf[0:needLen]...)
        lastPacket.Accept += needLen
        buf = buf[needLen:]
        return buf, nil
    }
}

func parsePacket(buf []byte) (*Packet, []byte, error) {
    if len(buf) <= 0 {
        return nil, nil, nil
    }
    p := Packet{}
    var err error
    buf, err = parseLastPacket(&p, buf)
    if err != nil {
        return nil, nil, err
    } else {
        return &p, buf, nil
    }
}
