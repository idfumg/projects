package main

import "fmt"

type Node interface {
	SetValue(int)
	GetValue() int
}

type SimpleNode struct {
	value             int
	next              *SimpleNode
	simpleNodeMessage string
}

type PowerNode struct {
	value            int
	next             *PowerNode
	powerNodeMessage string
}

func (node *SimpleNode) SetValue(value int) {
	node.value = value
}

func (node *SimpleNode) GetValue() int {
	return node.value
}

func (node *PowerNode) SetValue(value int) {
	node.value = 10 * value
}

func (node *PowerNode) GetValue() int {
	return node.value
}

func NewSimpleNode() *SimpleNode {
	return &SimpleNode{simpleNodeMessage: "Message from a simple node"}
}

func NewPowerNode() *PowerNode {
	return &PowerNode{powerNodeMessage: "Message from a power node"}
}

func CreateNode() Node {
	node := NewPowerNode()
	node.SetValue(4)
	return node
}

func PrintType(i interface{}) {
	switch i := i.(type) {
	case string:
		fmt.Println("This is a string type", i)
	case int:
		fmt.Println("This is a int type", i)
	case float64:
		fmt.Println("This is a float64 type", i)
	}
}

type MagicStore struct {
	value interface{}
	name string
}

func (m *MagicStore) SetValue(value interface{}) {
	m.value = value
}

func (m *MagicStore) GetValue() interface{} {
	return m.value
}

func NewMagicStore(name string) *MagicStore {
	return &MagicStore{name: name}
}

func main() {
	node := CreateNode()
	switch concreteType := node.(type) {
	case *SimpleNode:
		fmt.Println("Type is SimpleNode, message:", concreteType.simpleNodeMessage)
	case *PowerNode:
		fmt.Println("Type is PowerNode, message:", concreteType.powerNodeMessage)
	}

	PrintType("hello")
	PrintType(1)
	PrintType(4.0)

	istore := NewMagicStore("Integer Store")
	istore.SetValue(4)
	if v, ok := istore.GetValue().(int); ok {
		v *= 10
		fmt.Println(v)
	}
}
