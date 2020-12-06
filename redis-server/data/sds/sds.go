package sds

const M = 1024 * 1024
const M3 = 3 * M

type Sds struct {
    buf  []byte //字符串数据
    len  int    //已使用字节长度
    free int    //未使用数量
}

func New() *Sds {
    sds := new(Sds)
    sds.buf = make([]byte, 16)
    sds.len = 0
    sds.free = len(sds.buf)
    return sds
}

func NewWithStr(str string) *Sds {
    sds := new(Sds)
    sds.buf = make([]byte, len(str)*2)
    copy(sds.buf, str)
    sds.len = len(str)
    sds.free = sds.len * 2
    return sds
}

func (sds *Sds) Len() int {
    return sds.len
}

func (sds *Sds) String() string {
    return string(sds.buf[0:sds.len])
}

func (sds *Sds) Cat(str string) {
    sds.extend(len(str))
    copy(sds.buf[sds.len:], str)
    sds.len += len(str)
    sds.free -= len(str)
}

func (sds *Sds) CatSds(other *Sds) {
    sds.extend(other.len)
    copy(sds.buf[sds.len:], other.buf[0:other.len])
    sds.len += other.len
    sds.free -= other.len
}

func (sds *Sds) extend(len int) {
    if sds.free >= len {
        return
    }
    if sds.len+len > M3 {
        //扩展1M
        sds.buf = append(sds.buf, make([]byte, M+len-sds.free)...)
        sds.free = M + len
    } else {
        //扩两倍
        sds.buf = append(sds.buf, make([]byte, (sds.len+len)*2-sds.free)...)
        sds.free = (sds.len + len) * 2
    }
}

func (sds *Sds) freed() {
    if sds.free > 2*sds.len {
        sds.buf = sds.buf[0 : 2*sds.len]
        sds.free = len(sds.buf) - sds.len
    }
}
