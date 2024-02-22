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
	if(list.next == nil){
		next := *list;
		*list = Node{value: value, next: nil}
		list.next = &next
	}
}

// pushBack добавляет новый элемент в конец списка
func pushBack(list *Node, value int) {
	if(list.next != nil){
		node := list.next
		for node.next != nil {
			node = node.next
		}
		node.next = &Node{value: value, next: nil}
	}else{
		list.next = &Node{value: value, next: nil}
	}
}

// count возвращает кол-во элементов в списке
func count(list *Node) int {
	if list == nil {
		return 0
	}
	count := 1
	node := list.next
	for node != nil {
		node = node.next
		count++
	}
	return count
}

// popFront возвращает значение первого элемента и удаляет его из списка
func popFront(list *Node) int {
	if list == nil{
		return -1
	}
	if list.next != nil {
		value := list.value
		node := *list.next
		*list = node
		return value
	}
	return -1
}

// popBack возвращает значение последнего элемента и удаляет его из списка
func popBack(list *Node) int {
	var value int
	if list == nil{
		return -1
	}
	if list.next != nil {
		prevNode := list
		currNode := list.next
		for currNode.next != nil {
			prevNode = currNode
			currNode = prevNode.next
		}
		value = currNode.value
		prevNode.next = nil
		list = prevNode
		return value
	}
	value = list.value
	list = nil
	return value
	
}

// isValueInList ищет значение в списке и возвращает true, если оно найдено, в ином случае - false
func isValueInList(list *Node, value int) bool {
	if list == nil {
		return false
	}
	node := list
	for node.value != value && node.next != nil{
		node = node.next
	}
	return node.value == value
}

// getValueByIndex возвращает значение из списка по индексу, если оно есть, в ином случае - error("index out of range")
func getValueByIndex(list *Node, index int) (int, error) {
	if list == nil {
		return -1, fmt.Errorf("index out of range")
	}
	count := 0
	node := list
	for count != index && node.next != nil{
		node = node.next
		count ++
	}
	if(count != index){
		return -1, fmt.Errorf("index out of range")
	}
	return node.value, nil
}

// insert добавляет элемент в список в соответствующий индекс
func insert(list *Node, index, value int) {
	if count(list) < index {
		fmt.Println("Can't insert, empty linked list")
		return
	}
	count := 0
	var prevNode *Node
	curNode := list
	for count != index && curNode.next != nil {
		prevNode = curNode
		curNode = curNode.next
		count++
	}
	if(count == index){
		newNode := &Node{value: value, next: curNode}
		prevNode.next = newNode
		return
	}
	fmt.Println("Can't insert, index of element out of range")
}

// deleteByIndex удаляет элемент из списка по индексу и возвращает его значение. Если индекс неправильный - возвращает error("index out of range")
func deleteByIndex(list *Node, index int) (int, error) {
	if count(list) < index {
		return -1, fmt.Errorf("Can't delete, empty linked list")
	}
	count := 0
	var prevNode *Node
	curNode := list
	for count != index && curNode.next != nil {
		prevNode = curNode
		curNode = curNode.next
		count++
	}
	if(count == index){
		value := curNode.value
		prevNode.next = curNode.next
		return value, nil
	}
	return -1, fmt.Errorf("index out of range")
}

// sort сортирует список (*)
func sort(list *Node) {
	panic("not implemented")
}

func quickSortByCenter(list *Node) {
	quickSortByCenterImpl(list, 0, count(list) - 1)
}

func quickSortByCenterImpl(list *Node, start int, end int) {
	left := start
	right := end
	center, error := getValueByIndex(list, (left+right)/2)
	if error != nil {
		fmt.Println(error)
		return
	}

	for left <= right {
		for arr[right] > center {
			right--
		}
		for arr[left] < center {
			left++
		}
		if left <= right {
			arr[right], arr[left] = arr[left], arr[right]
			left++
			right--
		}
	}
	if right > start {
		quickSortByCenterImpl(arr, start, right)
	}
	if left < end {
		quickSortByCenterImpl(arr, left, end)
	}
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
	fmt.Println("Linked list after getting the value of the lst element and removing it: ")
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
}
