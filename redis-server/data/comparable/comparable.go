package comparable

type Comparable interface {
    // return 0 两元素相等
    // -1 小于
    // 1 大于
    CompareTo(c interface{}) int
}
