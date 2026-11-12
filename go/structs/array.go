package structs

import (
	"errors"
	"fmt"
)

type ArrayList struct {
	Data []string
}

// Добавить в конец
func (a *ArrayList) PushBack(value string) {
	a.Data = append(a.Data, value)
}

// Вставить по индексу
func (a *ArrayList) Insert(value string, index int) error {
	if index < 0 || index > len(a.Data) {
		return errors.New("index out of range")
	}

	if index == len(a.Data) {
		a.PushBack(value)
		return nil
	}

	a.Data = append(a.Data, "")

	copy(a.Data[index+1:], a.Data[index:])
	
	a.Data[index] = value
	return nil
}

// Получить по индексу
func (a *ArrayList) Get(index int) (string, error) {
	if index < 0 || index >= len(a.Data) {
		return "", errors.New("index out of range")
	}
	return a.Data[index], nil
}

// Заменить по индексу
func (a *ArrayList) Set(index int, value string) error {
	if index < 0 || index >= len(a.Data) {
		return errors.New("index out of range")
	}
	a.Data[index] = value
	return nil
}

// Удалить по индексу
func (a *ArrayList) Remove(index int) error {
	if index < 0 || index >= len(a.Data) {
		return errors.New("index out of range")
	}
	a.Data = append(a.Data[:index], a.Data[index+1:]...)
	return nil
}

// Получить размер
func (a *ArrayList) Size() int {
	return len(a.Data)
}

// Вывести на экран
func (a *ArrayList) Print() {
	for _, v := range a.Data {
		fmt.Print(v, " ")
	}
	fmt.Println()
}