package main

import "fmt"

// Node represents node use for stack
type Node struct {
	Value int
}

// Request represents printing method
func (n *Node) String() string {
	return fmt.Sprint(n.Value)
}

// NewStack represents configuring new stack
func NewStack() *Stack {
	return &Stack{}
}

// Stack represents structure for stack
type Stack struct {
	nodes []*Node
	count int
}

// Push represents put nodes in stack
func (s *Stack) Push(n *Node) {
	s.nodes = append(s.nodes[:s.count], n)
	s.count++
}

// Pop represents remove nodes in stack
func (s *Stack) Pop() *Node {
	if s.count == 0 {
		return nil
	}
	s.count--
	return s.nodes[s.count]
}

func main() {
	s := NewStack()
	s.Push(&Node{1})
	s.Push(&Node{2})
	s.Push(&Node{3})
	s.Push(&Node{4})
	s.Push(&Node{5})

	fmt.Println(s.Pop())
	fmt.Println(s.Pop())
	fmt.Println(s.Pop())
	fmt.Println(s.Pop())
	fmt.Println(s.Pop())

}
