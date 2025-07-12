package linkedlist

import (
	"testing"
)

// Table-driven test example (most common Go pattern)
func TestPushTableDriven(t *testing.T) {
	tests := []struct {
		name     string
		values   []int
		wantLen  int
		wantHead int
		wantTail int
	}{
		{"single element", []int{1}, 1, 1, 1},
		{"multiple elements", []int{1, 2, 3}, 3, 1, 3},
		{"empty then push", []int{}, 0, 0, 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ll := &LinkedList{}
			for _, val := range tt.values {
				ll.Push(val)
			}

			if len(tt.values) > 0 {
				if ll.Length != tt.wantLen || ll.Head.Val != tt.wantHead || ll.Tail.Val != tt.wantTail {
					t.Errorf("Push() = len:%d, head:%d, tail:%d, want len:%d, head:%d, tail:%d",
						ll.Length, ll.Head.Val, ll.Tail.Val, tt.wantLen, tt.wantHead, tt.wantTail)
				}
			}
		})
	}
}

func TestPushAndPop(t *testing.T) {
	ll := &LinkedList{}

	// Test push
	ll.Push(1)
	ll.Push(2)
	ll.Push(3)

	if ll.Length != 3 || ll.Head.Val != 1 || ll.Tail.Val != 3 {
		t.Errorf("Push failed: length=%d, head=%d, tail=%d", ll.Length, ll.Head.Val, ll.Tail.Val)
	}

	// Test pop
	node := ll.Pop()
	if node.Val != 3 || ll.Length != 2 || ll.Tail.Val != 2 {
		t.Errorf("Pop failed: popped=%d, length=%d, tail=%d", node.Val, ll.Length, ll.Tail.Val)
	}

	// Test pop empty
	ll.Pop()
	ll.Pop()
	if ll.Pop() != nil || ll.Length != 0 {
		t.Error("Pop from empty list should return nil")
	}
}

func TestShiftAndUnshift(t *testing.T) {
	ll := &LinkedList{}

	// Test unshift
	ll.Unshift(1)
	ll.Unshift(2)

	if ll.Length != 2 || ll.Head.Val != 2 || ll.Tail.Val != 1 {
		t.Errorf("Unshift failed: length=%d, head=%d, tail=%d", ll.Length, ll.Head.Val, ll.Tail.Val)
	}

	// Test shift
	node := ll.Shift()
	if node.Val != 2 || ll.Length != 1 || ll.Head.Val != 1 {
		t.Errorf("Shift failed: shifted=%d, length=%d, head=%d", node.Val, ll.Length, ll.Head.Val)
	}

	// Test shift empty
	ll.Shift()
	if ll.Shift() != nil || ll.Length != 0 {
		t.Error("Shift from empty list should return nil")
	}
}

func TestGetAndSet(t *testing.T) {
	ll := &LinkedList{}
	ll.Push(1)
	ll.Push(2)
	ll.Push(3)

	// Test get
	if ll.Get(0).Val != 1 || ll.Get(1).Val != 2 || ll.Get(2).Val != 3 {
		t.Error("Get failed to retrieve correct values")
	}

	if ll.Get(-1) != nil || ll.Get(3) != nil {
		t.Error("Get should return nil for invalid indices")
	}

	// Test set
	if !ll.Set(1, 20) || ll.Get(1).Val != 20 {
		t.Error("Set failed to update value")
	}

	if ll.Set(-1, 10) || ll.Set(3, 10) {
		t.Error("Set should return false for invalid indices")
	}
}

func TestInsertAndRemove(t *testing.T) {
	ll := &LinkedList{}

	// Test insert
	ll.Insert(0, 2) // [2]
	ll.Insert(0, 1) // [1, 2]
	ll.Insert(2, 3) // [1, 2, 3]
	ll.Insert(1, 5) // [1, 5, 2, 3]

	expected := []int{1, 5, 2, 3}
	actual := ll.List()
	if len(actual) != len(expected) {
		t.Errorf("Insert failed: expected length %d, got %d", len(expected), len(actual))
	}

	for i, val := range expected {
		if actual[i] != val {
			t.Errorf("Insert failed: expected %v, got %v", expected, actual)
			break
		}
	}

	// Test remove
	ll.Remove(1) // Remove 5: [1, 2, 3]
	ll.Remove(2) // Remove 3: [1, 2]

	if ll.Length != 2 || ll.Get(0).Val != 1 || ll.Get(1).Val != 2 {
		t.Error("Remove failed")
	}

	// Test invalid operations
	if ll.Insert(-1, 10) || ll.Insert(10, 10) || ll.Remove(-1) || ll.Remove(10) {
		t.Error("Invalid insert/remove operations should return false")
	}
}

func TestUtilityMethods(t *testing.T) {
	ll := &LinkedList{}

	// Test empty list
	if ll.Size() != 0 || ll.Contains(1) || len(ll.List()) != 0 {
		t.Error("Empty list methods failed")
	}

	// Add elements
	ll.Push(1)
	ll.Push(2)
	ll.Push(3)

	// Test size
	if ll.Size() != 3 {
		t.Errorf("Size failed: expected 3, got %d", ll.Size())
	}

	// Test contains
	if !ll.Contains(2) || ll.Contains(4) {
		t.Error("Contains failed")
	}

	// Test list
	list := ll.List()
	expected := []int{1, 2, 3}
	if len(list) != len(expected) {
		t.Errorf("List failed: expected length %d, got %d", len(expected), len(list))
	}

	for i, val := range expected {
		if list[i] != val {
			t.Errorf("List failed: expected %v, got %v", expected, list)
			break
		}
	}

	// Test clear
	ll.Clear()
	if ll.Size() != 0 || ll.Head != nil || ll.Tail != nil {
		t.Error("Clear failed")
	}
}
