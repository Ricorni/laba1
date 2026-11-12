package structs

import (
	"errors"
	"fmt"
)

type DoubleList struct {
	Head *DoubleNode
	Tail *DoubleNode
}

// Добавить в конец
func (d *DoubleList) PushBack(data string) {
	newNode := &DoubleNode{Data: data, Prev: d.Tail}
	if d.Tail != nil {
		d.Tail.Next = newNode
	} else {
		d.Head = newNode
	}
	d.Tail = newNode
}

// Добавить в начало
func (d *DoubleList) PushFront(data string) {
	newNode := &DoubleNode{Data: data, Next: d.Head}
	if d.Head != nil {
		d.Head.Prev = newNode
	} else {
		d.Tail = newNode
	}
	d.Head = newNode
}

// Поиск узла
func (d *DoubleList) find(data string) (*DoubleNode, error) {
	cur := d.Head
	for cur != nil {
		if cur.Data == data {
			return cur, nil
		}
		cur = cur.Next
	}
	return nil, errors.New("position element not found")
}

// Вставить после элемента
func (d *DoubleList) InsertAfter(data string, target string) error {
	node, err := d.find(target)
	if err != nil {
		return err
	}

	newNode := &DoubleNode{Data: data, Next: node.Next, Prev: node}
	if node.Next != nil {
		node.Next.Prev = newNode
	} else {
		d.Tail = newNode
	}
	node.Next = newNode
	return nil
}

// Вставить до элемента
func (d *DoubleList) InsertBefore(data string, target string) error {
	node, err := d.find(target)
	if err != nil {
		return err
	}

	newNode := &DoubleNode{Data: data, Next: node, Prev: node.Prev}
	if node.Prev != nil {
		node.Prev.Next = newNode
	} else {
		d.Head = newNode
	}
	node.Prev = newNode
	return nil
}

// Удалить из начала
func (d *DoubleList) PopFront() error {
	if d.Head == nil {
		return errors.New("list is empty")
	}
	d.Head = d.Head.Next
	if d.Head != nil {
		d.Head.Prev = nil
	} else {
		d.Tail = nil
	}
	return nil
}

// Удалить из конца
func (d *DoubleList) PopBack() error {
	if d.Tail == nil {
		return errors.New("list is empty")
	}
	d.Tail = d.Tail.Prev
	if d.Tail != nil {
		d.Tail.Next = nil
	} else {
		d.Head = nil
	}
	return nil
}

// Удалить после элемента
func (d *DoubleList) RemoveAfter(target string) error {
	node, err := d.find(target)
	if err != nil {
		return err
	}
	if node.Next == nil {
		return errors.New("no node after this one")
	}

	delNode := node.Next
	node.Next = delNode.Next
	if delNode.Next != nil {
		delNode.Next.Prev = node
	} else {
		d.Tail = node
	}
	return nil
}

// Удалить до элемента
func (d *DoubleList) RemoveBefore(target string) error {
	node, err := d.find(target)
	if err != nil {
		return err
	}
	if node.Prev == nil {
		return errors.New("no node before this one")
	}

	delNode := node.Prev
	node.Prev = delNode.Prev
	if delNode.Prev != nil {
		delNode.Prev.Next = node
	} else {
		d.Head = node
	}
	return nil
}

// Удалить по значению
func (d *DoubleList) DeleteValue(data string) error {
	node, err := d.find(data)
	if err != nil {
		return err
	}

	if node.Prev != nil {
		node.Prev.Next = node.Next
	} else {
		d.Head = node.Next
	}

	if node.Next != nil {
		node.Next.Prev = node.Prev
	} else {
		d.Tail = node.Prev
	}
	return nil
}

// Вывести на экран
func (d *DoubleList) Print() {
	cur := d.Head
	for cur != nil {
		fmt.Print(cur.Data, " ")
		cur = cur.Next
	}
	fmt.Println()
}