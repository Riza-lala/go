//go:build !solution

package linkedlist

// Prepend добавляет новый узел в начало списка.
func Prepend(head *Node, value int) *Node {
	_, _ = head, value

	return nil
}

// Length возвращает число узлов списка.
func Length(head *Node) int {
	_ = head

	return 0
}

// Find ищет первый узел со значением value.
func Find(head *Node, value int) *Node {
	_, _ = head, value

	return nil
}

// Reverse разворачивает список, меняя связи между существующими узлами.
func Reverse(head *Node) *Node {
	_ = head

	return nil
}
