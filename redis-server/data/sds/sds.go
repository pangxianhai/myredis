package sds

const M = 1024 * 1024
const M3 = 3 * M

type Sds struct {
    Buf  []byte //字符串数据
    len  int    //已使用字节长度
    free int    //未使用数量
}

func New() *Sds {
    sds := new(Sds)
    sds.Buf = make([]byte, 16)
    sds.len = 0
    sds.free = len(sds.Buf)
    return sds
}

func NewWithStr(str string) *Sds {
    sds := new(Sds)
    sds.Buf = make([]byte, len(str)*2)
    copy(sds.Buf, str)
    sds.len = len(str)
    sds.free = sds.len * 2
    return sds
}

func (sds *Sds) Len() int {
    return sds.len
}

func (sds *Sds) String() string {
    return string(sds.Buf[0:sds.len])
}

func (sds *Sds) CompareTo(o interface{}) int {
    o1, ok := o.(*Sds)
    if !ok {
        return -1
    }
    if sds.String() < o1.String() {
        return -1
    } else if sds.String() == o1.String() {
        return 0
    } else {
        return 1
    }
}

func (sds *Sds) Cat(str string) {
    sds.extend(len(str))
    copy(sds.Buf[sds.len:], str)
    sds.len += len(str)
    sds.free -= len(str)
}

func (sds *Sds) CatSds(other *Sds) {
    sds.extend(other.len)
    copy(sds.Buf[sds.len:], other.Buf[0:other.len])
    sds.len += other.len
    sds.free -= other.len
}

func (sds *Sds) extend(len int) {
    if sds.free >= len {
        return
    }
    if sds.len+len > M3 {
        //扩展1M
        sds.Buf = append(sds.Buf, make([]byte, M+len-sds.free)...)
        sds.free = M + len
    } else {
        //扩两倍
        sds.Buf = append(sds.Buf, make([]byte, (sds.len+len)*2-sds.free)...)
        sds.free = (sds.len + len) * 2
    }
}

func (sds *Sds) freed() {
    if sds.free > 2*sds.len {
        sds.Buf = sds.Buf[0 : 2*sds.len]
        sds.free = len(sds.Buf) - sds.len
    }
}

func IsEmpty(sds *Sds) bool {
    return sds == nil || sds.Len() == 0
}
