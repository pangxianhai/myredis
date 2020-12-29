package db

import (
    "redis-server/common/config"
    "redis-server/common/logger"
    "redis-server/data/object"
    "redis-server/data/sds"
    "testing"
)

func TestPut(t *testing.T) {
    config.Load("./redis.conf")
    logger.Init()
    Init()
    key := sds.NewWithStr("a")
    value := sds.NewWithStr("b")
    Put(key, object.New(value))

    v := Get(key)
    t.Log(v.SdsVal())
}
