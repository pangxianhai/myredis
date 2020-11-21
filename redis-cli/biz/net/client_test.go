package net

import "testing"

func TestSendMessage(t *testing.T) {
    b := []byte{0x7f}

    var x uint64

    for _, b := range b[0:] {
        x = x<<8 | uint64(b)
    }
    t.Log(x)
    t.Log(uint64(b[0]))
}
