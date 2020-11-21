package bytes

import "testing"

func TestToUint32(t *testing.T) {
    b, _ := FromUint32(^uint32(0))
    t.Log("结果：", b)
    v, _ := ToUint32(b)
    t.Log("结果：", v)

}

func TestToUint16(t *testing.T) {
    b, _ := FromUint16(^uint16(0))
    t.Log("结果：", b)
    v, _ := ToUint16(b)
    t.Log("结果：", v)
}

func TestToUint8(t *testing.T) {
    b, _ := FromUint8(^uint8(0))
    t.Log("结果：", b)
    v, _ := ToUint8(b)
    t.Log("结果：", v)
}

func TestToInt64(t *testing.T) {
    b, _ := FromInt64(^int64(0))
    t.Log("结果：", b)
    v, _ := ToInt64(b)
    t.Log("结果：", v)
}

func TestToInt32(t *testing.T) {
    b, _ := FromInt32(^int32(0))
    t.Log("结果：", b)
    v, _ := ToInt32(b)
    t.Log("结果：", v)
}

func TestToInt16(t *testing.T) {
    b, _ := FromInt16(^int16(0))
    t.Log("结果：", b)
    v, _ := ToInt16(b)
    t.Log("结果：", v)
}

func TestToInt8(t *testing.T) {
    b, _ := FromInt8(^int8(0))
    t.Log("结果：", b)
    v, _ := ToInt8(b)
    t.Log("结果：", v)
}
