package linkedlist_test

import (
	"reflect"
	"testing"

	"github.com/martishin/golang-algorithms/linkedlist"
)

func TestNew(t *testing.T) {
	list := linkedlist.New[int]()

	if list.Head != nil || list.Tail != nil {
		t.Fatalf("Head and tail should be nil after initialization")
	}
}

func TestAddToEmptyList(t *testing.T) {
	list := linkedlist.New[int]()

	list.Add(1)

	if list.Head == nil || list.Tail == nil {
		t.Fatalf("Expected head and tail to be non-nil after adding a value")
	}

	if list.Head != list.Tail {
		t.Fatalf("Expected head and tail to point to the same node for a single element list")
	}

	if list.Head.Val != 1 {
		t.Errorf("Expected node value to be 1, got %v", list.Head.Val)
	}
}

func TestAddToNonEmptyList(t *testing.T) {
	list := linkedlist.New[int]()

	list.Add(1)
	list.Add(2)

	if list.Head == nil || list.Tail == nil {
		t.Fatalf("Expected head and tail to be non-nil after adding values")
	}

	if list.Head == list.Tail {
		t.Fatalf("Expected head and tail to point to different nodes for a two element list")
	}

	if list.Head.Val != 1 {
		t.Errorf("Expected head node value to be 1, got %v", list.Head.Val)
	}
	if list.Tail.Val != 2 {
		t.Errorf("Expected tail node value to be 2, got %v", list.Tail.Val)
	}
}

func TestGet(t *testing.T) {
	list := linkedlist.New[int]()

	list.Add(1)
	list.Add(2)

	node := list.Get(1)

	if node == nil {
		t.Fatalf("Expected to find node with value 2")
	}

	if node.Val != 2 {
		t.Errorf("Expected node value to be 2, got %v", node.Val)
	}
}

func TestToSlice(t *testing.T) {
	list := linkedlist.New[int]()

	list.Add(1)
	list.Add(2)
	list.Add(3)

	slice := list.ToSlice()

	if !reflect.DeepEqual(slice, []int{1, 2, 3}) {
		t.Errorf("Expected list to be [1, 2, 3], got %v", slice)
	}
}

func TestReverse(t *testing.T) {
	list := linkedlist.New[int]()

	list.Add(1)
	list.Add(2)
	list.Add(3)

	list.Reverse()

	reversed := list.ToSlice()

	if !reflect.DeepEqual(reversed, []int{3, 2, 1}) {
		t.Errorf("Expected reversed list to be [3, 2, 1], got %v", reversed)
	}
}
