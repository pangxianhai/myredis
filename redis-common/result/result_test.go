package result

import "testing"

func TestToJson(t *testing.T) {
    r := NewOfData("你好")
    s, _ := ToJson(r)
    t.Log(s)

    r1, _ := FromJson(s)

    t.Log(r1)
}
