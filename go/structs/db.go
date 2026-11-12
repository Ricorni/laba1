package structs

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Database struct {
	Stacks map[string]*Stack
	Queues map[string]*Queue
	FLists map[string]*ForwardList
	DLists map[string]*DoubleList
	Arrays map[string]*ArrayList
	Trees  map[string]*FullBinaryTree
}

// Конструктор базы данных (создает пустые мапы)
func NewDatabase() *Database {
	return &Database{
		Stacks: make(map[string]*Stack),
		Queues: make(map[string]*Queue),
		FLists: make(map[string]*ForwardList),
		DLists: make(map[string]*DoubleList),
		Arrays: make(map[string]*ArrayList),
		Trees:  make(map[string]*FullBinaryTree),
	}
}

func Load(path string, db *Database) {
	file, err := os.Open(path)
	if err != nil {
		return // Если файла нет, просто выходим (создастся при сохранении)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var currentType, currentName string
	var values []string

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}

		if line == "END" {
			// Когда встречаем END, сохраняем накопленные данные в нужную структуру
			switch currentType {
			case "ARRAY":
				if db.Arrays[currentName] == nil {
					db.Arrays[currentName] = &ArrayList{}
				}
				for _, v := range values {
					db.Arrays[currentName].PushBack(v)
				}
			case "STACK":
				if db.Stacks[currentName] == nil {
					db.Stacks[currentName] = &Stack{}
				}
				for i := len(values) - 1; i >= 0; i-- {
					db.Stacks[currentName].Push(values[i])
				}
			case "QUEUE":
				if db.Queues[currentName] == nil {
					db.Queues[currentName] = &Queue{}
				}
				for _, v := range values {
					db.Queues[currentName].Push(v)
				}
			case "FLIST":
				if db.FLists[currentName] == nil {
					db.FLists[currentName] = &ForwardList{}
				}
				for _, v := range values {
					db.FLists[currentName].PushBack(v)
				}
			case "DLIST":
				if db.DLists[currentName] == nil {
					db.DLists[currentName] = &DoubleList{}
				}
				for _, v := range values {
					db.DLists[currentName].PushBack(v)
				}
			case "TREE":
				if db.Trees[currentName] == nil {
					db.Trees[currentName] = &FullBinaryTree{}
				}
				for _, v := range values {
					val, _ := strconv.Atoi(v)
					db.Trees[currentName].Insert(val)
				}
			}
			values = []string{}
			currentType = ""
			continue
		}

		parts := strings.Fields(line)

		if len(parts) == 2 && currentType == "" {
			currentType = parts[0]
			currentName = parts[1]
		} else {
			values = append(values, line)
		}
	}
}

func Save(path string, db *Database) {
	file, _ := os.Create(path)
	defer file.Close()

	// Сохраняем массивы
	for name, a := range db.Arrays {
		fmt.Fprintf(file, "ARRAY %s\n", name)
		for _, v := range a.Data {
			fmt.Fprintln(file, v)
		}
		fmt.Fprintln(file, "END")
	}

	// Сохраняем стеки
	for name, s := range db.Stacks {
		fmt.Fprintf(file, "STACK %s\n", name)
		cur := s.Head
		for cur != nil {
			fmt.Fprintln(file, cur.Data)
			cur = cur.Next
		}
		fmt.Fprintln(file, "END")
	}

	// Сохраняем очереди
	for name, q := range db.Queues {
		fmt.Fprintf(file, "QUEUE %s\n", name)
		cur := q.Head
		for cur != nil {
			fmt.Fprintln(file, cur.Data)
			cur = cur.Next
		}
		fmt.Fprintln(file, "END")
	}

	// Сохраняем односвязные списки
	for name, f := range db.FLists {
		fmt.Fprintf(file, "FLIST %s\n", name)
		cur := f.Head
		for cur != nil {
			fmt.Fprintln(file, cur.Data)
			cur = cur.Next
		}
		fmt.Fprintln(file, "END")
	}

	// Сохраняем двусвязные списки
	for name, d := range db.DLists {
		fmt.Fprintf(file, "DLIST %s\n", name)
		cur := d.Head
		for cur != nil {
			fmt.Fprintln(file, cur.Data)
			cur = cur.Next
		}
		fmt.Fprintln(file, "END")
	}

	// Сохраняем деревья (используем GetPreOrder)
	for name, t := range db.Trees {
		fmt.Fprintf(file, "TREE %s\n", name)
		for _, v := range t.GetPreOrder() {
			fmt.Fprintln(file, v)
		}
		fmt.Fprintln(file, "END")
	}
}