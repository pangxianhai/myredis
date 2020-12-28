package db

import (
    "redis-server/common/logger"
    "redis-server/data/dict"
    "redis-server/data/sds"
)

type Db struct {
    dict    *dict.Dict
    expires *dict.Dict
}

var db *Db

func Init() {
    db := new(Db)
    db.dict = dict.New()
    db.expires = dict.New()
    logger.Info("DB init finish !")
}

func Get(key *sds.Sds) interface{} {
    return db.dict.Get(key)
}

func Put(key *sds.Sds, value interface{}) {
    db.dict.Put(key, value)
}

func SetExpire(key *sds.Sds, expireAt int64) int {
    v := db.dict.Get(key)
    if v == nil {
        return 0
    }
    db.expires.Put(key, expireAt)
    return 1
}
