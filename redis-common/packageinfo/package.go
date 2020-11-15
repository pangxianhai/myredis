package packageinfo

import "redis-common/util"

// PackageInfo 定义数据包格式
type PackageInfo struct {
    MsgID  uint32
    Len    uint32
    Accept uint32
    Buf    []byte
}

// ByteToPackage 将二进制数据转换为 PackageInfo 对象
// buf 格式为 4位MsgID + 4位len + 数据
// 如果发生包解析错误 直接返回 error 服务端收到该错误可以主动与客户端断开连接 客户端重连服务器
func ByteToPackage(lastPackage *PackageInfo, buf []byte) ([]PackageInfo, error) {
    bufL := uint32(len(buf))
    if bufL == 0 {
        return nil, nil
    }
    var newBuf []byte
    if lastPackage != nil && lastPackage.Buf != nil {
        newBuf = append(lastPackage.Buf, buf...)
        //清空元数据
        lastPackage.Buf = []byte{}
    } else {
        newBuf = buf
    }
    var newPackageList []PackageInfo
    for len(newBuf) > 0 {
        var err error
        newBuf, err = parseLastPackage(lastPackage, newBuf)
        if err != nil {
            return nil, err
        }
        var newPackage *PackageInfo
        newPackage, newBuf, err = parsePackage(newBuf)
        if err != nil {
            return nil, err
        }
        if newPackage != nil {
            newPackageList = append(newPackageList, *newPackage)
        }
    }
    return newPackageList, nil
}

// PackageToByte 将 PackageInfo 对象转换二进制
func PackageToByte(packageInfo *PackageInfo) ([]byte, error) {
    buf := make([]byte, 0)
    b, err := util.Uint32ToBytes(packageInfo.MsgID)
    if err != nil {
        return nil, err
    }
    buf = append(buf, b...)

    b, err = util.Uint32ToBytes(packageInfo.Len)
    if err != nil {
        return nil, err
    }
    buf = append(buf, b...)
    buf = append(buf, packageInfo.Buf...)
    return buf, nil
}

func parseLastPackage(lastPackage *PackageInfo, buf []byte) ([]byte, error) {
    if lastPackage == nil {
        //数据没有发生变化
        return buf, nil
    }
    if lastPackage.MsgID == 0 {
        if len(buf) < 4 {
            //数据不够 数据存于 lastPackage.Buf中
            lastPackage.Buf = buf
            return nil, nil
        }
        id, err := util.ByteToUint32(buf[0:4])
        if err != nil {
            return nil, err
        }
        lastPackage.MsgID = id
        //更新剩余数据
        buf = buf[4:]
    }
    if lastPackage.Len == 0 {
        if len(buf) < 4 {
            //数据不够 数据存于 lastPackage.Buf中
            lastPackage.Buf = buf
            return nil, nil
        }
        l, err := util.ByteToUint32(buf[0:4])
        if err != nil {
            return nil, err
        }
        lastPackage.Len = l
        //更新剩余数据
        buf = buf[4:]
    }
    bufL := uint32(len(buf))
    if bufL == 0 {
        return nil, nil
    }
    needLen := lastPackage.Len - lastPackage.Accept
    if needLen == 0 {
        return buf, nil
    }
    if bufL <= needLen {
        lastPackage.Buf = append(lastPackage.Buf, buf...)
        lastPackage.Accept += bufL
        //buf中数据已经用完
        return nil, nil
    } else {
        lastPackage.Buf = append(lastPackage.Buf, buf[0:needLen]...)
        lastPackage.Accept += needLen
        buf = buf[needLen:]
        return buf, nil
    }
}

func parsePackage(buf []byte) (*PackageInfo, []byte, error) {
    if len(buf) <= 0 {
        return nil, nil, nil
    }
    p := PackageInfo{}
    var err error
    buf, err = parseLastPackage(&p, buf)
    if err != nil {
        return nil, nil, err
    } else {
        return &p, buf, nil
    }
}
