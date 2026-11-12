package structs

// Узел для Односвязного списка, Стека и Очереди
type Node struct {
	Data string
	Next *Node
}

// Узел для Двусвязного списка
type DoubleNode struct {
	Data string
	Next *DoubleNode
	Prev *DoubleNode
}

// Узел для Бинарного дерева 
type TreeNode struct {
	Data  int
	Left  *TreeNode
	Right *TreeNode
}