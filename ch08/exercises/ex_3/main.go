package main

import (
	"cmp"
	"errors"
	"fmt"
)

type OrderableFunc[T any] func(t1, t2 T) int

type LinkedList[T any] struct {
	Compare OrderableFunc[T]
	Root    *Node[T]
}

type Node[T any] struct {
	Position int
	Value    T
	Next     *Node[T]
}

func NewLinkedList[T any](compare OrderableFunc[T]) *LinkedList[T] {
	return &LinkedList[T]{Compare: compare}
}

func (ll *LinkedList[T]) Add(value T) {
	ll.Root = ll.Root.Add(value)
}

func (node *Node[T]) Add(value T) *Node[T] {
	if node == nil {
		return &Node[T]{
			Position: 0,
			Value:    value,
		}
	}

	if node.Next == nil {
		node.Next = &Node[T]{
			Position: node.Position + 1,
			Value:    value,
		}
	} else {
		node.Next.Add(value)
	}

	return node
}

func (ll *LinkedList[T]) Insert(value T, index int) (bool, error) {
	if ll.Root == nil {
		return false, errors.New("list is empty")
	}

	root, ok, err := ll.Root.Insert(ll.Compare, value, index)
}

func (node *Node[T]) Insert(compare OrderableFunc[T], value T, index int) (*Node[T], bool, error) {
	if node == nil {
		return nil, false, errors.New("index out of range")
	}

	if node.Position < index {
		next, ok, err := node.Next.Insert(compare, value, index)
		if !ok {
			return nil, false, err
		}
		node.Next = next
		return node, true, nil
	}
	if node.Position == index {

	}
}

func main() {
	ll := NewLinkedList[int](cmp.Compare)
	ll.Add(2)
	ll.Add(1)
	ll.Add(3)
	ll.Add(6)
	fmt.Println(ll.Root.Value)
	fmt.Println(ll.Root.Next.Value)
	fmt.Println(ll.Root.Next.Next.Value)
	fmt.Println(ll.Root.Next.Next.Next.Value)
}
