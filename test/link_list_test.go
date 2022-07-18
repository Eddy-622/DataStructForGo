package test

import (
	"DataStructures/data_struct"
	"testing"
)

func TestLinkList(t *testing.T) {
	list := data_struct.LinkedList{}
	list.AddFirst(1)
	list.AddFirst(2)
	list.AddFirst(3)
	list.AddLast("abc")
	t.Log(list.Len())
	t.Log(list.Value())
	t.Log(list.Contains("ccc")) // false
	t.Log(list.ValueOfIndex(3)) // abc
	t.Log(list.IndexOfValue(3)) // 0
	//list.DelFirst()
	//t.Log(list.Value()) // false
	//list.DelLast()
	//t.Log(list.Value())
	//t.Log(list.Len())
	list.Reverse()
	t.Log(list.Value())
	t.Log(list.FindKthNodeFromEnd(3))
}
