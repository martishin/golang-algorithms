package linkedlist_test

import (
	"reflect"
	"testing"

	"github.com/martishin/golang-algorithms/linkedlist"
)

func TestNew(t *testing.T) {
	list := linkedlist.New[int]()

	if list.Front != nil || list.Back != nil {
		t.Fatalf("Front and tail should be nil after initialization")
	}
}

func TestPushBackToEmptyList(t *testing.T) {
	list := linkedlist.New[int]()

	node := list.PushBack(1)

	if node == nil {
		t.Fatalf("Expected node to be non-nil after adding a value")
	}

	if list.Front == nil || list.Back == nil {
		t.Fatalf("Expected head and tail to be non-nil after adding a value")
	}

	if list.Front != list.Back {
		t.Fatalf("Expected head and tail to point to the same node for a single element list")
	}

	if list.Front.Val != 1 {
		t.Errorf("Expected node value to be 1, got %v", list.Front.Val)
	}
}

func TestPushBackToNonEmptyList(t *testing.T) {
	list := linkedlist.New[int]()

	list.PushBack(1)
	list.PushBack(2)

	if list.Front == nil || list.Back == nil {
		t.Fatalf("Expected head and tail to be non-nil after adding values")
	}

	if list.Front == list.Back {
		t.Fatalf("Expected head and tail to point to different nodes for a two element list")
	}

	if list.Front.Val != 1 {
		t.Errorf("Expected head node value to be 1, got %v", list.Front.Val)
	}
	if list.Back.Val != 2 {
		t.Errorf("Expected tail node value to be 2, got %v", list.Back.Val)
	}
}

func TestPushFrontToEmptyList(t *testing.T) {
	list := linkedlist.New[int]()

	node := list.PushFront(1)

	if node == nil {
		t.Fatalf("Expected node to be non-nil after adding a value")
	}

	if list.Front == nil || list.Back == nil {
		t.Fatalf("Expected head and tail to be non-nil after adding a value")
	}

	if list.Front != list.Back {
		t.Fatalf("Expected head and tail to point to the same node for a single element list")
	}

	if list.Front.Val != 1 {
		t.Errorf("Expected node value to be 1, got %v", list.Front.Val)
	}
}

func TestPushFrontToNonEmptyList(t *testing.T) {
	list := linkedlist.New[int]()

	list.PushFront(1)
	list.PushFront(2)

	if list.Front == nil || list.Back == nil {
		t.Fatalf("Expected head and tail to be non-nil after adding values")
	}

	if list.Front == list.Back {
		t.Fatalf("Expected head and tail to point to different nodes for a two element list")
	}

	if list.Front.Val != 2 {
		t.Errorf("Expected head node value to be 2, got %v", list.Front.Val)
	}
	if list.Back.Val != 1 {
		t.Errorf("Expected tail node value to be 1, got %v", list.Back.Val)
	}
}

func TestInsertBefore(t *testing.T) {
	list := linkedlist.New[int]()

	node := list.PushBack(1)
	list.InsertBefore(2, node)

	if list.Front.Val != 2 {
		t.Errorf("Expected head node value to be 2, got %v", list.Front.Val)
	}
	if list.Back.Val != 1 {
		t.Errorf("Expected tail node value to be 1, got %v", list.Back.Val)
	}
}

func TestGet(t *testing.T) {
	list := linkedlist.New[int]()

	list.PushBack(1)
	list.PushBack(2)

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

	list.PushBack(1)
	list.PushBack(2)
	list.PushBack(3)

	slice := list.ToSlice()

	if !reflect.DeepEqual(slice, []int{1, 2, 3}) {
		t.Errorf("Expected list to be [1, 2, 3], got %v", slice)
	}
}

func TestReverse(t *testing.T) {
	list := linkedlist.New[int]()

	list.PushBack(1)
	list.PushBack(2)
	list.PushBack(3)

	list.Reverse()

	reversed := list.ToSlice()

	if !reflect.DeepEqual(reversed, []int{3, 2, 1}) {
		t.Errorf("Expected reversed list to be [3, 2, 1], got %v", reversed)
	}
}
