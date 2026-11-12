package structs

import (
	"errors"
	"fmt"
)

type ForwardList struct {
	Head *Node
	Tail *Node
}

// Добавить в конец
func (f *ForwardList) PushBack(data string) {
	newNode := &Node{Data: data}
	if f.Head == nil {
		f.Head = newNode
		f.Tail = newNode
	} else {
		f.Tail.Next = newNode
		f.Tail = newNode
	}
}

// Добавить в начало
func (f *ForwardList) PushFront(data string) {
	f.Head = &Node{Data: data, Next: f.Head}
	if f.Tail == nil {
		f.Tail = f.Head
	}
}

// Поиск узла по значению (возвращает узел или ошибку)
func (f *ForwardList) find(data string) (*Node, error) {
	cur := f.Head
	for cur != nil {
		if cur.Data == data {
			return cur, nil
		}
		cur = cur.Next
	}
	return nil, errors.New("position element not found")
}

// Вставить после элемента
func (f *ForwardList) InsertAfter(data string, target string) error {
	node, err := f.find(target)
	if err != nil {
		return err
	}

	newNode := &Node{Data: data, Next: node.Next}
	node.Next = newNode

	if node == f.Tail {
		f.Tail = newNode
	}
	return nil
}

// Вставить до элемента
func (f *ForwardList) InsertBefore(data string, target string) error {
	if f.Head == nil {
		return errors.New("list is empty")
	}
	if f.Head.Data == target {
		f.PushFront(data)
		return nil
	}

	cur := f.Head
	for cur.Next != nil && cur.Next.Data != target {
		cur = cur.Next
	}

	if cur.Next == nil {
		return errors.New("position element not found")
	}

	newNode := &Node{Data: data, Next: cur.Next}
	cur.Next = newNode
	return nil
}

// Удалить из начала
func (f *ForwardList) PopFront() error {
	if f.Head == nil {
		return errors.New("list is empty")
	}
	f.Head = f.Head.Next
	if f.Head == nil {
		f.Tail = nil
	}
	return nil
}

// Удалить из конца
func (f *ForwardList) PopBack() error {
	if f.Head == nil {
		return errors.New("list is empty")
	}
	if f.Head == f.Tail {
		f.Head = nil
		f.Tail = nil
		return nil
	}

	cur := f.Head
	for cur.Next != f.Tail {
		cur = cur.Next
	}
	cur.Next = nil
	f.Tail = cur
	return nil
}

// Удалить после элемента
func (f *ForwardList) RemoveAfter(target string) error {
	node, err := f.find(target)
	if err != nil {
		return err
	}
	if node.Next == nil {
		return errors.New("no node after this one")
	}

	if node.Next == f.Tail {
		f.Tail = node
	}
	node.Next = node.Next.Next
	return nil
}

// Удалить до элемента
func (f *ForwardList) RemoveBefore(target string) error {
	if f.Head == nil {
		return errors.New("list is empty")
	}
	if f.Head.Data == target {
		return errors.New("no node before this one")
	}
	if f.Head.Next.Data == target {
		f.Head = f.Head.Next
		return nil
	}

	cur := f.Head
	for cur.Next.Next != nil && cur.Next.Next.Data != target {
		cur = cur.Next
	}
	if cur.Next.Next == nil {
		return errors.New("position element not found")
	}

	cur.Next = cur.Next.Next
	return nil
}

// Удалить по значению
func (f *ForwardList) DeleteValue(data string) error {
	if f.Head == nil {
		return errors.New("list is empty")
	}
	if f.Head.Data == data {
		return f.PopFront()
	}

	cur := f.Head
	for cur.Next != nil && cur.Next.Data != data {
		cur = cur.Next
	}
	if cur.Next == nil {
		return errors.New("value not found")
	}

	if cur.Next == f.Tail {
		f.Tail = cur
	}
	cur.Next = cur.Next.Next
	return nil
}

// Вывести на экран
func (f *ForwardList) Print() {
	cur := f.Head
	for cur != nil {
		fmt.Print(cur.Data, " ")
		cur = cur.Next
	}
	fmt.Println()
}