package list

import "redis-server/common/iterator"

type Node struct {
    value interface{}
    next  *Node
    prev  *Node
}

type List struct {
    head *Node
    tail *Node
    len  int
}

type Iterator struct {
    list *List
    cur  *Node
    next *Node
}

func New() *List {
    list := new(List)
    return list
}

func (list *List) Len() int {
    return list.len
}

func (list *List) Rpush(value interface{}) {
    node := new(Node)
    node.value = value
    prev := list.tail
    if prev != nil {
        prev.next = node
        node.prev = prev
    }
    if list.head == nil {
        list.head = node
    }
    list.tail = node
    list.len++
}

func (list *List) Lpush(value interface{}) {
    node := new(Node)
    node.value = value
    next := list.head
    if next != nil {
        node.next = next
        next.prev = node
    }
    list.head = node
    if list.tail == nil {
        list.tail = node
    }
    list.len++
}

func (list *List) Remove(value interface{}) {
    if list.len == 0 {
        return
    }
    for p := list.head; p != nil; p = p.next {
        if p.value == value {
            if p.next == nil {
                p.prev.next = nil
                list.tail = p.prev
                p.prev = nil
                list.len--
                break
            } else {
                p.prev = p.next
                p.next.prev = p.prev
                list.len--
            }
        }
    }
}

func (list *List) Lpop() (value interface{}) {
    if list.head == nil {
        return nil
    }
    value = list.head.value
    if list.head.next != nil {
        list.head.prev = nil
    }
    list.head = list.head.next
    list.len--
    return
}

func (list *List) Rpop() (value interface{}) {
    if list.tail == nil {
        return nil
    }
    value = list.tail.value
    if list.tail.prev != nil {
        list.tail.next = nil
    }
    list.tail = list.tail.prev
    list.len--
    return
}

func (list *List) Get(index int) interface{} {
    if index < 0 || list.len <= index {
        return nil
    }
    i := 0
    for p := list.head; p != nil; p = p.next {
        if i == index {
            return p.value
        } else {
            i++
        }
    }
    return nil
}

func (list *List) Iterator() iterator.Iterator {
    ite := new(Iterator)
    ite.list = list
    ite.next = list.head
    return ite
}

func (ite *Iterator) HasNext() bool {
    return ite.next != nil
}

func (ite *Iterator) Next() interface{} {
    ite.cur = ite.next
    ite.next = ite.cur.next
    return ite.cur.value
}
