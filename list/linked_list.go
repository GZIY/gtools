package list

import "github.com/GZIY/gtools/errs"

var (
	_ List[any] = &LinkedList[any]{}
)

// 双向循环链表的节点
type node[T any] struct {
	prev *node[T]
	next *node[T]
	val  T
}

// 双向循环链表
type LinkedList[T any] struct {
	// 虚头虚尾，不存储任何元素
	head   *node[T]
	tail   *node[T]
	length int
}

// NewLinkedList 创建一个双向循环链表
func NewLinkedList[T any]() *LinkedList[T] {
	head := &node[T]{}
	tail := &node[T]{}

	head.prev, head.next = tail, tail
	tail.prev, tail.next = head, head

	return &LinkedList[T]{
		head:   head,
		tail:   tail,
		length: 0,
	}
}

// NewLinkedListOf 将切片转换为双向循环链表, 直接使用了切片元素的值，而没有进行复制
func NewLinkedListOf[T any](ts []T) *LinkedList[T] {
	list := NewLinkedList[T]()

	err := list.Append(ts...)
	if err != nil {
		panic(err)
	}

	return list
}

func (l *LinkedList[T]) Get(idx int) (T, error) {
	if !l.checkIndex(idx) {
		// idx 越界
		var zeroVlue T
		return zeroVlue, errs.NewErrIndexOutOfRange(l.Len(), idx)
	}

	// 不越界，则找节点
	n := l.findNode(idx)
	return n.val, nil
}

func (l *LinkedList[T]) Append(ts ...T) error {
	for _, t := range ts {

		// 在尾部追加节点
		node := &node[T]{
			prev: l.tail.prev,
			next: l.tail,
			val:  t,
		}

		// 调整指向
		node.prev.next = node
		node.next.prev = node
		l.length++
	}

	return nil
}

func (l *LinkedList[T]) Add(idx int, t T) error {
	if idx < 0 || idx > l.Len() {
		return errs.NewErrIndexOutOfRange(l.Len(), idx)
	}

	if idx == l.Len() {
		// idx = 长度 则 等同于 append
		return l.Append(t)
	}

	// 找到 idx 下标的 node, 是插入节点的 下一个node
	next := l.findNode(idx)

	// 插入节点创建 并 插入
	node := &node[T]{
		prev: next.prev,
		next: next,
		val:  t,
	}

	// 调整指向
	node.prev.next = node
	node.next.prev = node
	l.length++

	return nil
}

func (l *LinkedList[T]) Set(idx int, t T) error {
	if !l.checkIndex(idx) {
		return errs.NewErrIndexOutOfRange(l.Len(), idx)
	}

	node := l.findNode(idx)
	node.val = t
	return nil
}

func (l *LinkedList[T]) Delete(idx int) (T, error) {
	if !l.checkIndex(idx) {
		var zeroVal T
		return zeroVal, errs.NewErrIndexOutOfRange(l.Len(), idx)
	}

	// 找到删除节点
	n := l.findNode(idx)
	n.prev.next = n.next
	n.next.prev = n.prev

	n.next, n.prev = nil, nil

	l.length--

	return n.val, nil
}

func (l *LinkedList[T]) Len() int {
	return l.length
}

func (l *LinkedList[T]) Cap() int {
	return l.Len()
}

func (l *LinkedList[T]) Range(fn func(idx int, t T) error) error {
	cur := l.head.next
	i := 0

	for cur != l.tail {
		err := fn(i, cur.val)
		if err != nil {
			return err
		}

		cur = cur.next
		i++
	}

	return nil
}

func (l *LinkedList[T]) AsSlice() []T {
	// 长度 容量均 l.length
	s := make([]T, l.length)

	for cur, i := l.head.next, 0; i < l.length; i++ {
		s[i] = cur.val
		cur = cur.next
	}

	return s
}

// 检查一个 索引 是否 越界
func (l *LinkedList[T]) checkIndex(idx int) bool {
	return 0 <= idx && idx < l.Len()
}

// 通过索引找节点
func (l *LinkedList[T]) findNode(idx int) *node[T] {
	// 由于是循环双向链表
	// idx <= 长度 / 2, 则顺序找， 否则逆序找
	var cur *node[T]
	if idx <= l.Len()/2 {
		// 顺序找
		cur = l.head
		for i := 0; i <= idx; i++ {
			cur = cur.next
		}

	} else {
		// 逆序找
		cur = l.tail
		for i := l.Len(); i > idx; i-- {
			cur = cur.prev
		}
	}

	return cur
}
