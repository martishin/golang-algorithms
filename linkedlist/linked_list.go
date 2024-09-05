package linkedlist

type Node[T comparable] struct {
	Val  T
	Next *Node[T]
	Prev *Node[T]
}

type List[T comparable] struct {
	Front *Node[T]
	Back  *Node[T]
	Len   int
}

func New[T comparable]() *List[T] {
	return &List[T]{
		Front: nil,
		Back:  nil,
		Len:   0,
	}
}

func (l *List[T]) PushBack(val T) *Node[T] {
	var newNode *Node[T]

	if l.Back != nil {
		newNode = &Node[T]{
			Val:  val,
			Next: nil,
			Prev: l.Back,
		}

		l.Back.Next = newNode
		l.Back = newNode
	} else {
		newNode = &Node[T]{
			Val:  val,
			Next: nil,
			Prev: nil,
		}

		l.Front = newNode
		l.Back = newNode
	}
	l.Len++

	return newNode
}

func (l *List[T]) PushFront(val T) *Node[T] {
	var newNode *Node[T]

	if l.Front != nil {
		newNode = &Node[T]{
			Val:  val,
			Next: l.Front,
			Prev: nil,
		}

		l.Front.Prev = newNode
		l.Front = newNode
	} else {
		newNode = &Node[T]{
			Val:  val,
			Next: nil,
			Prev: nil,
		}

		l.Front = newNode
		l.Back = newNode
	}
	l.Len++

	return newNode
}

func (l *List[T]) Find(val T) *Node[T] {
	current := l.Front
	for current != nil {
		if current.Val == val {
			return current
		}
		current = current.Next
	}
	return nil
}

func (l *List[T]) Contains(val T) bool {
	found := l.Find(val)

	return found != nil
}

func (l *List[T]) Remove(node *Node[T]) bool {
	if node != nil {
		if node.Next != nil {
			node.Next.Prev = node.Prev
		} else {
			l.Back = node.Prev
		}

		if node.Prev != nil {
			node.Prev.Next = node.Next
		} else {
			l.Front = node.Next
		}
		l.Len--

		return true
	}

	return false
}

func (l *List[T]) Reverse() {
	current := l.Front
	var prev *Node[T]
	l.Back = l.Front

	for current != nil {
		next := current.Next
		current.Next = prev
		current.Prev = next
		prev = current
		current = next
	}
	l.Front = prev
}

func (l *List[T]) InsertAfter(newVal T, node *Node[T]) *Node[T] {
	if node == nil {
		return nil
	}

	newNode := &Node[T]{
		Val:  newVal,
		Next: node.Next,
		Prev: node,
	}
	if node.Next != nil {
		node.Next.Prev = newNode
	} else {
		l.Back = newNode
	}
	node.Next = newNode
	l.Len++

	return newNode
}

func (l *List[T]) InsertBefore(newVal T, node *Node[T]) *Node[T] {
	if node == nil {
		return nil
	}

	newNode := &Node[T]{
		Val:  newVal,
		Next: node,
		Prev: node.Prev,
	}
	if node.Prev != nil {
		node.Prev.Next = newNode
	} else {
		l.Front = newNode
	}
	node.Prev = newNode
	l.Len++

	return newNode
}

func (l *List[T]) Clear() {
	l.Front = nil
	l.Back = nil
	l.Len = 0
}

func (l *List[T]) Traverse(f func(T)) {
	current := l.Front
	for current != nil {
		f(current.Val)
		current = current.Next
	}
}

func (l *List[T]) ToSlice() []T {
	var result []T
	l.Traverse(func(val T) {
		result = append(result, val)
	})
	return result
}

func (l *List[T]) Get(index int) *Node[T] {
	current := l.Front

	for range index {
		if current == nil {
			return nil
		}
		current = current.Next
	}

	return current
}
