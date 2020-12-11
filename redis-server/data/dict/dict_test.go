package dict

import (
    _ "database/sql"
    "fmt"
    "log"
    "testing"
)

func TestDict(t *testing.T) {
    dict := New()
    dict.Put("a", "b")
    dict.Put("aa", "c")
    dict.Put("aa", "d")
    //t.Log(dict.Get("a"))
    //
    ////for k, v := range dict {
    ////    t.Log(k, "=>", v)
    ////}
    //
    //println(dict)
    //
    //t.Log(dict)

    for ite := dict.Iterator(); ite.HasNext(); {
        entry := ite.Next()
        e := entry.(*Entry)
        t.Log(e.Key(), "=>", e.Value())
    }
}

func Test2(t *testing.T) {
    ss("a", "b")
}

func ss(v ...interface{}) {
    //a := make([]interface{}, 1)
    //a[0] = "[INFO]"
    //v1 := append(a, v...)
    s := fmt.Sprintln(v...)
    log.Println(s)
}
