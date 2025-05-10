package main

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v homework_test.go

type OrderedMap struct {
	root *TreeNode
	size int
}

type TreeNode struct {
	key   int
	value int
	left  *TreeNode
	right *TreeNode
}

func newTreeNode(key, value int, left *TreeNode, right *TreeNode) *TreeNode {
	return &TreeNode{
		key:   key,
		value: value,
		left:  left,
		right: right,
	}

}

func NewOrderedMap() OrderedMap {
	return OrderedMap{
		root: nil,
		size: 0,
	}
}

func (m *OrderedMap) Insert(key, value int) {
	if m.root != nil {
		var curNode *TreeNode
		curNode = m.root
		for {
			if curNode.key < key {
				if curNode.right != nil {
					curNode = curNode.right
				} else {
					curNode.right = newTreeNode(key, value, nil, nil)
					m.size++
					return
				}
			} else if curNode.key > key {
				if curNode.left != nil {
					curNode = curNode.left
				} else {
					curNode.left = newTreeNode(key, value, nil, nil)
					m.size++
					return
				}
			} else {
				curNode.value = value
				return
			}
		}
	} else {
		m.root = newTreeNode(key, value, nil, nil)
		m.size++
	}
}

func (m *OrderedMap) Erase(key int) {
	var curNode *TreeNode
	curNode = m.root
	if curNode != nil {
		if curNode.key == key {
			if curNode.left == nil && curNode.right == nil {
				m.root = nil
				m.size = 0
			}
		}
	}
	prevNewCurNodeMain := curNode
	for {
		if curNode.key == key {
			if curNode.left == nil && curNode.right == nil {
				if prevNewCurNodeMain.left == curNode {
					prevNewCurNodeMain.left = nil
					m.size--
				} else if prevNewCurNodeMain.right == curNode {
					prevNewCurNodeMain.right = nil
					m.size--
				}
				return
			} else if curNode.left == nil {
				curNode.value = curNode.right.value
				curNode.key = curNode.right.key
				curNode.left = curNode.right.left
				curNode.right = curNode.right.right
				m.size--
				return
			} else if curNode.right == nil {
				curNode.value = curNode.left.value
				curNode.key = curNode.left.key
				curNode.right = curNode.left.right
				curNode.left = curNode.left.left
				m.size--
				return
			} else {
				var newCurNode = curNode.right
				var prevNewCurNode = curNode
				for {
					if newCurNode.left != nil {
						prevNewCurNode = newCurNode
						newCurNode = newCurNode.left
					} else {
						break
					}
				}
				curNode.key = newCurNode.key
				curNode.value = newCurNode.value
				prevNewCurNode.left = nil
				m.size--
				return
			}
		} else if curNode.key < key {
			if curNode.right != nil {
				prevNewCurNodeMain = curNode
				curNode = curNode.right
			} else {
				// key doesn't exist
				return
			}
		} else if curNode.key > key {
			if curNode.left != nil {
				prevNewCurNodeMain = curNode
				curNode = curNode.left
			} else {
				// key doesn't exist
				return
			}
		}
	}
}

func (m *OrderedMap) Contains(key int) bool {
	if m.root != nil {
		var curNode *TreeNode
		curNode = m.root
		for {
			if curNode.key < key {
				if curNode.right != nil {
					curNode = curNode.right
				} else {
					return false
				}
			} else if curNode.key > key {
				if curNode.left != nil {
					curNode = curNode.left
				} else {
					return false
				}
			} else {
				return true
			}
		}
	}
	return false
}

func (m *OrderedMap) Size() int {
	return m.size
}

func (m *OrderedMap) ForEach(action func(int, int)) {
	if m.root != nil {
		callFunction(m.root, action)
	}
}

func callFunction(node *TreeNode, action func(int, int)) {
	if node.left != nil {
		callFunction(node.left, action)
	}
	action(node.value, 0)
	if node.right != nil {
		callFunction(node.right, action)
	}
}

func TestCircularQueue(t *testing.T) {
	data := NewOrderedMap()
	assert.Zero(t, data.Size())

	data.Insert(10, 10)
	data.Insert(5, 5)
	data.Insert(15, 15)
	data.Insert(2, 2)
	data.Insert(4, 4)
	data.Insert(12, 12)
	data.Insert(14, 14)

	assert.Equal(t, 7, data.Size())
	assert.True(t, data.Contains(4))
	assert.True(t, data.Contains(12))
	assert.False(t, data.Contains(3))
	assert.False(t, data.Contains(13))

	var keys []int
	expectedKeys := []int{2, 4, 5, 10, 12, 14, 15}
	data.ForEach(func(key, _ int) {
		keys = append(keys, key)
	})

	assert.True(t, reflect.DeepEqual(expectedKeys, keys))

	data.Erase(15)
	data.Erase(14)
	data.Erase(2)

	assert.Equal(t, 4, data.Size())
	assert.True(t, data.Contains(4))
	assert.True(t, data.Contains(12))
	assert.False(t, data.Contains(2))
	assert.False(t, data.Contains(14))

	keys = nil
	expectedKeys = []int{4, 5, 10, 12}
	data.ForEach(func(key, _ int) {
		keys = append(keys, key)
	})

	assert.True(t, reflect.DeepEqual(expectedKeys, keys))
}
