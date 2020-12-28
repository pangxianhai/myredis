package bytes

import (
	"testing"
)

func TestToUint64(t *testing.T) {
	bs := []byte{0x0e, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}
	t.Log(bs)
	n := ToUint64(bs)
	t.Log(n)

}

func TestToUint32(t *testing.T) {
	b := FromUint32(^uint32(0))
	t.Log("结果：", b)
	v := ToUint32(b)
	t.Log("结果：", v)

}

func TestToUint16(t *testing.T) {
	b := FromUint16(^uint16(0))
	t.Log("结果：", b)
	v := ToUint16(b)
	t.Log("结果：", v)
}

func TestToUint8(t *testing.T) {
	b := FromUint8(^uint8(0))
	t.Log("结果：", b)
	v := ToUint8(b)
	t.Log("结果：", v)
}

func TestToInt64(t *testing.T) {
	b := FromInt64(^int64(0))
	t.Log("结果：", b)
	v := ToInt64(b)
	t.Log("结果：", v)
}

func TestToInt32(t *testing.T) {
	b := FromInt32(^int32(0))
	t.Log("结果：", b)
	v := ToInt32(b)
	t.Log("结果：", v)
}

func TestToInt16(t *testing.T) {
	b := FromInt16(^int16(0))
	t.Log("结果：", b)
	v := ToInt16(b)
	t.Log("结果：", v)
}

func TestToInt8(t *testing.T) {
	b := FromInt8(^int8(0))
	t.Log("结果：", b)
	v := ToInt8(b)
	t.Log("结果：", v)
}
