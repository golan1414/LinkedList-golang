package linkedlist

import (
	"errors"
	"fmt"
	"sort"
)

var (
	errInvalidCellNum = errors.New("error: can't access cell number")
	errMultipleIndex  = errors.New("indexes slice can't contain multiple appearances of the same index")
	errPopEmpty       = errors.New("cant pop from an empty list")
)

type LinkedList struct {
	// holds the length, tail and head for simple pop push and to be able to validate input quickly.
	root *Node
	len  int
}

type Node struct {
	// val is interface in order to hold every type of value, next and prev are held to keep a doubly linked list.
	val  interface{}
	next *Node
	prev *Node
	list *LinkedList
}

// GetNext returns the node's next node.
func (n *Node) GetNext() *Node {
	if n.next != n.list.root {
		return n.next
	}
	return nil
}

// GetPrev returns the node's previous node.
func (n *Node) GetPrev() *Node {
	if n.prev != n.list.root {
		return n.prev
	}
	return nil
}

// GetVal returns the Node's value.
func (n *Node) GetVal() interface{} {
	return n.val
}

// NewLinkedList returns a new LinkedList instance.
func NewLinkedList() LinkedList {
	l := LinkedList{&Node{nil, nil, nil, nil}, 0}
	l.root.list = &l
	return l
}

func (l *LinkedList) insertAt(at, toInsert *Node) *Node {
	l.len++
	toInsert.next = at.next
	at.next = toInsert
	toInsert.prev = at
	toInsert.next.prev = toInsert
	return toInsert
}

// PushBack push to the front of the tail of the list.
func (l *LinkedList) PushBack(value interface{}) *Node {
	return l.insertAt(l.root.next, &Node{value, nil, nil, l})
}

// PushFront push to the back of the head of the list.
func (l *LinkedList) PushFront(value interface{}) *Node {
	return l.insertAt(l.root.prev, &Node{value, nil, nil, l})
}

// PopBack pop the last element.
func (l *LinkedList) PopBack() (interface{}, error) {
	if l.len == 0 {
		return 0, errPopEmpty
	}

	val := l.root.prev.val

	l.eraseNode(l.root.prev)

	return val, nil
}

// PopFront pop the first element.
func (l *LinkedList) PopFront() (interface{}, error) {
	if l.len == 0 {
		return 0, errPopEmpty
	}

	val := l.root.next.val

	l.eraseNode(l.root.next)

	return val, nil
}

// Peek this function returns the i'th value, an error is returned for illegal indexes.
func (l *LinkedList) Peek(i int) (interface{}, error) {
	if i < 0 || i >= l.len {
		return 0, fmt.Errorf("%w: %d", errInvalidCellNum, i)
	}

	curNode := l.root.next
	// progress until you reach the i'th element of the list
	for n := 0; n < i; n++ {
		curNode = curNode.next
	}

	return curNode.val, nil
}

// EraseIndexList this function receives an index list and removes all the nodes corresponding to these indexes in the
// linked list.
func (l *LinkedList) EraseIndexList(indexes []int) error {
	sort.Ints(indexes)

	cur := l.root.next
	// j is the index of the current
	j := 0

	for offset, i := range indexes {
		// every time we delete a node all the nodes from before have their index decrease by one so we need to subtract
		// the offset
		i -= offset
		if i < 0 || i >= l.len {
			return fmt.Errorf("%w: %d", errInvalidCellNum, i)
		}

		if j > i {
			return errMultipleIndex
		}

		for ; j < i; j++ {
			cur = cur.next
		}
		// we dont want to lose our current location so we stop one before.
		tmp := cur
		cur = cur.next

		l.eraseNode(tmp)
	}

	return nil
}

// Len return the list length.
func (l *LinkedList) Len() int { return l.len }

// Head return the list head.
func (l *LinkedList) Head() *Node { return l.root.next }

// Tail return the list tail.
func (l *LinkedList) Tail() *Node { return l.root.prev }

// erase a single node from the list.
func (l *LinkedList) eraseNode(tmp *Node) {
	tmp.next.prev = tmp.prev
	tmp.prev.next = tmp.next
	tmp = nil

	l.len--
}
