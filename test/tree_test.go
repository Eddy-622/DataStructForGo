package test

import (
	"DataStructures/data_struct"
	"math/rand"
	"testing"
	"time"
)

func TestTree(t *testing.T) {
	bt := data_struct.BinaryTree{}
	rand.Seed(time.Now().Unix())
	t.Log(bt.Insert(rand.Intn(100)))
	t.Log(bt.Insert(rand.Intn(100)))
	t.Log(bt.Insert(rand.Intn(100)))
	t.Log(bt.Insert(rand.Intn(100)))
	t.Log(bt.Insert(rand.Intn(100)))

	bt.Insert(1000)

	t.Log(bt.Size)
	t.Log("Root:", bt.Root.Value)
	t.Log("left", bt.Values("left"))
	t.Log("right", bt.Values("right"))
	t.Log("Isin", bt.Find(1000))

	bt.TracersInOrder()
	t.Log(bt.Height())

	t.Log("=======================")

	bt1 := data_struct.BinaryTree{}
	bt2 := data_struct.BinaryTree{}
	bt1.Insert(10)
	bt1.Insert(30)
	bt1.Insert(20)
	bt1.Insert(50)

	bt1.Insert(10)
	bt1.Insert(20)
	bt1.Insert(20)
	bt1.Insert(50)
	t.Log(bt1.Equals(&bt2))

	t.Log("===================")
	t.Log(bt1.IsBinarySearchTree())
	t.Log(bt1.GetDistanceArray(0))
	bt1.PrintALL()
}
