package message

import (
    "encoding/binary"
    "testing"
)

func TestNew(t *testing.T) {
    bs := []byte{0x1f, 0xff}

    t.Log(binary.BigEndian.Uint16(bs))
    t.Log(int16(binary.BigEndian.Uint16(bs)))

    n := toInt16(bs)
    t.Log(n)

    ab := fromInt16(n)
    t.Log(ab)

}

func toInt16(bs []byte) (n int16) {
    for _, b := range bs {
        n = n<<8 | int16(b)
    }
    return
}

func fromInt16(n int16) []byte {
    b := make([]byte, 2)

    b[1] = byte(n)
    b[0] = byte(n >> 8)
    return b
}
