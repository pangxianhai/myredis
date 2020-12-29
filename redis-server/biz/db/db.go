package db

import (
    "redis-server/common/logger"
    "redis-server/data/dict"
    "redis-server/data/object"
    "redis-server/data/sds"
)

type Db struct {
    dict    *dict.Dict
    expires *dict.Dict
}

var db *Db

func Init() {
    db = new(Db)
    db.dict = dict.New()
    db.expires = dict.New()
    logger.Info("DB init finish !")
}

func Get(key *sds.Sds) *object.Object {
    v := db.dict.Get(key)
    if v == nil {
        return nil
    }
    return v.(*object.Object)
}

func Put(key *sds.Sds, obj *object.Object) {
    db.dict.Put(key, obj)
}

func SetExpire(key *sds.Sds, expireAt int64) int {
    v := db.dict.Get(key)
    if v == nil {
        return 0
    }
    db.expires.Put(key, expireAt)
    return 1
}
