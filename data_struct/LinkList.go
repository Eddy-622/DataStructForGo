package data_struct

import (
	"errors"
)

// Node 每个节点的属性【自身的数据和下一个节点的地址】
type Node struct {
	value interface{}
	next  *Node
}

type LinkedList struct {
	first *Node
	last  *Node
	size  int
}

var ListNil = errors.New("链表为空")
var IndexOut = errors.New("索引超出范围")

// AddFirst 首位添加
func (l *LinkedList) AddFirst(v interface{}) {
	var node Node
	node.value = v
	if l.IsNone() {
		l.first = &node
		l.last = &node
	} else {
		node.next = l.first
		l.first = &node
	}
	l.size++
}

//AddLast 末尾添加节点
func (l *LinkedList) AddLast(v interface{}) {
	var node Node
	node.value = v
	if l.IsNone() {
		l.last = &node
		l.first = &node
	} else {
		l.last.next = &node
		l.last = &node
	}
	l.size++
}

//DelFirst 删除头部节点
func (l *LinkedList) DelFirst() {
	if l.IsNone() {
		panic(ListNil)
	}
	if l.first == l.last {
		l.first = nil
		l.last = nil
	} else {
		next := l.first.next
		l.first.next = nil
		l.first = next
	}
	l.size--
}

//DelLast 删除尾部节点
func (l *LinkedList) DelLast() {
	if l.IsNone() {
		panic(ListNil)
	}
	if l.first == l.last {
		l.first = nil
		l.last = nil
	} else {
		l.last = l.NodeOfIndex(l.Len() - 2)
		l.last.next = nil
	}
	l.size--
}

//Contains 是否包含v
func (l *LinkedList) Contains(v interface{}) bool {
	if l.IsNone() {
		panic(ListNil)
	}

	current := l.first
	for current != nil {
		if current.value == v {
			return true
		}
		current = current.next
	}
	return false
}

//ValueOfIndex 根绝索引取值
func (l *LinkedList) ValueOfIndex(i int) interface{} {
	if l.Len() <= i {
		panic(IndexOut)
	}

	var index int
	if l.IsNone() {
		panic(ListNil)
	}
	current := l.first
	for current != nil {
		if index == i {
			return current.value
		}
		index += 1
		current = current.next
	}
	return nil
}

// IndexOfValue 根据值取索引 return = -1 表示没有匹配
func (l *LinkedList) IndexOfValue(v interface{}) int {
	if l.IsNone() {
		panic(ListNil)
	}
	var index int
	current := l.first
	for current != nil {
		if v == current.value {
			return index
		}
		current = current.next
		index++
	}
	return -1
}

//NodeOfIndex 索引取节点对象
func (l *LinkedList) NodeOfIndex(i int) *Node {
	if l.Len() <= i {
		panic(IndexOut)
	}

	var index int
	if l.IsNone() {
		panic(ListNil)
	}
	current := l.first
	for current != nil {
		if index == i {
			return current
		}
		index += 1
		current = current.next
	}
	panic(IndexOut)
}

//IsNone 链表是否为空
func (l *LinkedList) IsNone() bool {
	return l.first == nil
}

//Len 返回链表长度
func (l *LinkedList) Len() int {
	return l.size
}

// Value 返回链表的值以数组方式
func (l *LinkedList) Value() []interface{} {
	temp := make([]interface{}, l.Len())
	var i int
	if l.IsNone() {
		return temp
	}
	current := l.first
	for current != nil {
		temp[i] = current.value
		i += 1
		current = current.next
	}
	return temp
}

// Reverse 倒转链表
func (l *LinkedList) Reverse() {
	if l.IsNone() || l.first == l.last {
		return
	}

	// f           l
	// 10 -> 20 -> 30
	//              p     c     n  //循环到此处是c为空就需要终止循环了
	/*
		c.next = p
		p = c
	*/

	p := l.first
	c := p.next
	for c != nil {
		n := c.next
		c.next = p
		p = c
		c = n
	}
	l.first.next = nil
	l.last = l.first
	l.first = p
}

// FindKthNodeFromEnd 查找倒数k节点的值
func (l *LinkedList) FindKthNodeFromEnd(k int) interface{} {
	// k = 3  寻找倒数第三个节点到数据  [双指针]
	// [10 -> 20 -> 30 -> 40 -> 50]
	//              a*  	    b*
	// 首先确定a与b之间当间隔 k-1 [从a到b到距离]
	// 当b指针到达末尾，a指针就是结果
	if k == 0 {
		panic(errors.New("k is a int and must be a gt 0"))
	}
	if l.IsNone() {
		panic(ListNil)
	}

	var a = l.first
	var b = l.first
	for i := 0; i < k-1; i++ {
		b = b.next
		if b == nil {
			panic(IndexOut)
		}
	}
	for b != l.last {
		a = a.next
		b = b.next
	}
	return a.value
}
