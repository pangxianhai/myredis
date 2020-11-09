package packageinfo

// PackageInfo 定义数据包格式
type PackageInfo struct {
    MsgID  uint32
    len    uint32
    accpet uint32
    Buf    []byte
}

// ByteToPackage 将二进制数据转换为 PackageInfo 对象
func (packageInfo *PackageInfo) ByteToPackage(buf []byte) []PackageInfo {
    return nil
}

// ToByte 将 PackageInfo 对象转换二进制
func (packageInfo *PackageInfo) ToByte() []byte {
    
    return nil
}
