//go:build !solution

package treetransform

// Map строит глубокую копию дерева, преобразуя значения узлов.
func Map(root *Node, transform Transform) (*Node, error) {
	_, _ = root, transform

	return nil, nil
}
