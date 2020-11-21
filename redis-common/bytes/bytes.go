package bytes

import (
    "bytes"
    "encoding/binary"
    "errors"
)

// ToUint64 64位byte 转成uint64
func ToUint64(byte []byte) (uint64, error) {
    if len(byte) != 8 {
        return 0, errors.New("转 uint64 byte 长度必须是8")
    }
    buffer := bytes.NewBuffer(byte)
    var value uint64
    err := binary.Read(buffer, binary.BigEndian, &value)
    if err != nil {
        return 0, err
    }
    return value, nil
}

// FromUint64 uint64 转 二进制
func FromUint64(n uint64) ([]byte, error) {
    buffer := bytes.NewBuffer([]byte{})
    err := binary.Write(buffer, binary.BigEndian, n)
    if err != nil {
        return nil, err
    }
    return buffer.Bytes(), nil
}

// ToUint32 32位byte 转成uint32
func ToUint32(byte []byte) (uint32, error) {
    if len(byte) != 4 {
        return 0, errors.New("转 uint32 byte 长度必须是4")
    }
    buffer := bytes.NewBuffer(byte)
    var value uint32
    err := binary.Read(buffer, binary.BigEndian, &value)
    if err != nil {
        return 0, err
    }
    return value, nil
}

// FromUint32 uint32 转 二进制
func FromUint32(n uint32) ([]byte, error) {
    buffer := bytes.NewBuffer([]byte{})
    err := binary.Write(buffer, binary.BigEndian, n)
    if err != nil {
        return nil, err
    }
    return buffer.Bytes(), nil
}

// ToUint16 16位byte 转成uint16
func ToUint16(byte []byte) (uint16, error) {
    if len(byte) != 2 {
        return 0, errors.New("转 uint16 byte 长度必须是2")
    }
    buffer := bytes.NewBuffer(byte)
    var value uint16
    err := binary.Read(buffer, binary.BigEndian, &value)
    if err != nil {
        return 0, err
    }
    return value, nil
}

// FromUint16 uint16 转 二进制
func FromUint16(n uint16) ([]byte, error) {
    buffer := bytes.NewBuffer([]byte{})
    err := binary.Write(buffer, binary.BigEndian, n)
    if err != nil {
        return nil, err
    }
    return buffer.Bytes(), nil
}

// ToUint8 8位byte 转成uint8
func ToUint8(byte []byte) (uint8, error) {
    if len(byte) != 1 {
        return 0, errors.New("转 uint32 byte 长度必须是1")
    }
    buffer := bytes.NewBuffer(byte)
    var value uint8
    err := binary.Read(buffer, binary.BigEndian, &value)
    if err != nil {
        return 0, err
    }
    return value, nil
}

// FromUint8 uint8 转 二进制
func FromUint8(n uint8) ([]byte, error) {
    buffer := bytes.NewBuffer([]byte{})
    err := binary.Write(buffer, binary.BigEndian, n)
    if err != nil {
        return nil, err
    }
    return buffer.Bytes(), nil
}

// ToInt64 64位byte 转成int64
func ToInt64(byte []byte) (int64, error) {
    if len(byte) != 8 {
        return 0, errors.New("转 int64 byte 长度必须是8")
    }
    buffer := bytes.NewBuffer(byte)
    var value int64
    err := binary.Read(buffer, binary.BigEndian, &value)
    if err != nil {
        return 0, err
    }
    return value, nil
}

// FromInt64 int64 转 二进制
func FromInt64(n int64) ([]byte, error) {
    buffer := bytes.NewBuffer([]byte{})
    err := binary.Write(buffer, binary.BigEndian, n)
    if err != nil {
        return nil, err
    }
    return buffer.Bytes(), nil
}

// ToInt32 64位byte 转成int64
func ToInt32(byte []byte) (int32, error) {
    if len(byte) != 4 {
        return 0, errors.New("转 int32 byte 长度必须是4")
    }
    buffer := bytes.NewBuffer(byte)
    var value int32
    err := binary.Read(buffer, binary.BigEndian, &value)
    if err != nil {
        return 0, err
    }
    return value, nil
}

// FromInt32 int32 转 二进制
func FromInt32(n int32) ([]byte, error) {
    buffer := bytes.NewBuffer([]byte{})
    err := binary.Write(buffer, binary.BigEndian, n)
    if err != nil {
        return nil, err
    }
    return buffer.Bytes(), nil
}

// ToInt16 16位byte 转成int16
func ToInt16(byte []byte) (int16, error) {
    if len(byte) != 2 {
        return 0, errors.New("转 uint16 byte 长度必须是2")
    }
    buffer := bytes.NewBuffer(byte)
    var value int16
    err := binary.Read(buffer, binary.BigEndian, &value)
    if err != nil {
        return 0, err
    }
    return value, nil
}

// FromInt16 int16 转 二进制
func FromInt16(n int16) ([]byte, error) {
    buffer := bytes.NewBuffer([]byte{})
    err := binary.Write(buffer, binary.BigEndian, n)
    if err != nil {
        return nil, err
    }
    return buffer.Bytes(), nil
}

// ToInt8 16位byte 转成int16
func ToInt8(byte []byte) (int8, error) {
    if len(byte) != 1 {
        return 0, errors.New("转 int8 byte 长度必须是1")
    }
    buffer := bytes.NewBuffer(byte)
    var value int8
    err := binary.Read(buffer, binary.BigEndian, &value)
    if err != nil {
        return 0, err
    }
    return value, nil
}

// FromInt8 int8 转 二进制
func FromInt8(n int8) ([]byte, error) {
    buffer := bytes.NewBuffer([]byte{})
    err := binary.Write(buffer, binary.BigEndian, n)
    if err != nil {
        return nil, err
    }
    return buffer.Bytes(), nil
}
