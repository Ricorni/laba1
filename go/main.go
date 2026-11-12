package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"dbms/structs" 
)

func main() {
	// 1. Парсинг аргументов командной строки
	var filePath, query string
	for i := 1; i < len(os.Args); i++ {
		if os.Args[i] == "--file" && i+1 < len(os.Args) {
			filePath = os.Args[i+1]
			i++
		} else if os.Args[i] == "--query" && i+1 < len(os.Args) {
			query = os.Args[i+1]
			i++
		}
	}

	if filePath == "" || query == "" {
		fmt.Println("Usage: ./dbms --file <path> --query '<command>'")
		return
	}

	// 2. Инициализация и загрузка БД
	db := structs.NewDatabase()
	structs.Load(filePath, db)

	// 3. Парсинг запроса
	tokens := strings.Fields(query)
	if len(tokens) == 0 {
		return
	}

	cmd := tokens[0]
	name := ""
	value := ""
	index := ""

	if len(tokens) > 1 { name = tokens[1] }
	if len(tokens) > 2 { value = tokens[2] }
	if len(tokens) > 3 { index = tokens[3] }

	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Error:", r)
		}
	}()

	// 4. Выполнение команды
	switch cmd {
	// --- ARRAY ---
	case "MPUSH":
		if db.Arrays[name] == nil { db.Arrays[name] = &structs.ArrayList{} }
		db.Arrays[name].PushBack(value)
		fmt.Println(value)
	case "MGET":
		idx, _ := strconv.Atoi(value)
		val, err := db.Arrays[name].Get(idx)
		if err != nil { fmt.Println("Error:", err) } else { fmt.Println(val) }
	case "MINSERT":
		idx, _ := strconv.Atoi(index)
		err := db.Arrays[name].Insert(value, idx)
		if err != nil { fmt.Println("Error:", err) }
	case "MSET":
		idx, _ := strconv.Atoi(index)
		err := db.Arrays[name].Set(idx, value)
		if err != nil { fmt.Println("Error:", err) }
	case "MDEL":
		idx, _ := strconv.Atoi(value)
		err := db.Arrays[name].Remove(idx)
		if err != nil { fmt.Println("Error:", err) }
	case "MSIZE":
		if db.Arrays[name] != nil { fmt.Println(db.Arrays[name].Size()) } else { fmt.Println(0) }

	// --- FORWARD LIST ---
	case "FPUSH":
		if db.FLists[name] == nil { db.FLists[name] = &structs.ForwardList{} }
		db.FLists[name].PushBack(value)
		fmt.Println(value)
	case "FPUSHFRONT":
		if db.FLists[name] == nil { db.FLists[name] = &structs.ForwardList{} }
		db.FLists[name].PushFront(value)
		fmt.Println(value)
	case "FPOPFRONT":
		if db.FLists[name] != nil { db.FLists[name].PopFront() }
	case "FPOPBACK":
		if db.FLists[name] != nil { db.FLists[name].PopBack() }
	case "FINSERTAFTER":
		if db.FLists[name] != nil {
			err := db.FLists[name].InsertAfter(value, index)
			if err != nil { fmt.Println("Error:", err) }
		}
	case "FINSERTBEFORE":
		if db.FLists[name] != nil {
			err := db.FLists[name].InsertBefore(value, index)
			if err != nil { fmt.Println("Error:", err) }
		}
	case "FREMOVEAFTER":
		if db.FLists[name] != nil {
			err := db.FLists[name].RemoveAfter(value)
			if err != nil { fmt.Println("Error:", err) }
		}
	case "FREMOVEBEFORE":
		if db.FLists[name] != nil {
			err := db.FLists[name].RemoveBefore(value)
			if err != nil { fmt.Println("Error:", err) }
		}
	case "FDEL":
		if db.FLists[name] != nil {
			err := db.FLists[name].DeleteValue(value)
			if err != nil { fmt.Println("Error:", err) }
		}

	// --- DOUBLE LIST ---
	case "LPUSH":
		if db.DLists[name] == nil { db.DLists[name] = &structs.DoubleList{} }
		db.DLists[name].PushBack(value)
		fmt.Println(value)
	case "LPUSHFRONT":
		if db.DLists[name] == nil { db.DLists[name] = &structs.DoubleList{} }
		db.DLists[name].PushFront(value)
		fmt.Println(value)
	case "LPOPFRONT":
		if db.DLists[name] != nil { db.DLists[name].PopFront() }
	case "LPOPBACK":
		if db.DLists[name] != nil { db.DLists[name].PopBack() }
	case "LINSERTAFTER":
		if db.DLists[name] != nil {
			err := db.DLists[name].InsertAfter(value, index)
			if err != nil { fmt.Println("Error:", err) }
		}
	case "LINSERTBEFORE":
		if db.DLists[name] != nil {
			err := db.DLists[name].InsertBefore(value, index)
			if err != nil { fmt.Println("Error:", err) }
		}
	case "LREMOVEAFTER":
		if db.DLists[name] != nil {
			err := db.DLists[name].RemoveAfter(value)
			if err != nil { fmt.Println("Error:", err) }
		}
	case "LREMOVEBEFORE":
		if db.DLists[name] != nil {
			err := db.DLists[name].RemoveBefore(value)
			if err != nil { fmt.Println("Error:", err) }
		}
	case "LDEL":
		if db.DLists[name] != nil {
			err := db.DLists[name].DeleteValue(value)
			if err != nil { fmt.Println("Error:", err) }
		}

	// --- STACK ---
	case "SPUSH":
		if db.Stacks[name] == nil { db.Stacks[name] = &structs.Stack{} }
		db.Stacks[name].Push(value)
		fmt.Println(value)
	case "SPOP":
		if db.Stacks[name] != nil {
			val, err := db.Stacks[name].Pop()
			if err != nil { fmt.Println("Error:", err) } else { fmt.Println(val) }
		}
	case "SGET":
		if db.Stacks[name] != nil {
			val, err := db.Stacks[name].Front()
			if err != nil { fmt.Println("Error:", err) } else { fmt.Println(val) }
		}

	// --- QUEUE ---
	case "QPUSH":
		if db.Queues[name] == nil { db.Queues[name] = &structs.Queue{} }
		db.Queues[name].Push(value)
		fmt.Println(value)
	case "QPOP":
		if db.Queues[name] != nil {
			val, err := db.Queues[name].Pop()
			if err != nil { fmt.Println("Error:", err) } else { fmt.Println(val) }
		}
	case "QGET":
		if db.Queues[name] != nil {
			val, err := db.Queues[name].Front()
			if err != nil { fmt.Println("Error:", err) } else { fmt.Println(val) }
		}
	case "QBACK":
		if db.Queues[name] != nil {
			val, err := db.Queues[name].Back()
			if err != nil { fmt.Println("Error:", err) } else { fmt.Println(val) }
		}

	// --- TREE ---
	case "TINSERT":
		if db.Trees[name] == nil { db.Trees[name] = &structs.FullBinaryTree{} }
		val, _ := strconv.Atoi(value)
		db.Trees[name].Insert(val)
		fmt.Println(value)
	case "TGET":
		if db.Trees[name] != nil {
			val, _ := strconv.Atoi(value)
			if db.Trees[name].Find(val) { fmt.Println("TRUE") } else { fmt.Println("FALSE") }
		}
	case "TISFULL":
		if db.Trees[name] != nil && db.Trees[name].IsFull() {
			fmt.Println("TRUE")
		} else {
			fmt.Println("FALSE")
		}
	case "TBFS":
		if db.Trees[name] != nil { db.Trees[name].PrintBFS() }


	case "PRINT":
		if db.Arrays[name] != nil {
			db.Arrays[name].Print()
		} else if db.Stacks[name] != nil {
			db.Stacks[name].Print()
		} else if db.Queues[name] != nil {
			db.Queues[name].Print()
		} else if db.FLists[name] != nil {
			db.FLists[name].Print()
		} else if db.DLists[name] != nil {
			db.DLists[name].Print()
		} else if db.Trees[name] != nil {
			db.Trees[name].Print()
		} else {
			fmt.Println("Container not found")
		}

	default:
		fmt.Println("Unknown command:", cmd)
	}

	structs.Save(filePath, db)
}