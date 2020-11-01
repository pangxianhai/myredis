package list

import "fmt"

type List struct {
}

func New() List {
    fmt.Println("新建链表")
    return List{}
}
