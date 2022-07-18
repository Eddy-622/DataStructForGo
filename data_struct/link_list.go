package data_struct

type node struct {
	value interface{}
	next  *node
}

type LinkList struct {
	first *node
	last  *node
	size  int
}

func NewLinkList() *LinkList {
	return &LinkList{}
}

func (l *LinkList) Put(v interface{}) {
	node := &node{v, nil}
	if l.IsEmpty() {
		l.first = node
		l.last = node
	} else {
		l.last.next = node
		l.last = node
	}
	l.size++
}

func (l *LinkList) Pop() interface{} {
	if l.IsEmpty() {
		return nil
	}
	lastValue := l.last.value

	if l.first == l.last {
		l.first = nil
		l.last = nil
		l.size--
		return lastValue
	}

	current := l.first
	for i := 0; i < l.size-1; i++ {
		current = current.next
	}
	current.next = nil
	l.last = current
	l.size--
	return lastValue
}

func (l *LinkList) DelHead() interface{} {
	if l.IsEmpty() {
		return nil
	}
	lastValue := l.first.value
	if l.first == l.last {
		l.first = nil
		l.last = nil
		l.size--
		return lastValue
	}

	next := l.first.next
	l.first.next = nil
	l.first = next
	return lastValue
}

func (l *LinkList) IsEmpty() bool {
	return l.size == 0
}

func (l *LinkList) Len() int {
	return l.size
}
