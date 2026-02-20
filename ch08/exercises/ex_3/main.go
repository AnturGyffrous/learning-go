package main

import "fmt"

type LinkedList[T comparable] struct {
	Root *Node[T]
}

type Node[T comparable] struct {
	Position int
	Value    T
	Next     *Node[T]
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

func main() {
	ll := LinkedList[int]{}
	ll.Add(2)
	ll.Add(1)
	ll.Add(3)
	ll.Add(6)
	fmt.Println(ll.Root.Value)
	fmt.Println(ll.Root.Next.Value)
	fmt.Println(ll.Root.Next.Next.Value)
	fmt.Println(ll.Root.Next.Next.Next.Value)
}
