package packageinfo

import "testing"

func TestPackage(t *testing.T) {
    p := PackageInfo{
        MsgID: 12,
        Len:   1,
        Buf:   []byte{0x2},
    }
    b, err := PackageToByte(&p)
    p2 := PackageInfo{
        MsgID: 19,
        Len:   2,
        Buf:   []byte{0x29, 0x32},
    }
    b2, err := PackageToByte(&p2)

    b = append(b, b2...)
    t.Log(len(b))

    pp, err := ByteToPackage(nil, b[0:5])
    t.Log("error=", err)
    t.Log("data=", pp)

    ppp, err := ByteToPackage(&pp[0], b[5:])
    t.Log("error1=", err)
    t.Log("data1=", ppp)
    t.Log("data=", pp)
}
