package main

import (
	"fmt"
)

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
	next := *list;
	*list = Node{value: value, next: &next}
}

// pushBack добавляет новый элемент в конец списка
func pushBack(list *Node, value int) {
	node := list
	for node.next != nil {
		node = node.Next()
	}
	node.next = &Node{value: value, next: nil}
}

// count возвращает кол-во элементов в списке
func count(list *Node) int {
	if list == nil{
		return 0
	}
	count := 0
	node := list
	for ; node != nil; count++ {
		node = node.Next()
	}
	return count
}

// popFront возвращает значение первого элемента и удаляет его из списка
func popFront(list *Node) int {
	value := list.value
	node := *list.Next()
	*list = node
	return value
}

// popBack возвращает значение последнего элемента и удаляет его из списка
func popBack(list *Node) int {
	var value int
	node := list
	for node.Next() != nil {
		node = node.Next()
	}
	previus(list, index(list, node)).next = nil
	return value
}

// isValueInList ищет значение в списке и возвращает true, если оно найдено, в ином случае - false
func isValueInList(list *Node, value int) bool {
	node := list
	for node.value != value && node.next != nil{
		node = node.next
	}
	return node.value == value
}

// getValueByIndex возвращает значение из списка по индексу, если оно есть, в ином случае - error("index out of range")
func getValueByIndex(list *Node, index int) (int, error) {
	node := list
	for i := 0; node != nil; i++ {
		if i == index {
			return node.value, nil
		}
		node = node.next
	}
	return -1, fmt.Errorf("index out of range")
}

// insert добавляет элемент в список в соответствующий индекс
func insert(list *Node, index, value int) {
	if count(list) < index {
		fmt.Println("Can't insert, index out of range")
		return
	}
	if (index == 0){
		*list = Node{value: value, next: list}
		return
	}
	node := list
	for i := 0; node != nil; i++ {
		if i == index {
			previus(list, index).next = &Node{value: value, next: nodeByIndex(list, index)}
		}
		node = node.Next()
	}
}

// deleteByIndex удаляет элемент из списка по индексу и возвращает его значение. Если индекс неправильный - возвращает error("index out of range")
func deleteByIndex(list *Node, index int) (int, error) {
	if count(list) < index {
		return -1, fmt.Errorf("Can't delete, index out of range")
	}
	var value int
	if index == 0 {
		value = list.value
		*list = *list.Next()
		return value, nil
	}
	node := list
	for i := 0; node != nil; i++ {
		if i == index {
			value = node.value
			if node.next == nil {
				previus(list, index).next = nil
			} else {
				previus(list, index).next = node.Next()
			}
			return value, nil
		}
		node = node.Next()
	}
	return -1, fmt.Errorf("index out of range")
}

func nodeByIndex(list *Node, index int) *Node {
	count := 0
	node := list
	for count != index && node.next != nil {
		node = node.Next()
		count++
	}
	if count == index {
		return node
	}
	return nil
}

func previus(list *Node, index int) *Node {
	node := list
	for i := 0; node != nil; i++ {
		if i == index - 1 {
			return node
		}
		node = node.next
	}
	return nil
}

func next(list *Node, index int) *Node {
	node := list
	for i := 0; node != nil; i++ {
		if i == index + 1 {
			return node
		}
		node = node.next
	}
	return nil
}

func swap(list *Node, left, right int) *Node{
	fmt.Println("-------")
	printList(list)

	leftPrev := previus(list, left)
	leftNode := nodeByIndex(list, left)
	leftNext := next(list, left)

	rightPrev := previus(list, right)
	rightNode := nodeByIndex(list, right)
	rightNext := next(list, right)

	leftNode.next = rightNext
	if rightPrev == leftNode {
		rightNode.next = leftNode
	} else {
		rightPrev.next = leftNode
		rightNode.next = leftNext
	}

	if leftPrev == nil {
		list = rightNode
	} else {
		leftPrev.next = rightNode
	}
	return list
}

func index(list, node *Node) int {
	curNode := list
	for i := 0; curNode != nil; i++ {
		if curNode == node {
			return i
		}
		curNode = curNode.Next()
	}
	return -1
}

// sort сортирует список (*)
func sort(list *Node) {
	//l1 := quickSort(list, 0, count(list) - 1)
	l2 := *quickSort(list, 0, count(list) - 1)
	// fmt.Println(l1)
	//fmt.Println(l2)
	*list = l2
	//*list = *quickSort(list, 0, count(list) - 1)
}

func quickSort(list *Node, start int, end int) *Node {
	left := start
	right := end
	center,_ := getValueByIndex(list, (left + right) / 2)
	for left <= right {
		value, _ := getValueByIndex(list, right)
		for value > center {
			right--
			value, _ = getValueByIndex(list, right)
		}
		value, _ = getValueByIndex(list, left)
		for value < center {
			left++
			value, _ = getValueByIndex(list, left)
		}
		if left <= right {
			list = swap(list, left, right)
			left++
			right--
		}
	}
	if right > start {
		list = quickSort(list, start, right)
	}
	if left < end {
		list = quickSort(list, left, end)
	}
	return list
}

func main() {
	linkedList := &Node{
		value: 0,
		next:  nil,
	}
	fmt.Printf("Linked list after initialization: ")
	printList(linkedList)
	pushFront(linkedList, 5)
	fmt.Println("Linked list after inserting a value at the beginning: ")
	printList(linkedList)
	pushBack(linkedList, 3)
	fmt.Println("Linked list after inserting a value at the end: ")
	printList(linkedList)
	fmt.Println("Lenght of linked list: ", count(linkedList))
	popFront(linkedList)
	fmt.Println("Linked list after getting the value of the first element and removing it: ")
	printList(linkedList)
	popBack(linkedList)
	fmt.Println("Linked list after getting the value of the last element and removing it: ")
	printList(linkedList)
	pushFront(linkedList, 5)
	pushBack(linkedList, 3)

	fmt.Println("Value 3 is present in the linked list: ", isValueInList(linkedList, 3))
	fmt.Println("Value 5 is present in the linked list: ", isValueInList(linkedList, 5))
	fmt.Println("Value 0 is present in the linked list: ", isValueInList(linkedList, 0))
	fmt.Println("Value 7 is present in the linked list: ", isValueInList(linkedList, 7))

	value, error := getValueByIndex(linkedList, 7)
	if error == nil {
		fmt.Println("Value of element at index 7: ", value)
	}else{
		fmt.Println("7 - ", error)
	}

	value, error = getValueByIndex(linkedList, 0)
	if error == nil {
		fmt.Println("Value of element at index 0: ", value)
	}else{
		fmt.Println("0 - ", error)
	}

	value, error = getValueByIndex(linkedList, 2)
	if error == nil {
		fmt.Println("Value of element at index 2: ", value)
	}else{
		fmt.Println("2 - ", error)
	}

	fmt.Println("Linked list after insert value 6 at postion 1: ")
	insert(linkedList, 1, 6)
	printList(linkedList)

	fmt.Println("Linked list after insert value 9 at postion 2: ")
	insert(linkedList, 2, 9)
	printList(linkedList)

	fmt.Println("Linked list after delete element at postion 3: ")
	value, error = deleteByIndex(linkedList, 3)
	if error != nil{
		fmt.Println(error)
	}else{
		fmt.Println("Value of delete element is ", value)
		printList(linkedList)
	}

	fmt.Println("Linked list after delete element at postion 4: ")
	value, error = deleteByIndex(linkedList, 4)
	if error != nil{
		fmt.Println(error)
	}else{
		fmt.Println("Value of delete element is ", value)
		printList(linkedList)
	}

	fmt.Println("Linked list after delete element at postion 0: ")
	value, error = deleteByIndex(linkedList, 0)
	if error != nil{
		fmt.Println(error)
	}else{
		fmt.Println("Value of delete element is ", value)
		printList(linkedList)
	}

	fmt.Println("Linked list after delete element at postion 2: ")
	value, error = deleteByIndex(linkedList, 2)
	if error != nil{
		fmt.Println(error)
	}else{
		fmt.Println("Value of delete element is ", value)
		printList(linkedList)
	}
	fmt.Println("-------")

	insert(linkedList, 1, 3)
	insert(linkedList, 1, 2)
	insert(linkedList, 3, 5)
	printList(linkedList)
	fmt.Println("-------")

	fmt.Println("\nSorted linked list")
	sort(linkedList)
	printList(linkedList)
}
