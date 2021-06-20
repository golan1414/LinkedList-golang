package linkedList

import (
	"errors"
	"fmt"
	"sort"
)

type LinkedList struct {
	head *Node
	tail *Node
	len int
}

type Node struct {
	val interface{}
	next *Node
	prev *Node
}

func (l *LinkedList) push(value interface{}) *Node {
	var newNode = Node{value, nil, nil}
	if l.head == nil {
		l.head = &newNode
	} else {
		l.tail.next = &newNode
		newNode.prev = l.tail
	}
	l.tail = &newNode
	l.len = l.len + 1
	return &newNode
}

func (l *LinkedList) pop() (interface{}, error) {
	if l.len == 0 {
		return 0, errors.New("cant pop from an empty list")
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

func (l *LinkedList) peek(i int) (interface{}, error) {
	if i < 0 || i >= l.len {
		return 0, errors.New(fmt.Sprintf("Cant access cell number %d", i))
	}
	n := 0
	var curNode = l.head
	for ; n < i; n++ {
		curNode = curNode.next
	}
	return curNode.val, nil
}

func (l *LinkedList) eraseIndexList(indexes []int) error {
	sort.Ints(indexes)
	cur := l.head
	// j is the index of the current
	j := 0
	for offset, i := range indexes {
		// every time we delete a node all the nodes from before have their index decrease by one so we need to subtract the offset
		i = i - offset
		if i < 0 || i >= l.len {
			return errors.New(fmt.Sprintf("cell number %d does not exist", i))
		}
		if j > i {
			return errors.New("indexes slice can't contain multiple appearances of the same index")
		}
		for ; j < i; j++ {
			cur = cur.next
		}
		// we dont want to lose our current location so we stop one before
		tmp := cur
		cur = cur.next
		l.eraseNode(tmp)
		l.len--
	}
	return nil
}

func (l *LinkedList) eraseNode(node *Node) {
	if node == l.head {
		l.head = node.next
	}
	if node == l.tail {
		l.tail = node.prev
	}
	if node.next != nil {
		node.next.prev = node.prev
	}
	if node.prev != nil {
		node.prev.next = node.next
	}
}