package data_struct

import (
	"testing"
)

func TestHashtable(t *testing.T) {

	hash := NewHashTable(10)
	hash.Put(1, "a")
	hash.Put(2, "a")
	hash.Put(3, "b")
	hash.Put(13, "bb")
	hash.Remove(2)
	t.Log(hash.Get(1))
	hash.Print()
}

func TestAVLTree(t *testing.T) {
	avl := NewAVLTree()
	avl.Insert(30)
	avl.Insert(20)
	avl.Insert(10)
	t.Log(avl.Root.Value)
	t.Log("Done")
}

func TestHeap(t *testing.T) {
	h := NewHeap(10)
	h.Insert(10)
	h.Insert(5)
	h.Insert(30)
	h.Insert(20)
	h.Insert(8)
	h.Insert(8)
	h.Insert(8)
	h.Remove()

	arr := []int{3, 10, 33, 21}

	maxHeap := new(MaxHeap)
	maxHeap.Heapify(arr)
	t.Log(arr)
	t.Log(maxHeap.GetKthMax(arr, 2))
}
