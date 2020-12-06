package sds

import "testing"

func Test(t *testing.T) {
    s := "你好123"
    sds := NewWithStr(s)
    t.Log(sds.buf)
    t.Log(sds.len)
    t.Log(sds.free)
    t.Log(sds)

    sds.Cat("你好123你好123你好123你好123你好123")

    t.Log(sds.buf)
    t.Log(len(sds.buf))
    t.Log(sds.len)
    t.Log(sds.free)
    t.Log(sds)

    sds.Cat("我好")

    t.Log(sds.buf)
    t.Log(len(sds.buf))
    t.Log(sds.len)
    t.Log(sds.free)
    t.Log(sds)

}
