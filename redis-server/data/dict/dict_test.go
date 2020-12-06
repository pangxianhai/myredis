package dict

import "testing"

func TestDict(t *testing.T) {
    dict := New()
    dict.Put("a", "b")
    dict.Put("aa", "c")
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
        t.Log(e.key, "=>", e.value)
    }

}
