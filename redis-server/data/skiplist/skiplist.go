package skiplist

import (
    "fmt"
    "math"
    "math/rand"
    "redis-server/common/iterator"
    "redis-server/data/comparable"
    "redis-server/data/list"
    "strconv"
    "time"
)

const (
    maxLevel = 32 //最大层
    randMax  = 2 ^ 32
)

type Node struct {
    v     comparable.Comparable //数据
    score float64               //分数
    bw    *Node                 //后退指针
    level []*LevelNode          //层

}

//层节点
type LevelNode struct {
    forward *Node //前进指针
    span    int   //跨度
}

type SkipList struct {
    head  *Node //头结点
    tail  *Node //尾结点
    len   int   //长度
    level int   //最高层数
    rand  *rand.Rand
}

type Iterator struct {
    list *SkipList
    cur  *Node
    next *Node
}

func (n *Node) CompareTo(o interface{}) int {
    n1, ok := o.(*Node)
    if !ok {
        return -1
    }
    if n.score < n1.score {
        return -1
    } else if n.score > n1.score {
        return 1
    } else {
        return n.v.CompareTo(n1.v)
    }
}

func New() *SkipList {
    skipList := new(SkipList)
    skipList.rand = rand.New(rand.NewSource(time.Now().UnixNano()))
    skipList.head = skipList.createNode(nil, 0, maxLevel)
    return skipList
}

func (skipList *SkipList) Add(v comparable.Comparable, score float64) {
    nLevel := skipList.randLevel()
    n := skipList.createNode(v, score, nLevel)
    update, rank := skipList.search(n)

    if nLevel > skipList.level {
        for i := skipList.level; i < nLevel; i++ {
            update[i] = skipList.head
            update[i].level[i].span = skipList.len
            rank[i] = 0
        }
        skipList.level = nLevel
    }

    //记录需要修改层节点 层指向新节点
    for k := 0; k < nLevel; k++ {
        n.level[k].forward = update[k].level[k].forward
        update[k].level[k].forward = n

        n.level[k].span = update[k].level[k].span - (rank[0] - rank[k])
        update[k].level[k].span = (rank[0] - rank[k]) + 1
    }
    //未接触的节点的 span 值也需要增一
    for i := nLevel; i < skipList.level; i++ {
        update[i].level[i].span++
    }
    if update[0] == skipList.head {
        n.bw = nil
    } else {
        n.bw = update[0]
    }
    if n.level[0].forward != nil {
        n.level[0].forward.bw = n
    } else {
        skipList.tail = n
    }
    skipList.len++
}

func (skipList *SkipList) Remove(v comparable.Comparable, score float64) {
    n := skipList.createNode(v, score, 0)
    update, _ := skipList.search(n)
    cur := update[0].level[0].forward
    if cur == nil || cur.CompareTo(n) != 0 {
        //不存在 无需删除
        return
    }
    for i := 0; i < len(cur.level); i++ {
        update[i].level[i].forward = cur.level[i].forward
        update[i].level[i].span = update[i].level[i].span + cur.level[i].span - 1
    }
    for i := len(cur.level); i < len(update); i++ {
        if update[i] != nil && update[i].level[i] != nil {
            update[i].level[i].span--
        }
    }
    if cur.level[0].forward != nil {
        cur.level[0].forward.bw = update[0]
    }
    if cur.CompareTo(skipList.tail) == 0 {
        skipList.tail = cur.bw
    }
    for i := skipList.level - 1; i >= 0; i-- {
        if skipList.head.level[i].forward == nil {
            //改层设置为无效
            skipList.head.level[i].span = 0
            skipList.level--
        } else {
            break
        }
    }
    skipList.len--
}

func (skipList *SkipList) Range(start, end int) *list.List {
    start, end, ok := skipList.rewordStartAndEnd(start, end)

    if !ok {
        //偏移量不合法 返回 nil
        return nil
    }
    node := skipList.getIndex(skipList.head, start+1)

    resList := list.New()
    for i := 0; i <= end-start; i++ {
        if node == nil {
            break
        }
        resList.Rpush(node)
        node = node.level[0].forward
    }
    return resList
}

func (skipList *SkipList) RevRange(start, end int) *list.List {
    start, end, ok := skipList.rewordStartAndEnd(start, end)

    if !ok {
        //偏移量不合法 返回 nil
        return nil
    }
    node := skipList.getIndex(skipList.head, end+1)
    resList := list.New()
    for i := 0; i <= end-start; i++ {
        if node == nil {
            break
        }
        resList.Rpush(node)
        node = node.bw
    }
    return resList
}

func (skipList *SkipList) RangeByScore(min, max float64, offset, count int) *list.List {
    if offset < 0 || offset > skipList.len || count <= 0 {
        return nil
    }
    if min > max || min > skipList.tail.score {
        return nil
    }
    node, _ := skipList.getScore(min)
    if node == nil {
        return nil
    }
    node = skipList.getIndex(node, offset)
    resList := list.New()
    for i := 0; i < count; i++ {
        if node == nil || node.score > max {
            break
        }
        resList.Rpush(node)
        node = node.level[0].forward
    }
    return resList
}

func (skipList *SkipList) RevRangeByScore(min, max float64, offset, count int) *list.List {
    if offset < 0 || offset > skipList.len || count <= 0 {
        return nil
    }
    if min > max || min > skipList.tail.score {
        return nil
    }
    node, span := skipList.getScore(max)
    if node == nil {
        return nil
    }
    node = skipList.getIndex(skipList.head, span)

    resList := list.New()
    for i := 0; i < count; i++ {
        if node == nil || node.score < min {
            break
        }
        resList.Rpush(node)
        node = node.bw
    }
    return resList
}

func (skipList *SkipList) Iterator() iterator.Iterator {
    ite := new(Iterator)
    ite.list = skipList
    ite.next = skipList.head.level[0].forward
    return ite
}

func (ite *Iterator) HasNext() bool {
    return ite.next != nil
}

func (ite *Iterator) Next() interface{} {
    ite.cur = ite.next
    ite.next = ite.cur.level[0].forward
    return ite.cur
}

func (n *Node) String() string {
    levelStr := ""
    for i := 0; i < len(n.level); i++ {
        levelStr += "[层数:" + strconv.Itoa(i) + ",跨度" + strconv.Itoa(n.level[i].span) + ",下一个:"
        if n.level[i].forward != nil {
            levelStr += fmt.Sprint(n.level[i].forward.v) + "] "
        } else {
            levelStr += "nil] "
        }
    }
    return fmt.Sprint("值:", n.v, ";分数:", n.score, "; 层:"+levelStr)
}

func (skipList *SkipList) search(t *Node) (update []*Node, rank []int) {
    l := skipList.level
    update = make([]*Node, maxLevel)
    rank = make([]int, maxLevel)
    n := skipList.head
    for l = l - 1; l >= 0; l-- {
        if l == skipList.level-1 {
            rank[l] = 0
        } else {
            rank[l] = rank[l+1]
        }
        for n.level[l] != nil && n.level[l].forward != nil && n.level[l].forward.CompareTo(t) < 0 {
            rank[l] = rank[l] + n.level[l].span
            n = n.level[l].forward
        }
        update[l] = n
    }
    return
}

func (skipList *SkipList) getIndex(startNode *Node, index int) *Node {
    if index < 0 || index >= skipList.len {
        return nil
    }
    //由于head 到第一个的跨度为1 所以 index 要加1
    span := 0
    for span < index && startNode != skipList.tail {
        for l := len(startNode.level) - 1; l >= 0; l-- {
            for index >= startNode.level[l].span+span && startNode.level[l].forward != nil {
                span = startNode.level[l].span + span
                startNode = startNode.level[l].forward
            }
        }
    }
    return startNode
}

// getScore 查询 <= score 的分数最大的节点 或 大于>= score 分数最小的节点 其实这两条件的含义是一样的
func (skipList *SkipList) getScore(score float64) (node *Node, span int) {
    node = skipList.head
    for l := skipList.level - 1; l >= 0; l-- {
        for node.level[l].forward != nil && node.level[l].forward.score <= score {
            span += node.level[l].span
            node = node.level[l].forward
        }
        if node.score >= score {
            break
        }
    }
    //node 可能是最后一个节点 node 不能继续向后移动
    if node.score < score {
        node = nil
    }
    return
}

func (skipList *SkipList) createNode(v comparable.Comparable, score float64, level int) *Node {
    n := new(Node)
    n.score = score
    n.v = v
    n.level = make([]*LevelNode, level)
    for i := 0; i < level; i++ {
        n.level[i] = &LevelNode{}
    }
    return n
}

func (skipList *SkipList) randLevel() int {
    i, t := 1, skipList.rand.Intn(randMax)
    for j := 2; i <= int(math.Log2(float64(skipList.len))); i, j = i+1, j+j {
        if t > randMax/j {
            break
        }
    }
    if i > maxLevel {
        i = maxLevel
    }
    return i
}

func (skipList *SkipList) String() string {

    str := fmt.Sprintf("[len:%d,level:%d", skipList.len, skipList.level)
    for ite := skipList.Iterator(); ite.HasNext(); {
        str += fmt.Sprint(ite.Next())
    }
    str += "]"
    return str
}

func (skipList *SkipList) rewordStartAndEnd(start, end int) (int, int, bool) {
    if end < 0 {
        //end 小于0 表示到倒数end位
        end = skipList.len + end + 1
    }
    if start < 0 {
        //start 小于0 表示到倒数start位
        start = skipList.len + start
    }
    if start < 0 || start >= skipList.len || end < 0 || start > end {
        //偏移量不合法 返回 nil
        return start, end, false
    } else {
        return start, end, true
    }
}
