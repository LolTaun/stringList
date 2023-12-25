package stringlist

// Типы данных:
// Node: элемент списка, хранящий символ и ссылку на следующий элемент.

type Node struct {
	Symbol rune
	Next *Node
};