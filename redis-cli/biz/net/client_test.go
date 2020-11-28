package net

import (
    "testing"
    "time"
)

func TestSendMessage(t *testing.T) {
    a := [10000][10000]int{}
    s := time.Now().UnixNano() / 1000000
    for i, _ := range a {
        for j, _ := range a[i] {
            a[i][j] = j
        }
    }
    t.Log(time.Now().UnixNano()/1000000 - s)
    s = time.Now().UnixNano() / 1000000
    for i := 0; i < len(a); i++ {
        for j := 0; j < len(a[i]); j++ {
            a[j][i] = j
        }
    }
    t.Log(time.Now().UnixNano()/1000000 - s)
    t.Log(time.Now().UnixNano())
}
