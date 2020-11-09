package util

import "testing"

func TestByteToUint32(t *testing.T) {
    b, _ := Uint32ToBytes(^uint32(0))
    t.Log("结果：", b)
    v, _ := ByteToUint32(b)
    t.Log("结果：", v)

}

func TestByteToUint16(t *testing.T) {
    b, _ := Uint16ToBytes(^uint16(0))
    t.Log("结果：", b)
    v, _ := ByteToUint16(b)
    t.Log("结果：", v)
}


func TestByteToUint8(t *testing.T) {
    b, _ := Uint8ToBytes(^uint8(0))
    t.Log("结果：", b)
    v, _ := ByteToUint8(b)
    t.Log("结果：", v)
}

func TestByteToInt64(t *testing.T) {
    b, _ := Int64ToBytes(^int64(0))
    t.Log("结果：", b)
    v, _ := ByteToInt64(b)
    t.Log("结果：", v)
}

func TestByteToInt32(t *testing.T) {
    b, _ := Int32ToBytes(^int32(0))
    t.Log("结果：", b)
    v, _ := ByteToInt32(b)
    t.Log("结果：", v)
}

func TestByteToInt16(t *testing.T) {
    b, _ := Int16ToBytes(^int16(0))
    t.Log("结果：", b)
    v, _ := ByteToInt16(b)
    t.Log("结果：", v)
}

func TestByteToInt8(t *testing.T) {
    b, _ := Int8ToBytes(^int8(0))
    t.Log("结果：", b)
    v, _ := ByteToInt8(b)
    t.Log("结果：", v)
}
