package linkedlist

import "testing"

func TestPrependAndLength(t *testing.T) {
	if got := Length(nil); got != 0 {
		t.Fatalf("Length(nil) = %d, want 0", got)
	}

	tail := &Node{Value: 2}

	head := Prepend(tail, 1)
	if head == nil {
		t.Fatal("Prepend returned nil")
	}

	if head.Value != 1 || head.Next != tail {
		t.Fatalf("Prepend returned %#v, want a new node pointing to the old head", head)
	}

	if got := Length(head); got != 2 {
		t.Fatalf("Length(head) = %d, want 2", got)
	}
}

func TestFind(t *testing.T) {
	duplicate := &Node{Value: 2}
	third := &Node{Value: 3, Next: duplicate}
	second := &Node{Value: 2, Next: third}
	first := &Node{Value: 1, Next: second}

	if got := Find(first, 2); got != second {
		t.Fatalf("Find must return the existing matching node: got %p, want %p", got, second)
	}

	if got := Find(first, 99); got != nil {
		t.Fatalf("Find(first, 99) = %#v, want nil", got)
	}

	if got := Find(nil, 1); got != nil {
		t.Fatalf("Find(nil, 1) = %#v, want nil", got)
	}
}

func TestReverse(t *testing.T) {
	third := &Node{Value: 3}
	second := &Node{Value: 2, Next: third}
	first := &Node{Value: 1, Next: second}

	head := Reverse(first)
	if head != third || head.Next != second || second.Next != first || first.Next != nil {
		t.Fatalf("Reverse did not reverse the existing nodes: head=%#v", head)
	}

	if got := Reverse(nil); got != nil {
		t.Fatalf("Reverse(nil) = %#v, want nil", got)
	}
}
