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

	if list.Head.Next != nil || list.Head.Prev != nil {
		t.Errorf("Expected next and prev pointers of the node to be nil, but got Next: %v, Prev: %v", list.Head.Next, list.Head.Prev)
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

	if list.Head.Next != list.Tail {
		t.Errorf("Expected head's next to point to the tail, but got %v", list.Head.Next)
	}
	if list.Tail.Prev != list.Head {
		t.Errorf("Expected tail's prev to point to the head, but got %v", list.Tail.Prev)
	}

	if list.Head.Prev != nil {
		t.Errorf("Expected head's prev to be nil, but got %v", list.Head.Prev)
	}
	if list.Tail.Next != nil {
		t.Errorf("Expected tail's next to be nil, but got %v", list.Tail.Next)
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
