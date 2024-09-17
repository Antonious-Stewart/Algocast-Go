package linked_list

import (
	"fmt"
)

type Node struct {
	data interface{}
	next *Node
}

func NewNode(value interface{}) (*Node, error) {
	if value == nil {
		return nil, fmt.Errorf("invalid value")
	}
	node := Node{
		data: value,
		next: nil,
	}

	return &node, nil
}

type List struct {
	Head   *Node
	Tail   *Node
	Length int
}

func (l *List) NewList(value interface{}) error {
	node, err := NewNode(value)

	if err != nil {
		return err
	}

	l.Head = node
	l.Tail = l.Head
	l.Length = 1
	return nil
}

func (l *List) GetLength() int {
	return l.Length
}

func (l *List) GetHead() *Node {
	return l.Head
}

func (l *List) GetTail() *Node {
	return l.Tail
}
