package list

import "testing"

func TestList(t *testing.T) {
    list := New()
    list.Rpush("a")
    list.Rpush("b")
    list.Rpush("c")
    list.Rpush("d")

    for ite := list.Iterator(); ite.HasNext(); {
        t.Log(ite.Next())
    }
}
