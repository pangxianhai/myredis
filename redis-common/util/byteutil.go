package util

import (
    "bytes"
    "encoding/binary"
    "errors"
)

// ByteToUint64 64位byte 转成uint64
func ByteToUint64(byte []byte) (uint64, error) {
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

// Uint64ToBytes uint64 转 二进制
func Uint64ToBytes(n uint64) ([]byte, error) {
    buffer := bytes.NewBuffer([]byte{})
    err := binary.Write(buffer, binary.BigEndian, n)
    if err != nil {
        return nil, err
    }
    return buffer.Bytes(), nil
}

// ByteToUint32 32位byte 转成uint32
func ByteToUint32(byte []byte) (uint32, error) {
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

// Uint32ToBytes uint32 转 二进制
func Uint32ToBytes(n uint32) ([]byte, error) {
    buffer := bytes.NewBuffer([]byte{})
    err := binary.Write(buffer, binary.BigEndian, n)
    if err != nil {
        return nil, err
    }
    return buffer.Bytes(), nil
}

// ByteToUint16 16位byte 转成uint16
func ByteToUint16(byte []byte) (uint16, error) {
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

// Uint16ToBytes uint16 转 二进制
func Uint16ToBytes(n uint16) ([]byte, error) {
    buffer := bytes.NewBuffer([]byte{})
    err := binary.Write(buffer, binary.BigEndian, n)
    if err != nil {
        return nil, err
    }
    return buffer.Bytes(), nil
}

// ByteToUint8 8位byte 转成uint8
func ByteToUint8(byte []byte) (uint8, error) {
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

// Uint8ToBytes uint8 转 二进制
func Uint8ToBytes(n uint8) ([]byte, error) {
    buffer := bytes.NewBuffer([]byte{})
    err := binary.Write(buffer, binary.BigEndian, n)
    if err != nil {
        return nil, err
    }
    return buffer.Bytes(), nil
}

// ByteToInt64 64位byte 转成int64
func ByteToInt64(byte []byte) (int64, error) {
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

// Int64ToBytes int64 转 二进制
func Int64ToBytes(n int64) ([]byte, error) {
    buffer := bytes.NewBuffer([]byte{})
    err := binary.Write(buffer, binary.BigEndian, n)
    if err != nil {
        return nil, err
    }
    return buffer.Bytes(), nil
}

// ByteToInt32 64位byte 转成int64
func ByteToInt32(byte []byte) (int32, error) {
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

// Int32ToBytes int32 转 二进制
func Int32ToBytes(n int32) ([]byte, error) {
    buffer := bytes.NewBuffer([]byte{})
    err := binary.Write(buffer, binary.BigEndian, n)
    if err != nil {
        return nil, err
    }
    return buffer.Bytes(), nil
}

// ByteToInt16 16位byte 转成int16
func ByteToInt16(byte []byte) (int16, error) {
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

// Int16ToBytes int16 转 二进制
func Int16ToBytes(n int16) ([]byte, error) {
    buffer := bytes.NewBuffer([]byte{})
    err := binary.Write(buffer, binary.BigEndian, n)
    if err != nil {
        return nil, err
    }
    return buffer.Bytes(), nil
}

// ByteToInt8 16位byte 转成int16
func ByteToInt8(byte []byte) (int8, error) {
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

// Int8ToBytes int8 转 二进制
func Int8ToBytes(n int8) ([]byte, error) {
    buffer := bytes.NewBuffer([]byte{})
    err := binary.Write(buffer, binary.BigEndian, n)
    if err != nil {
        return nil, err
    }
    return buffer.Bytes(), nil
}
