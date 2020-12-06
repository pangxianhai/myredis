package numbers

// 求不大于 n 的最大 2的n次方值
func Max2N32(n uint32) uint32 {
    n |= n >> 1
    n |= n >> 2
    n |= n >> 4
    n |= n >> 8
    n |= n >> 16
    return (n + 1) >> 1
}

// 求不大于 n 的最大 2的n次方值
func Max2N64(n uint64) uint64 {
    n |= n >> 1
    n |= n >> 2
    n |= n >> 4
    n |= n >> 8
    n |= n >> 16
    n |= n >> 32
    return (n + 1) >> 1
}
