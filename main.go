package main

import "fmt"

// DO NOT EDIT ---------------------------------------------------------------------------------------------------------
type Node struct {
	value int
	next  *Node
}

func (n *Node) Value() int {
	return n.value
}

func (n *Node) Next() *Node {
	return n.next
}

func printList(list *Node) {
	curNode := list
	for i := 0; curNode != nil; i++ {
		fmt.Printf("%d: %d\n", i, curNode.Value())
		curNode = curNode.Next()
	}
}

// ---------------------------------------------------------------------------------------------------------------------

// pushFront добавляет новый элемент в начало списка
func pushFront(list *Node, value int) {
	panic("not implemented")
}

// pushBack добавляет новый элемент в конец списка
func pushBack(list *Node, value int) {
	panic("not implemented")
}

// count возвращает кол-во элементов в списке
func count(list *Node) int {
	panic("not implemented")
}

// popFront возвращает значение первого элемента и удаляет его из списка
func popFront(list *Node) int {
	panic("not implemented")
}

// popBack возвращает значение последнего элемента и удаляет его из списка
func popBack(list *Node) int {
	panic("not implemented")
}

// isValueInList ищет значение в списке и возвращает true, если оно найдено, в ином случае - false
func isValueInList(list *Node, value int) bool {
	panic("not implemented")
}

// getValueByIndex возвращает значение из списка по индексу, если оно есть, в ином случае - error("index out of range")
func getValueByIndex(list *Node, index int) (int, error) {
	panic("not implemented")
}

// insert добавляет элемент в список в соответствующий индекс
func insert(list *Node, index, value int) {
	panic("not implemented")
}

// deleteByIndex удаляет элемент из списка по индексу и возвращает его значение. Если индекс неправильный - возвращает error("index out of range")
func deleteByIndex(list *Node, index int) (int, error) {
	panic("not implemented")
}

// sort сортирует список (*)
func sort(list *Node) {
	panic("not implemented")
}

func main() {
	linkedList := &Node{
		value: 0,
		next:  nil,
	}
	printList(linkedList)
}
