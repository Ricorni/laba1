package structs

import (
	"errors"
	"fmt"
)

type Queue struct {
	Head *Node
	Tail *Node
}

// Добавить в очередь (в конец)
func (q *Queue) Push(data string) {
	newNode := &Node{Data: data}
	if q.Tail != nil {
		q.Tail.Next = newNode
	} else {
		q.Head = newNode
	}
	q.Tail = newNode
}

// Извлечь из начала очереди
func (q *Queue) Pop() (string, error) {
	if q.Head == nil {
		return "", errors.New("queue is empty")
	}
	data := q.Head.Data
	q.Head = q.Head.Next
	if q.Head == nil {
		q.Tail = nil
	}
	return data, nil
}

// Посмотреть начало (не удаляя)
func (q *Queue) Front() (string, error) {
	if q.Head == nil {
		return "", errors.New("queue is empty")
	}
	return q.Head.Data, nil
}

// Посмотреть конец (не удаляя)
func (q *Queue) Back() (string, error) {
	if q.Tail == nil {
		return "", errors.New("queue is empty")
	}
	return q.Tail.Data, nil
}

// Вывести на экран
func (q *Queue) Print() {
	cur := q.Head
	for cur != nil {
		fmt.Print(cur.Data, " ")
		cur = cur.Next
	}
	fmt.Println()
}