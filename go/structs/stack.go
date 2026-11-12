package structs

import (
	"errors"
	"fmt"
)

type Stack struct {
	Head *Node
}

// Добавить в стек (на вершину)
func (s *Stack) Push(data string) {
	s.Head = &Node{Data: data, Next: s.Head}
}

// Извлечь с вершины
func (s *Stack) Pop() (string, error) {
	if s.Head == nil {
		return "", errors.New("stack is empty")
	}
	data := s.Head.Data
	s.Head = s.Head.Next
	return data, nil
}

// Посмотреть вершину (не удаляя)
func (s *Stack) Front() (string, error) {
	if s.Head == nil {
		return "", errors.New("stack is empty")
	}
	return s.Head.Data, nil
}

// Вывести на экран
func (s *Stack) Print() {
	cur := s.Head
	for cur != nil {
		fmt.Print(cur.Data, " ")
		cur = cur.Next
	}
	fmt.Println()
}