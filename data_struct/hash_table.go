package data_struct

import (
	"fmt"
)

type HashTable struct {
	list   []*LinkList
	size   int
	kvSize int
}

func NewHashTable(size int) *HashTable {
	return &HashTable{make([]*LinkList, size), size, 0}
}

func (h *HashTable) Put(k int, v string) {
	entry := h.Entry(k, v)

	h.LL(k).Put(entry)
	h.kvSize++
}

func (h *HashTable) Get(k int) interface{} {
	ll := h.LL(k)

	current := ll.first
	for current != nil {
		if entry, ok := current.value.(*entry); ok && k == entry.k {
			return entry.v
		}
	}
	return nil
}

func (h *HashTable) Remove(k int) interface{} {
	ll := h.LL(k)

	current := ll.first
	for current != nil {
		if entry, ok := current.value.(*entry); ok && entry.k == k {
			switch current {
			case ll.first:
				ll.DelHead()
			case ll.last:
				ll.Pop()
			default:
				before := ll.first
				for before != nil {
					if before.next == current {
						before.next = current.next
						break
					}
					before = before.next
				}
			}
			h.kvSize--
			return entry.v
		}

		current = current.next
	}

	return nil
}

type entry struct {
	k int
	v string
}

func (h *HashTable) Entry(k int, v string) *entry {
	return &entry{k, v}
}

func (h *HashTable) LL(k int) *LinkList {
	if h.list[k%h.size] == nil {
		h.list[k%h.size] = NewLinkList()
	}
	return h.list[k%h.size]
}

func (h *HashTable) Print() {
	printSlice := make([]string, h.kvSize)
	var i int
	for _, ll := range h.list {
		if ll == nil {
			continue
		}
		current := ll.first
		for current != nil {
			if entry, ok := current.value.(*entry); ok {
				str := fmt.Sprintf("%v:%v", entry.k, entry.v)
				printSlice[i] = str
				i++
			}
			current = current.next
		}
	}
	fmt.Println(printSlice)
}
