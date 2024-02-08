package tree_test

import "testing"

type Node[T any] struct {
	left  *Node[T]
	right *Node[T]
	value T
}

func TestTree(t *testing.T) {
	root := &Node[any]{nil, nil, 1}
	root.left = &Node[any]{nil, nil, 2}
	root.right = &Node[any]{nil, nil, 3}
	root.left.left = &Node[any]{nil, nil, 4}
	root.left.right = &Node[any]{nil, nil, 5}
	root.right.left = &Node[any]{nil, nil, 6}
	root.right.right = &Node[any]{nil, nil, 7}

	t.Log(root)

}
