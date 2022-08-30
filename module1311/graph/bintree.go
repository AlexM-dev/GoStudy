package graph

import (
	"fmt"
)

type node struct {
	value int
	right *node
	left  *node
}

type tree struct {
	root *node
}

func NewTree(v int) *tree {
	return &tree{root: &node{value: v, right: nil, left: nil}}
}

func (n *node) add(value int) {

	if value < n.value {
		if n.left == nil {
			n.left = &node{value: value, left: nil, right: nil}
		} else {
			n.left.add(value)
		}
	} else {
		if n.right == nil {
			n.right = &node{value: value, left: nil, right: nil}
		} else {
			n.right.add(value)
		}
	}
}

func (t *tree) Add(value int) {
	if t.root == nil {
		t.root = &node{value: value, right: nil, left: nil}
	} else {
		t.root.add(value)
	}
}

func (n *node) find(value int) bool {
	if n == nil {
		return false
	}
	switch {
	case value == n.value:
		return true
	case value < n.value:
		return n.left.find(value)
	default:
		return n.right.find(value)
	}
}

func (t *tree) Find(value int) bool {
	if t.root == nil {
		return false
	}
	return t.root.find(value)
}

func (n *node) findMax(parent *node) (*node, *node) {
	if n == nil {
		return nil, parent
	}
	if n.right == nil {
		return n, parent
	}
	return n.right.findMax(n)
}

func (n *node) replaceNode(parent, replacement *node) {
	if n == nil {
		return
	}

	if n == parent.left {
		parent.left = replacement
	} else {
		parent.right = replacement
	}
}

func (n *node) mostLeftValue() int {
	if n.left == nil {
		return n.value
	}
	return n.right.mostLeftValue()
}

func (n *node) delete(parent *node, v int) error {
	switch {
	case n.value == v:
		if n.left != nil && n.right != nil {
			n.value = n.right.mostLeftValue()
			return n.right.delete(n, n.value)
		}
		if parent.left == n {
			if n.left != nil {
				parent.left = n.left
			} else {
				parent.left = n.right
			}
		} else if parent.right == n {
			if n.left != nil {
				parent.right = n.left
			} else {
				parent.right = n.right
			}
		}
		return nil
	case n.value > v:
		if n.left == nil {
			return fmt.Errorf("no element %d", v)
		}
		return n.left.delete(n, v)
	case n.value < v:
		if n.right == nil {
			return fmt.Errorf("no element %d", v)
		}
		return n.right.delete(n, v)
	}
	return fmt.Errorf("?")
}

func (t *tree) Delete(n int) error {
	if t.root == nil {
		return fmt.Errorf("empty tree")
	}

	if t.root.value == n {
		tempRoot := &node{value: 0, left: nil, right: nil}
		tempRoot.left = t.root
		r := t.root.delete(tempRoot, n)
		t.root = tempRoot.left
		return r
	}
	er1, er2 := t.root.left.delete(t.root, n), t.root.right.delete(t.root, n)
	if er1 == nil || er2 == nil {
		return nil
	} else {
		if er1 == nil {
			return er2
		}
		return er1
	}
}

func (n *node) print() {
	if n != nil {
		n.left.print()
		fmt.Println(n.value)
		n.right.print()
	}
}

func (t *tree) Print() {
	t.root.print()
}
