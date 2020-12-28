package bytes

import "github.com/gogo/protobuf/proto"

// ToUint64 64位byte 转成uint64
func ToUint64(byte []byte) uint64 {
    var n uint64
    for i, b := range byte {
        n = n | uint64(b)<<(i*8)
    }
    return n
}

// FromUint64 uint64 转 二进制
func FromUint64(n uint64) []byte {
    b := make([]byte, 8)
    b[7] = byte(n >> 56)
    b[6] = byte(n >> 48)
    b[5] = byte(n >> 40)
    b[4] = byte(n >> 32)
    b[3] = byte(n >> 24)
    b[2] = byte(n >> 16)
    b[1] = byte(n >> 8)
    b[0] = byte(n)
    return b
}

// ToUint32 32位byte 转成uint32
func ToUint32(byte []byte) uint32 {
    var n uint32
    for i, b := range byte {
        n = n | uint32(b)<<(i*8)
    }
    return n
}

// FromUint32 uint32 转 二进制
func FromUint32(n uint32) []byte {
    b := make([]byte, 4)
    b[3] = byte(n >> 24)
    b[2] = byte(n >> 16)
    b[1] = byte(n >> 8)
    b[0] = byte(n)
    return b
}

// ToUint16 16位byte 转成uint16
func ToUint16(byte []byte) uint16 {
    var n uint16
    for i, b := range byte {
        n = n | uint16(b)<<(i*8)
    }
    return n
}

// FromUint16 uint16 转 二进制
func FromUint16(n uint16) []byte {
    b := make([]byte, 2)
    b[1] = byte(n >> 8)
    b[0] = byte(n)
    return b
}

// ToUint8 8位byte 转成uint8
func ToUint8(byte []byte) uint8 {
    return byte[0]
}

// FromUint8 uint8 转 二进制
func FromUint8(n uint8) []byte {
    b := make([]byte, 1)
    b[0] = n
    return b
}

// ToInt64 64位byte 转成int64
func ToInt64(byte []byte) int64 {
    var n int64
    for i, b := range byte {
        n = n | int64(b)<<(i*8)
    }
    return n
}

// FromInt64 int64 转 二进制
func FromInt64(n int64) []byte {
    b := make([]byte, 8)
    b[7] = byte(n >> 56)
    b[6] = byte(n >> 48)
    b[5] = byte(n >> 40)
    b[4] = byte(n >> 32)
    b[3] = byte(n >> 24)
    b[2] = byte(n >> 16)
    b[1] = byte(n >> 8)
    b[0] = byte(n)
    return b
}

// ToInt32 64位byte 转成int32
func ToInt32(byte []byte) int32 {
    var n int32
    for i, b := range byte {
        n = n | int32(b)<<(i*8)
    }
    return n
}

// FromInt32 int32 转 二进制
func FromInt32(n int32) []byte {
    b := make([]byte, 4)
    b[3] = byte(n >> 24)
    b[2] = byte(n >> 16)
    b[1] = byte(n >> 8)
    b[0] = byte(n)
    return b
}

// ToInt16 16位byte 转成int16
func ToInt16(byte []byte) int16 {
    var n int16
    for i, b := range byte {
        n = n | int16(b)<<(i*8)
    }
    return n
}

// FromInt16 int16 转 二进制
func FromInt16(n int16) []byte {
    b := make([]byte, 2)
    b[1] = byte(n >> 8)
    b[0] = byte(n)
    return b
}

// ToInt8 16位byte 转成int16
func ToInt8(byte []byte) int8 {
    return int8(byte[0])
}

// FromInt8 int8 转 二进制
func FromInt8(n int8) []byte {
    b := make([]byte, 1)
    b[0] = byte(n)
    return b
}

func FromPb(pb proto.Message) []byte {
    if pb == nil {
        return nil
    }
    b, err := proto.Marshal(pb)
    if err != nil {
        panic(err)
    }
    return b
}

func ToPb(b []byte, pb proto.Message) {
    if b == nil {
        return
    }
    err := proto.Unmarshal(b, pb)
    if err != nil {
        panic(err)
    }
}
