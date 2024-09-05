package linkedlist

type Node[T comparable] struct {
	Val  T
	Next *Node[T]
	Prev *Node[T]
}

type List[T comparable] struct {
	Head *Node[T]
	Tail *Node[T]
	Size int
}

func New[T comparable]() *List[T] {
	return &List[T]{
		Head: nil,
		Tail: nil,
		Size: 0,
	}
}

func (l *List[T]) Add(val T) {
	if l.Tail != nil {
		newNode := &Node[T]{
			Val:  val,
			Next: nil,
			Prev: l.Tail,
		}
		l.Tail.Next = newNode
		l.Tail = newNode
	} else {
		newNode := &Node[T]{
			Val:  val,
			Next: nil,
			Prev: nil,
		}
		l.Head = newNode
		l.Tail = newNode
	}
	l.Size++
}

func (l *List[T]) Find(val T) *Node[T] {
	current := l.Head
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

func (l *List[T]) Remove(val T) bool {
	found := l.Find(val)

	if found != nil {
		if found.Next != nil {
			found.Next.Prev = found.Prev
		} else {
			l.Tail = found.Prev
		}

		if found.Prev != nil {
			found.Prev.Next = found.Next
		} else {
			l.Head = found.Next
		}

		l.Size--
		return true
	}

	return false
}

func (l *List[T]) Reverse() {
	current := l.Head
	var prev *Node[T]
	l.Tail = l.Head

	for current != nil {
		next := current.Next
		current.Next = prev
		current.Prev = next
		prev = current
		current = next
	}
	l.Head = prev
}

func (l *List[T]) InsertAfter(targetVal T, newVal T) {
	current := l.Find(targetVal)
	if current != nil {
		newNode := &Node[T]{
			Val:  newVal,
			Next: current.Next,
			Prev: current,
		}
		if current.Next != nil {
			current.Next.Prev = newNode
		}
		current.Next = newNode
		if current == l.Tail {
			l.Tail = newNode
		}
	}
}

func (l *List[T]) Clear() {
	l.Head = nil
	l.Tail = nil
}

func (l *List[T]) Traverse(f func(T)) {
	current := l.Head
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
	current := l.Head

	for range index {
		if current == nil {
			return nil
		}
		current = current.Next
	}

	return current
}
