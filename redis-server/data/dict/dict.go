package dict

import (
    "bytes"
    "encoding/gob"
    "log"
    "redis-server/common/iterator"
    "redis-server/common/murmur3"
    "redis-server/common/numbers"
)

type Entry struct {
    key   interface{}
    value interface{}
    next  *Entry
}

type Ht struct {
    table []*Entry
    mask  uint32
    size  uint32
}

type Dict struct {
    ht       [2]*Ht
    size     uint32
    isRehash bool // 当 isRehash == true 当时正在rehash
}

type Iterator struct {
    dict    *Dict  //字典
    hti     int    //ht 索引
    index   int    // table 索引
    current *Entry // 当前遍历实体
    next    *Entry // 下一个遍历实体
}

func New() *Dict {
    ht := new(Ht)
    ht.table = make([]*Entry, 16)
    ht.mask = 15
    ht.size = 0

    dict := new(Dict)
    dict.size = 0
    dict.isRehash = false
    dict.ht[0] = ht

    return dict
}

func (dict *Dict) Put(key interface{}, value interface{}) {
    //先进行 rehash ,是否需要 rehash 在 rehash里判断
    dict.rehash()
    if dict.isRehash {
        // rehash 还未完成
        dict.ht[1].put(key, value)
    } else {
        // 已经 rehash 完成
        dict.ht[0].put(key, value)
    }
}

func (dict *Dict) Get(key interface{}) interface{} {
    //先进行 rehash ,是否需要 rehash 在 rehash里判断
    dict.rehash()
    var v interface{}
    if dict.ht[0] != nil {
        v = dict.ht[0].get(key)
        if v != nil {
            return v
        }
    } else if dict.ht[1] != nil {
        v = dict.ht[1].get(key)
        if v != nil {
            return v
        }
    }
    return nil
}

func (dict *Dict) Remove(key interface{}) {
    //先进行 rehash ,是否需要 rehash 在 rehash里判断
    dict.rehash()
    if dict.ht[0] != nil {
        dict.ht[0].remove(key)
    } else if dict.ht[1] != nil {
        dict.ht[1].remove(key)
    }
}

func (dict *Dict) Iterator() iterator.Iterator {
    ite := new(Iterator)
    ite.dict = dict
    ite.hti, ite.index, ite.next = dict.findNext(0, 0)
    return ite
}

func (dict *Dict) findNext(hti, index int) (int, int, *Entry) {
    htL := len(dict.ht)
    for hi := hti; hi < htL; hi++ {
        if dict.ht[hi] == nil {
            continue
        }
        table := dict.ht[hi].table
        tL := len(table)
        for i := index + 1; i < tL; i++ {
            if table[i] != nil {
                return hi, i, table[i]
            }
        }
    }
    return 0, 0, nil
}

func (dict *Dict) rehash() {
    need := dict.needRehash()
    dict.isRehash = need
    if !need {
        return
    }
    dict.onRehash()
}

func (dict *Dict) onRehash() {
    if dict.ht[0] == nil || dict.ht[1] == nil {
        //无数据 不rehash
        return
    }
    maxRehashCount := 1024
    count := 0
    for i := 0; i < len(dict.ht[0].table); i++ {
        entry := dict.ht[0].table[i]
        for entry != nil {
            dict.ht[1].put(entry.key, entry.value)
            entry = entry.next
            count++
        }
        dict.ht[0].table[i] = nil
        if i == len(dict.ht[0].table)-1 {
            //已经同步完成
            dict.ht[0] = dict.ht[1]
            dict.ht[1] = nil
            dict.isRehash = false
            break
        }
        if count >= maxRehashCount {
            break
        }
    }
}

func (dict *Dict) needRehash() bool {
    if dict.isRehash {
        return true
    } else {
        //当前没有进行 rehash
        loadFactor := float32(dict.ht[0].size) / float32(dict.ht[0].mask+1)
        if loadFactor >= 1 {
            //需要扩大 扩大两倍
            h1Size := numbers.Max2N32((dict.ht[0].mask + 1) * 2)
            ht1 := new(Ht)
            ht1.table = make([]*Entry, h1Size)
            ht1.mask = h1Size - 1
            ht1.size = 0
            dict.ht[1] = ht1
            return true
        } else if loadFactor <= 0.1 && dict.ht[0].size > 16 {
            //收缩
            h1Size := numbers.Max2N32((dict.ht[0].mask + 1) / 2)
            ht1 := new(Ht)
            ht1.table = make([]*Entry, h1Size)
            ht1.mask = h1Size - 1
            ht1.size = 0
            dict.ht[1] = ht1
            return true
        } else {
            return false
        }
    }
}

func (ht *Ht) put(key interface{}, value interface{}) {
    index := ht.hash(key)
    entry := new(Entry)
    entry.key = key
    entry.value = value

    v := ht.table[index]
    if v == nil {
        ht.table[index] = entry
    } else {
        p := ht.table[index]
        //找到相同key 只替换value
        for p != nil {
            if p.key == key {
                p.value = value
                return
            }
        }
        //无相同key  加到头部
        entry.next = v
        ht.table[index] = entry
    }
    return
}

func (ht *Ht) get(key interface{}) interface{} {
    index := ht.hash(key)

    entry := ht.table[index]
    if entry == nil {
        return nil
    }
    p := entry
    for p != nil {
        if p.key == key {
            return p.value
        }
        p = p.next
    }
    return nil
}

func (ht *Ht) remove(key interface{}) {
    index := ht.hash(key)

    entry := ht.table[index]
    if entry == nil {
        return
    }
    if entry.key == key {
        ht.table[index] = entry.next
        entry.next = nil
    } else {
        p := entry.next
        pre := entry
        for p != nil {
            if p.key == key {
                pre.next = p.next
                p.next = nil
                break
            }
            p = p.next
            pre = pre.next
        }
    }
}

func (ht *Ht) hash(key interface{}) uint32 {
    var buf bytes.Buffer
    enc := gob.NewEncoder(&buf)
    err := enc.Encode(key)
    if err != nil {
        log.Println("hash key failed", key, err)
        return 0
    }
    b := buf.Bytes()
    hashV := murmur3.Sum32(b)
    return hashV & ht.mask
}

func (ite *Iterator) HasNext() bool {
    return ite.next != nil
}

func (ite *Iterator) Next() interface{} {
    ite.current = ite.next
    if ite.current != nil && ite.current.next != nil {
        ite.next = ite.current.next
    } else {
        ite.hti, ite.index, ite.next = ite.dict.findNext(ite.hti, ite.index)
    }
    return ite.current
}

func (e *Entry) Key() interface{} {
    return e.key
}

func (e *Entry) Value() interface{} {
    return e.value
}
