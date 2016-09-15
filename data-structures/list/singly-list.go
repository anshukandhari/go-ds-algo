package list

import (
	"errors"
	"fmt"
)

// ######################################### Structure Declarations #######################################
type SLL struct {
	Length int
	Head   *SllNode
	Tail   *SllNode
}

type SllNode struct {
	Value int
	Next  *SllNode
}

// ######################################### Creational Functions #######################################
func NewSLL() *SLL {
	l := new(SLL)
	l.Length = 0
	return l
}

func NewSllNode(value int) *SllNode {
	return &SllNode{Value: value}
}

// ######################################### Util Functions #######################################

func (l *SLL) IsEmpty() bool {
	return l.Length == 0
}

func (l *SLL) PrintSLL() {
	for node := l.Head; node != nil; node = node.Next {
		fmt.Print(node.Value, " ")
	}
	fmt.Println()
	fmt.Println("Head, Tail & Length are :", l.Head.Value, l.Tail.Value, l.Length)
}

func (l *SLL) GetAtIndex(index int) (*SllNode, error) {
	if index > l.Length {
		return nil, errors.New("Index out of range")
	}

	node := l.Head
	for i := 0; i < index; i++ {
		node = node.Next
	}

	return node, nil
}

func (l *SLL) Clear() {
	l.Length = 0
	l.Head = nil
	l.Tail = nil
}

// ######################################### SLL Basic Operations #######################################

func (l *SLL) Prepend(value int) {
	node := NewSllNode(value)
	if l.Length == 0 {
		l.Head = node
		l.Tail = l.Head
	} else {
		formerHead := l.Head

		node.Next = formerHead
		l.Head = node
	}

	l.Length++
}

func (l *SLL) Append(value int) {
	node := NewSllNode(value)

	if l.Length == 0 {
		l.Head = node
		l.Tail = l.Head
	} else {
		formerTail := l.Tail
		formerTail.Next = node

		l.Tail = node
	}

	l.Length++
}

func (l *SLL) Add(value int, index int) error {
	if index > l.Length || index < 0 {
		return errors.New("Index out of range")
	}

	if l.Length == 0 || index == 0 {
		l.Prepend(value)
		return nil
	}

	if l.Length == index {
		l.Append(value)
		return nil
	}

	prevSllNode, err := l.GetAtIndex(index - 1)
	if err != nil {
		return err
	}

	node := NewSllNode(value)
	node.Next = prevSllNode.Next
	prevSllNode.Next = node

	l.Length++
	return nil
}

func (l *SLL) Remove(value int) error {
	if l.Length == 0 {
		return errors.New("Empty SLL")
	}

	if l.Head.Value == value {
		l.Head = l.Head.Next
		l.Length--
		return nil
	}

	found := 0
	var prev *SllNode
	for n := l.Head; n != nil; prev, n = n, n.Next {
		if n.Value == value && found == 0 {
			prev.Next = n.Next
			l.Length--
			found++
			break
		}
	}

	if found == 0 {
		return errors.New("SllNode not found")
	}

	return nil
}
