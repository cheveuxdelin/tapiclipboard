package main

import "fmt"

type node struct {
	value string
	next  *node
}

type Stack struct {
	head      *node
	size      int
	sizeLimit int
}

func (s *Stack) push(str string) {
	var newNode *node = &node{value: str}

	if s.size == 0 {
		s.head = newNode
	} else {
		newNode.next = s.head
		s.head = newNode
	}

	s.size++
}

func (s *Stack) pop() (rtn string) {
	if s.size > 0 {
		var poppedNode *node = s.head
		s.head = s.head.next
		s.size--
		rtn = poppedNode.value
	}
	return
}

func (s *Stack) front() (rtn string) {
	if s.size > 0 {
		var poppedNode *node = s.head
		rtn = poppedNode.value
	}
	return
}

func (s *Stack) clear() {
	s.head = nil
	s.size = 0
}

func (s *Stack) print() (rtn string) {
	var current *node = s.head
	var index int = 0
	for current != nil {
		rtn += fmt.Sprintf("%d: %s\n", index, current.value)
		index++
		current = current.next
	}
	return
}
