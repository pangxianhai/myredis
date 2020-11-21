package packet

import (
    "testing"
)

func TestPackage(t *testing.T) {
    p := Packet{
        Len: 1,
        Buf: []byte{0x2},
    }
    b, err := ToByte(&p)
    p2 := Packet{
        Len: 2,
        Buf: []byte{0x29, 0x32},
    }
    b2, err := ToByte(&p2)

    b = append(b, b2...)
    t.Log(len(b))

    pp, err := FromByte(nil, b[0:5])
    t.Log("error=", err)
    t.Log("data=", pp)

    ppp, err := FromByte(&pp[0], b[5:])
    t.Log("error1=", err)
    t.Log("data1=", ppp)
    t.Log("data=", pp)
}

func TestA(t *testing.T) {
    var p []Packet
    p = append(p, Packet{Len: 12})
    t.Log(p[0].Len)
    a(p)
    t.Log(p[0].Len)
}

func a(l []Packet) {
    l[0].Len = 45
}
