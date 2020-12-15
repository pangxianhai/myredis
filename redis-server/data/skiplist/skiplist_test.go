package skiplist

import (
    "fmt"
    "math"
    "redis-server/data/comparable"
    "testing"
)

func TestSkipListAdd(t *testing.T) {
    l := New()
    l.Add(&A{data: "a"}, 1)
    l.Add(&A{data: "f"}, 6)
    l.Add(&A{data: "c"}, 3)
    l.Add(&A{data: "b"}, 2)
    l.Add(&A{data: "h"}, 8)
    l.Add(&A{data: "d"}, 4)
    l.Add(&A{data: "g"}, 7)
    l.Add(&A{data: "e"}, 5)
    l.Add(&A{data: "i"}, 9)

    //t.Log("heald:", l.head)
    //
    //for ite := l.Iterator(); ite.HasNext(); {
    //    t.Log(ite.Next())
    //}
    t.Log(l)
}

func TestSkipListRemove(t *testing.T) {
    l := New()
    l.Add(&A{data: "a"}, 1)
    l.Add(&A{data: "f"}, 6)
    l.Add(&A{data: "c"}, 3)
    l.Add(&A{data: "b"}, 2)
    l.Add(&A{data: "h"}, 8)
    l.Add(&A{data: "d"}, 4)
    l.Add(&A{data: "g"}, 7)
    l.Add(&A{data: "e"}, 5)
    l.Add(&A{data: "i"}, 9)

    for ite := l.Iterator(); ite.HasNext(); {
        t.Log(ite.Next())
    }

    l.Remove(&A{data: "d"}, 4)

    t.Log("删除d后....")
    t.Log("层数", l.level, "heald:", l.head)
    for ite := l.Iterator(); ite.HasNext(); {
        t.Log(ite.Next())
    }

    l.Remove(&A{data: "e"}, 5)

    t.Log("删除e后....")
    t.Log("层数", l.level, "heald:", l.head)
    for ite := l.Iterator(); ite.HasNext(); {
        t.Log(ite.Next())
    }

    l.Remove(&A{data: "i"}, 9)

    t.Log("删除i后....")
    t.Log("层数", l.level, "heald:", l.head)
    for ite := l.Iterator(); ite.HasNext(); {
        t.Log(ite.Next())
    }

    l.Remove(&A{data: "a"}, 1)

    t.Log("删除a后....")
    t.Log("层数", l.level, "heald:", l.head)
    for ite := l.Iterator(); ite.HasNext(); {
        t.Log(ite.Next())
    }

    l.Remove(&A{data: "f"}, 6)
    t.Log("删除f后....")
    t.Log("层数", l.level, "heald:", l.head)
    for ite := l.Iterator(); ite.HasNext(); {
        t.Log(ite.Next())
    }

    l.Remove(&A{data: "b"}, 2)
    t.Log("删除b后....")
    t.Log("层数", l.level, "heald:", l.head)
    for ite := l.Iterator(); ite.HasNext(); {
        t.Log(ite.Next())
    }

    l.Remove(&A{data: "g"}, 7)
    t.Log("删除g后....")
    t.Log("层数", l.level, "heald:", l.head)
    for ite := l.Iterator(); ite.HasNext(); {
        t.Log(ite.Next())
    }

    l.Remove(&A{data: "c"}, 3)
    t.Log("删除c后....")
    t.Log("层数", l.level, "heald:", l.head)
    for ite := l.Iterator(); ite.HasNext(); {
        t.Log(ite.Next())
    }

    l.Remove(&A{data: "h"}, 8)
    t.Log("删除h后....")
    t.Log("层数", l.level, "heald:", l.head)
    for ite := l.Iterator(); ite.HasNext(); {
        t.Log(ite.Next())
    }
}

func TestSkipListRange(t *testing.T) {
    l := New()
    l.Add(&A{data: "a"}, 1)
    l.Add(&A{data: "f"}, 6)
    l.Add(&A{data: "c"}, 3)
    l.Add(&A{data: "b"}, 2)
    l.Add(&A{data: "h"}, 8)
    l.Add(&A{data: "d"}, 4)
    l.Add(&A{data: "g"}, 7)
    l.Add(&A{data: "e"}, 5)
    l.Add(&A{data: "i"}, 9)

    ll := l.Range(3, 5)

    t.Log("ll len", ll.Len())
    if ll.Len() == 0 {
        t.Log(l.head)
        for ite := l.Iterator(); ite.HasNext(); {
            t.Log(ite.Next())
        }
        //ll = l.Range(3, 5)
    }

    for ite := ll.Iterator(); ite.HasNext(); {
        t.Log(ite.Next())
    }
}

type A struct {
    data string
}

func (a *A) CompareTo(c interface{}) int {
    if a.data == c.(*A).data {
        return 0
    }
    if a.data < c.(*A).data {
        return -1
    }
    return 1
}

func (a *A) String() string {
    return a.data
}

type B struct {
    a comparable.Comparable
}

func TestA(t *testing.T) {
    a1 := A{data: "a"}
    a2 := A{data: "c"}

    b1 := B{a: &a1}
    b2 := B{a: &a2}

    //c := a1.CompareTo(a2)

    c := b1.a.CompareTo(b2)

    t.Log(c)
}

func b() {
    panic("error hhh")
}

func TestB(t *testing.T) {
    defer func() {
        err := recover()
        if err != nil {
            t.Log("发生了错误", err)
        }
    }()
    fmt.Println("aaaa....")
    b()
    fmt.Println("bbbb....")
}

func TestC(t *testing.T) {
    a := int(math.Log2(float64(7)))
    t.Log(a)
}
