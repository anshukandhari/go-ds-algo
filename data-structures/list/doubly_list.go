package list

import (
	"errors"
	"fmt"
)

// ######################################### Structure Declarations #######################################
type DLL struct {
	Length int
	Head   *DllNode
	Tail   *DllNode
}

type DllNode struct {
	Value int
	Prev  *DllNode
	Next  *DllNode
}

// ######################################### Creational Functions #######################################
func NewDLL() *DLL {
	l := new(DLL)
	l.Length = 0
	return l
}

func NewDllNode(value int) *DllNode {
	return &DllNode{Value: value}
}

// #################################### DLL Useful Functions (utilities) #######################################

func (l *DLL) IsEmpty() bool {
	return l.Length == 0
}

func (l *DLL) PrintDLL() {
	for node := l.Head; node != nil; node = node.Next {
		fmt.Print(node.Value, " ")
	}
	fmt.Println()
	// fmt.Println("Head, Tail & Length are :", l.Head.Value, l.Tail.Value, l.Length)
}

func (l *DLL) GetAtIndex(index int) (*DllNode, error) {
	if index > l.Length {
		return nil, errors.New("Index out of range")
	}

	node := l.Head
	for i := 0; i < index; i++ {
		node = node.Next
	}

	return node, nil
}

func (l *DLL) Clear() {
	l.Length = 0
	l.Head = nil
	l.Tail = nil
}

// ######################################### DLL Basic Operations #######################################

func (l *DLL) Prepend(value int) {
	node := NewDllNode(value)
	if l.Length == 0 {
		l.Head = node
		l.Tail = l.Head
	} else {
		formerHead := l.Head
		formerHead.Prev = node

		node.Next = formerHead
		l.Head = node
	}

	l.Length++
}

func (l *DLL) Append(value int) {
	node := NewDllNode(value)

	if l.Length == 0 {
		l.Head = node
		l.Tail = l.Head
	} else {
		formerTail := l.Tail
		formerTail.Next = node

		node.Prev = formerTail
		l.Tail = node
	}

	l.Length++
}

func (l *DLL) DelAtEnd() {
	l.Tail = l.Tail.Prev
	l.Tail.Next = nil
	l.Length--
}

func (l *DLL) DelAtStart() {
	l.Head = l.Head.Next
	l.Head.Prev = nil
	l.Length--
}

func (l *DLL) Add(value int, index int) error {
	if index > l.Length || index < 0 {
		return errors.New("Index out of range")
	}

	if l.Length == 0 || index == 0 {
		l.Prepend(value)
		return nil
	}

	if l.Length-1 == index {
		l.Append(value)
		return nil
	}

	nextDllNode, err := l.GetAtIndex(index)
	if err != nil {
		return err
	}

	node := NewDllNode(value)
	prevDllNode := nextDllNode.Prev

	prevDllNode.Next = node
	node.Prev = prevDllNode

	nextDllNode.Prev = node
	node.Next = nextDllNode

	l.Length++

	return nil
}

func (l *DLL) Remove(value int) error {
	if l.Length == 0 {
		return errors.New("Empty list")
	}

	if l.Head.Value == value {
		l.DelAtStart()
		return nil
	}

	if l.Tail.Value == value {
		l.DelAtEnd()
		return nil
	}

	found := 0
	for n := l.Head; n != nil; n = n.Next {
		if n.Value == value && found == 0 {
			n.Next.Prev, n.Prev.Next = n.Prev, n.Next
			l.Length--
			found++
			break
		}
	}

	if found == 0 {
		return errors.New("DllNode not found")
	}

	return nil
}

// ######################################### DLL Interview Questions #######################################

/*
	Reversing a DLL in iterative manner
*/
func (l *DLL) ReverseIterative() {
	var temp = l.Head
	var p, q, r *DllNode
	for r = l.Head; r != nil; {
		p = q
		q = r
		r = r.Next
		q.Next, q.Prev = p, r
	}
	l.Head, l.Tail = q, temp
}

/*
	Reversing a DLL using recursion
*/
func (l *DLL) RecursiveReverse(n *DllNode) *DllNode {
	if l.Length <= 1 {
		return l.Head
	}

	if n.Next == nil {
		l.Head = n
		return n
	}

	curr := l.RecursiveReverse(n.Next)
	curr.Next = n
	if n.Prev == nil {
		n.Next = nil
		l.Tail = n
	}
	n.Prev = curr
	return n
}

/*
	Question Explanation for SLL: http://www.geeksforgeeks.org/reverse-a-list-in-groups-of-given-size/

	func ReverseInGroups
	1. Save the start node
	2. Go to the kth DllNode in iteration, reverse every pair in iteration
	3. Return Kth node
		a. If the start was the head node, then make kth DllNode, the new Head
		b. If the iteration reached the end of the list then update Tail & make it point to the saved start
	4. start.next = ReverseInGroups(K+1)
*/
func (l *DLL) ReverseInGroups(start *DllNode, k int) *DllNode {
	if l.Length <= 1 || k <= 1 {
		return l.Head
	}

	if start == nil {
		return start
	}

	var p, q, r *DllNode
	i := 0
	for r = start; i < k; i++ {
		p = q
		q = r
		r = r.Next
		q.Next, q.Prev = p, r
		if r == nil {
			l.Tail = start
			break
		}
	}

	if start == l.Head {
		l.Head = q
	}

	start.Next = l.ReverseInGroups(r, k)
	return q
}
