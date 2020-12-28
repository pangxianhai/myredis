package zset

import (
    "redis-server/data/dict"
    "redis-server/data/list"
    "redis-server/data/sds"
    "redis-server/data/skiplist"
)

type Zset struct {
    skipList *skiplist.SkipList
    dict     *dict.Dict
}

func New() *Zset {
    zset := new(Zset)
    zset.skipList = skiplist.New()
    zset.dict = dict.New()
    return zset
}

func (zset *Zset) Add(v *sds.Sds, score float64) {
    s := zset.dict.Get(v)
    if s != nil {
        if s.(float64) == score {
            //数据没变化无需操作
            return
        }
        //如果元素已经存在先删除在保存
        zset.skipList.Remove(v, s.(float64))
    }
    zset.dict.Put(v, score)
    zset.skipList.Add(v, score)
}

func (zset *Zset) Remove(v *sds.Sds) {
    s := zset.dict.Get(v)
    if s == nil {
        return
    }
    zset.skipList.Remove(v, s.(float64))
    zset.dict.Remove(v)
}

func (zset *Zset) Range(start, end int) *list.List {
    return zset.skipList.Range(start, end)
}

func (zset *Zset) RevRange(start, end int) *list.List {
    return zset.skipList.RevRange(start, end)
}

func (zset *Zset) RangeByScore(min, max float64, offset, count int) *list.List {
    return zset.skipList.RangeByScore(min, max, offset, count)
}

func (zset *Zset) RevRangeByScore(min, max float64, offset, count int) *list.List {
    return zset.skipList.RevRangeByScore(min, max, offset, count)
}
