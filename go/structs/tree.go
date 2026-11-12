package structs

import "fmt"

type FullBinaryTree struct {
	Root *TreeNode
}

// Добавление элемента
func (t *FullBinaryTree) Insert(data int) {
	t.Root = insertNode(t.Root, data)
}

func insertNode(node *TreeNode, data int) *TreeNode {
	if node == nil {
		return &TreeNode{Data: data}
	}
	if data < node.Data {
		node.Left = insertNode(node.Left, data)
	} else if data > node.Data {
		node.Right = insertNode(node.Right, data)
	}
	return node
}

// Поиск элемента
func (t *FullBinaryTree) Find(data int) bool {
	cur := t.Root
	for cur != nil {
		if data == cur.Data {
			return true
		} else if data < cur.Data {
			cur = cur.Left
		} else {
			cur = cur.Right
		}
	}
	return false
}

// Проверка на полноту (Full Binary Tree)
func (t *FullBinaryTree) IsFull() bool {
	return checkFull(t.Root)
}

func checkFull(node *TreeNode) bool {
	if node == nil {
		return true
	}
	if node.Left == nil && node.Right == nil {
		return true
	}
	if node.Left != nil && node.Right != nil {
		return checkFull(node.Left) && checkFull(node.Right)
	}
	return false
}

// Вывод в ширину (BFS)
func (t *FullBinaryTree) PrintBFS() {
	if t.Root == nil {
		return
	}
	queue := []*TreeNode{t.Root}
	for len(queue) > 0 {
		cur := queue[0]
		queue = queue[1:]
		fmt.Print(cur.Data, " ")
		if cur.Left != nil {
			queue = append(queue, cur.Left)
		}
		if cur.Right != nil {
			queue = append(queue, cur.Right)
		}
	}
	fmt.Println()
}

// Сбор данных для сохранения в файл (Pre-Order: корень-лево-право)
func (t *FullBinaryTree) GetPreOrder() []int {
	var result []int
	var traverse func(node *TreeNode)
	traverse = func(node *TreeNode) {
		if node == nil {
			return
		}
		result = append(result, node.Data)
		traverse(node.Left)
		traverse(node.Right)
	}
	traverse(t.Root)
	return result
}

// Красивый вывод дерева в консоль
func (t *FullBinaryTree) Print() {
	if t.Root == nil {
		fmt.Println("(empty)")
		return
	}
	printNode(t.Root, "", false)
}

func printNode(node *TreeNode, prefix string, isLeft bool) {
	if node == nil {
		return
	}

	fmt.Print(prefix)
	if isLeft {
		fmt.Print("|-- ")
	} else {
		fmt.Print("\\-- ")
	}
	fmt.Println(node.Data)

	newPrefix := prefix
	if isLeft {
		newPrefix += "|   "
	} else {
		newPrefix += "    "
	}

	if node.Left != nil || node.Right != nil {
		if node.Left != nil {
			printNode(node.Left, newPrefix, true)
		} else {
			fmt.Println(newPrefix + "|-- [null]")
		}

		if node.Right != nil {
			printNode(node.Right, newPrefix, false)
		} else {
			fmt.Println(newPrefix + "\\-- [null]")
		}
	}
}