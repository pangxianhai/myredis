package db

import "redis-server/data/object"

type Db struct {
    dict map[string]*object.Object
}

var db Db

func new() {
    db = Db{
        dict: make(map[string]*object.Object, 1),
    }
}

func GetByKey(key string) *object.Object {
    return db.dict[key]
}
