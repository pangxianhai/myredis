//对象数据结构
package object

import (
    "redis-server/data/dict"
    "redis-server/data/list"
    "redis-server/data/sds"
    "redis-server/data/zset"
)

type Type uint8

const (
    Sds  Type = 1
    List Type = 2
    Zset Type = 3
    Hash Type = 4
)

type Object struct {
    t Type        //类型
    o interface{} //数据类型
}

func New(v interface{}) *Object {
    obj := new(Object)
    switch v.(type) {
    case sds.Sds:
        obj.t = Sds
        obj.o = v
    case list.List:
        obj.t = List
        obj.o = v
    case zset.Zset:
        obj.t = Zset
        obj.o = v
    case dict.Dict:
        obj.t = Hash
        obj.o = v
    }
    return obj
}

func (obj *Object) SdsVal() *sds.Sds {
    v, ok := obj.o.(*sds.Sds)
    if ok {
        return v
    } else {
        return nil
    }
}

func (obj *Object) ListVal() *list.List {
    v, ok := obj.o.(*list.List)
    if ok {
        return v
    } else {
        return nil
    }
}

func (obj *Object) ZsetVal() *zset.Zset {
    v, ok := obj.o.(*zset.Zset)
    if ok {
        return v
    } else {
        return nil
    }
}

func (obj *Object) HasVal() *dict.Dict {
    v, ok := obj.o.(*dict.Dict)
    if ok {
        return v
    } else {
        return nil
    }
}
