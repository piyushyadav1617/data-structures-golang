package linkedlist

type ListNode struct {
	Val  int
	Next *ListNode
}

type LinkedList struct {
	Head   *ListNode
	Tail   *ListNode
	Length int
}

func (l *LinkedList) Push(val int) {
	node := &ListNode{Val: val}
	if l.Head == nil {
		l.Head = node
		l.Tail = node
	} else {
		l.Tail.Next = node
		l.Tail = node
	}
	l.Length++
}

func (l *LinkedList) Pop() *ListNode {
	if l.Head == nil {
		return nil
	}
	current := l.Head
	newTail := current
	for current.Next != nil {
		newTail = current
		current = current.Next
	}
	l.Tail = newTail
	l.Tail.Next = nil
	l.Length--
	if l.Length == 0 {
		l.Head = nil
		l.Tail = nil
	}
	return current
}

func (l *LinkedList) Shift() *ListNode {
	if l.Head == nil {
		return nil
	}
	currentHead := l.Head
	l.Head = currentHead.Next
	l.Length--
	if l.Length == 0 {
		l.Tail = nil
	}
	return currentHead
}

func (l *LinkedList) Unshift(val int) {
	node := &ListNode{Val: val}
	if l.Head == nil {
		l.Head = node
		l.Tail = node
	} else {
		node.Next = l.Head
		l.Head = node
	}
	l.Length++
}

func (l *LinkedList) Get(index int) *ListNode {
	if index < 0 || index >= l.Length {
		return nil
	}
	i := 0
	current := l.Head
	for i < index {
		current = current.Next
		i++
	}
	return current
}

func (l *LinkedList) Set(index, val int) bool {
	foundNode := l.Get(index)
	if foundNode == nil {
		return false
	}
	foundNode.Val = val
	return true
}

func (l *LinkedList) Insert(index, val int) bool {
	if index < 0 || index > l.Length {
		return false
	}
	switch index {
	case 0:
		l.Unshift(val)
	case l.Length - 1:
		l.Push(val)
	default:
		prevNode := l.Get(index - 1)
		node := &ListNode{Val: val}
		node.Next = prevNode.Next
		prevNode.Next = node
		l.Length++
	}
	return true
}

func (l *LinkedList) Remove(index int) bool {
	if index < 0 || index >= l.Length {
		return false
	}
	switch index {
	case 0:
		l.Shift()
	case l.Length - 1:
		l.Pop()

	default:
		prevNode := l.Get(index - 1)
		prevNode.Next = prevNode.Next.Next
		l.Length--
	}
	return true
}

func (l *LinkedList) Size() int {
	return l.Length
}
func (l *LinkedList) Contains(val int) bool {
	current := l.Head
	for current != nil {
		if current.Val == val {
			return true
		}
		current = current.Next
	}
	return false
}

func (l *LinkedList) Clear() {
	l.Head = nil
	l.Tail = nil
	l.Length = 0
}

func (l *LinkedList) List() []int {
	list := []int{}
	current := l.Head
	for current != nil {
		list = append(list, current.Val)
		current = current.Next
	}
	return list
}
