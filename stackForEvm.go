package main

import "fmt"

// FunctionNode represents node use for stack
type FunctionNode struct {
	FunctionName    string
	ContractName    string
	ContractAddress string
	FuntionOrder    string
	FuntionID       string
}

// Request represents printing method
func (n *FunctionNode) String() string {
	return fmt.Sprint(n.FunctionName, n.ContractName, n.ContractAddress, n.FuntionOrder, n.FuntionID)
}

// NewCallStack represents configuring new stack
func NewCallStack() *CallStack {
	return &CallStack{}
}

// CallStack represents structure for stack
type CallStack struct {
	functions []*FunctionNode
	count     int
}

// Push represents put nodes in stack
func (s *CallStack) Push(n *FunctionNode) {
	s.functions = append(s.functions[:s.count], n)
	s.count++
}

// Pop represents remove nodes in stack
func (s *CallStack) Pop() *FunctionNode {
	if s.count == 0 {
		fmt.Println("nothing to pop")
		return nil
	}
	s.count--
	return s.functions[s.count]
}
