package linkedlist

import (
	"errors"
	"fmt"
	"sort"
)

var errInvalidCellNum = errors.New("error: can't access cell number")
var errMultipleIndex = errors.New("indexes slice can't contain multiple appearances of the same index")
var errPopEmpty = errors.New("cant pop from an empty list")

type LinkedList struct {
	// holds the length, tail and head for simple pop push and to be able to validate input quickly.
	head *Node
	tail *Node
	len  int
}

type Node struct {
	// val is interface in order to hold every type of value, next and prev are held to keep a doubly linked list.
	val  interface{}
	next *Node
	prev *Node
}

// GetNext returns the node's next node
func (n *Node) GetNext() *Node {
	return n.next
}

// GetPrev returns the node's previous node
func (n *Node) GetPrev() *Node {
	return n.prev
}

// GetVal returns the Node's value
func (n *Node) GetVal() interface{} {
	return n.val
}

func NewLinkedList() LinkedList {
	return LinkedList{nil, nil, 0}
}

// PushBack push to the front of the tail of the list.
func (l *LinkedList) PushBack(value interface{}) *Node {
	var newNode = Node{value, nil, nil}
	// if head is nil then this list is empty
	if l.head == nil {
		l.head = &newNode
	} else {
		l.tail.next = &newNode
		newNode.prev = l.tail
	}

	l.tail = &newNode
	l.len++

	return &newNode
}

// PushFront push to the back of the head of the list.
func (l *LinkedList) PushFront(value interface{}) *Node {
	var newNode = Node{value, nil, nil}
	// if tail is nil then the list is empty
	if l.tail == nil {
		l.tail = &newNode
	} else {
		l.head.prev = &newNode
		newNode.next = l.head
	}

	l.head = &newNode
	l.len++

	return &newNode
}

// PopBack pop the last element.
func (l *LinkedList) PopBack() (interface{}, error) {
	if l.len == 0 {
		return 0, errPopEmpty
	}

	l.len--
	var val = l.tail.val

	if l.len == 0 {
		l.head = nil
		l.tail = nil
	} else {
		l.tail = l.tail.prev
		l.tail.next = nil
	}

	return val, nil
}

// PopFront pop the first element.
func (l *LinkedList) PopFront() (interface{}, error) {
	if l.len == 0 {
		return 0, errPopEmpty
	}

	l.len--
	var val = l.head.val

	if l.len == 0 {
		l.head = nil
		l.tail = nil
	} else {
		l.head = l.head.next
		l.head.prev = nil
	}

	return val, nil
}

// Peek this function returns the i'th value, an error is returned for illegal indexes.
func (l *LinkedList) Peek(i int) (interface{}, error) {
	if i < 0 || i >= l.len {
		return 0, fmt.Errorf("%w: %d", errInvalidCellNum, i)
	}

	n := 0
	var curNode = l.head
	// progress until you reach the i'th element of the list
	for ; n < i; n++ {
		curNode = curNode.next
	}

	return curNode.val, nil
}

// EraseIndexList this function receives an index list and removes all the nodes corresponding to these indexes in the
// linked list.
func (l *LinkedList) EraseIndexList(indexes []int) error {
	sort.Ints(indexes)

	cur := l.head
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

func (l *LinkedList) Len() int { return l.len }

func (l *LinkedList) Head() *Node { return l.head }

func (l *LinkedList) eraseNode(tmp *Node) {
	if tmp == l.head {
		l.head = tmp.next
	}

	if tmp == l.tail {
		l.tail = tmp.prev
	}

	if tmp.next != nil {
		tmp.next.prev = tmp.prev
	}

	if tmp.prev != nil {
		tmp.prev.next = tmp.next
	}

	l.len--
}
