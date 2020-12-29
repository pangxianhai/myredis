package object

import (
    "redis-server/data/sds"
    "testing"
)

func TestNew(t *testing.T) {
    s := sds.NewWithStr("aaa")
    o := New(s)

    t.Log(o.SdsVal())
}
