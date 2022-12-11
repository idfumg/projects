package main

import "fmt"

type NodeI interface {
	SetValue(int)
	GetValue() int
}

type Node struct {
	value int
	next  *Node
}

func (node *Node) SetValue(value int) {
	node.value = value
}

func (node *Node) GetValue() int {
	return node.value
}

func NewNode() *Node {
	return new(Node)
}

type PowerNode struct {
	value int
	next *Node
}

func (node *PowerNode) SetValue(value int) {
	node.value = value * 10
}

func (node *PowerNode) GetValue() int {
	return node.value
}

func NewPowerNode() *PowerNode {
	return new(PowerNode)
}

type SingleLinkedList struct {
	head *Node
	tail *Node
}

func NewSingleLinkedList() *SingleLinkedList {
	return new(SingleLinkedList)
}

func (list *SingleLinkedList) Add(value int) {
	newNode := &Node{value: value}
	if list.head == nil {
		list.head = newNode
	} else if list.head == list.tail {
		list.head.next = newNode
	} else {
		list.tail.next = newNode
	}
	list.tail = newNode
}

func (list *SingleLinkedList) String() string {
	ans := ""
	for n := list.head; n != nil; n = n.next {
		ans += fmt.Sprintf(" {%d}", n.value)
	}
	return ans
}

type Level int

func (level *Level) raiseShieldLevel(i int) {
	if *level == 0 {
		*level = 1
	}
	*level = *level * Level(i)
}

func main() {
	node := NewNode()
	node.SetValue(3)
	fmt.Println("Node has the value of:", node.GetValue())

	shieldLevel := new(Level)
	shieldLevel.raiseShieldLevel(4)
	shieldLevel.raiseShieldLevel(5)
	fmt.Println(*shieldLevel) // 20

	shieldLevel2 := Level(1)
	shieldLevel2.raiseShieldLevel(4)
	shieldLevel2.raiseShieldLevel(5)
	fmt.Println(shieldLevel2) // 20

	n1 := NewNode()
	n1.SetValue(4)
	n2 := NewPowerNode()
	n2.SetValue(4)
	fmt.Printf("n1: %v, n2: %v\n", n1.value, n2.value)

	list := NewSingleLinkedList()
	list.Add(2)
	list.Add(3)
	list.Add(4)
	list.Add(5)
	fmt.Println("List contains: ", list)
}
