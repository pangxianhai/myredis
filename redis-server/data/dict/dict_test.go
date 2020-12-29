package dict

import (
    _ "database/sql"
    "fmt"
    "redis-server/data/sds"
    "reflect"
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

func TestKey(t *testing.T) {
    s1 := sds.NewWithStr("a")
    s2 := sds.NewWithStr("a")
    dict := New()
    dict.Put(s1, "aaa")

    t.Log(dict.Get(s2))

    a(s1)
}

func a(i interface{}) {
    v := reflect.ValueOf(i)
    v.Elem()
    fmt.Println(v.Type())
}
