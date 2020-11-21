package message

import "testing"

func TestNew(t *testing.T) {
    msg := New(123, "asdfsfdfs")
    b, _ := ToByte(msg)

    msg1, _ := FromByte(b)

    t.Log(msg)
    t.Log(msg1)
}
